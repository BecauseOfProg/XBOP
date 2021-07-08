package connect_four

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var tokens = []string{":white_large_square:", ":red_circle:", ":yellow_circle:"}

var numbers = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣"}

func components(columns []string, disabled bool) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				selectMenu(columns, disabled),
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				stopButton(disabled),
			},
		},
	}
}

func stopButton(disabled bool) discordgo.Button {
	return discordgo.Button{
		Label:    "Arrêter la partie",
		Style:    discordgo.DangerButton,
		CustomID: "connectfour_stop",
		Disabled: disabled,
	}
}

func selectMenu(columns []string, _ bool) discordgo.SelectMenu {
	var options []discordgo.SelectMenuOption
	for i := 0; i < 7; i++ {
		if strings.Contains(columns[i], "0") {
			options = append(options, discordgo.SelectMenuOption{
				Label: fmt.Sprintf("Colonne %d", i+1),
				Value: strconv.Itoa(i),
				Emoji: discordgo.ComponentEmoji{
					Name: numbers[i],
				},
			})
		}
	}

	return discordgo.SelectMenu{
		CustomID:    "connectfour_turn",
		Placeholder: "Colonne où insérer le jeton",
		MinValues:   1,
		MaxValues:   1,
		Options:     options,
	}
}

func generateGrid(columns []string) string {
	var output string

	for rowIndex := 0; rowIndex < 7; rowIndex++ {
		for columnIndex := 0; columnIndex < 7; columnIndex++ {
			if rowIndex == 0 {
				output += numbers[columnIndex]
			} else {
				index, _ := strconv.Atoi(string(columns[columnIndex][6-rowIndex]))
				output += tokens[index]
			}
		}
		output += "\n"
	}

	return output
}

func generateTurnMessage(user *discordgo.User, token int) string {
	return fmt.Sprintf(
		"**:arrow_right: %s, à votre tour** (vous êtes %s)\n",
		user.Mention(),
		tokens[token],
	)
}
