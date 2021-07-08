package connect_four

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var tokens = []string{":white_large_square:", ":red_circle:", ":yellow_circle:"}

var numbers = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣"}

func components(columns []string, disabled bool) (list []discordgo.MessageComponent) {
	list = selectButtons(columns, disabled)
	list = append(list, discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			stopButton(disabled),
		},
	})

	return
}

func stopButton(disabled bool) discordgo.Button {
	return discordgo.Button{
		Label:    "Arrêter la partie",
		Style:    discordgo.DangerButton,
		CustomID: "connectfour_stop",
		Disabled: disabled,
	}
}

func selectButtons(columns []string, disabled bool) (buttons []discordgo.MessageComponent) {
	var row discordgo.ActionsRow
	for i := 0; i < 7; i++ {
		if strings.Contains(columns[i], "0") {
			row.Components = append(row.Components, discordgo.Button{
				Style:    discordgo.PrimaryButton,
				Disabled: disabled,
				Emoji: discordgo.ComponentEmoji{
					Name: numbers[i],
				},
				CustomID: fmt.Sprintf("connectfour_turn_%d", i),
			})
		}
		if len(row.Components) == 5 || i == 6 {
			buttons = append(buttons, row)
			row = discordgo.ActionsRow{}
		}
	}

	return
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

func generateTurnMessage(user *discordgo.Member, token int) string {
	return fmt.Sprintf(
		"**:arrow_right: %s, à votre tour** (vous êtes %s)\n",
		user.Mention(),
		tokens[token],
	)
}
