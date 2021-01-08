package connect_four

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
	"strconv"
)

func generateGrid(rows []string) string {
	var output string

	for row := 0; row < 7; row++ {
		for column := 0; column < 7; column++ {
			if row == 0 {
				output += numbers[column]
			} else {
				index, _ := strconv.Atoi(string(rows[column][6-row]))
				output += tokens[index]
			}
		}
		output += "\n"
	}

	return output
}

func sendTurnMessage(bot *onyxcord.Bot, user *discordgo.User, channel string, userToken int) string {
	message, _ := bot.Client.ChannelMessageSend(
		channel,
		fmt.Sprintf(
			"**:arrow_right: %s, à votre tour** (vous êtes %s)\n*Pour quitter la partie, envoyez `stop`.*",
			user.Mention(),
			tokens[userToken],
		),
	)
	return message.ID
}
