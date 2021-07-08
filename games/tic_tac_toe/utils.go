package tic_tac_toe

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var tokens = []string{"⬛", "❌", "⭕"}

var winningGrids = [][][]int{
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
}

func stopButton(disabled bool) discordgo.Button {
	return discordgo.Button{
		Label:    "Arrêter la partie",
		Style:    discordgo.DangerButton,
		CustomID: "tictactoe_stop",
		Disabled: disabled,
	}
}

func generateGrid(columns []string, disabled bool) (grid []discordgo.MessageComponent) {
	for columnIndex, column := range columns {
		var buttons []discordgo.MessageComponent
		for rowIndex := 0; rowIndex < 3; rowIndex++ {
			token, _ := strconv.Atoi(string(column[rowIndex]))
			var style discordgo.ButtonStyle
			if token == 0 {
				style = discordgo.SecondaryButton
			} else {
				style = discordgo.PrimaryButton
			}

			buttons = append(buttons, discordgo.Button{
				CustomID: fmt.Sprintf("tictactoe_%d_%d", columnIndex, rowIndex),
				Style:    style,
				Disabled: disabled,
				Emoji: discordgo.ButtonEmoji{
					Name: tokens[token],
				},
			})
		}
		grid = append(grid, discordgo.ActionsRow{
			Components: buttons,
		})
	}
	grid = append(grid, discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			stopButton(disabled),
		},
	})

	return
}

func generateTurnMessage(user *discordgo.User, token int) string {
	return fmt.Sprintf(
		"**:arrow_right: %s, à votre tour** (vous êtes %s)",
		user.Mention(),
		tokens[token],
	)
}
