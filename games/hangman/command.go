package hangman

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			rand.Seed(time.Now().UnixNano())
			word := strings.TrimRight(words[rand.Intn(len(words))], "\r")

			var maxErrors int
			options := interaction.ApplicationCommandData().Options
			if len(options) == 0 {
				maxErrors = 7
			} else {
				maxErrors = int(options[0].IntValue())
			}
			letters := string(word[0])

			_ = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    formatMessage(word, letters, "", maxErrors, ""),
					Components: stopButton(false),
				},
			})

			bot.Cache.HSet(context.Background(), "hangman:"+interaction.ChannelID,
				"word", word,
				"letters", letters,
				"wrongLetters", "",
				"maxErrors", maxErrors,
				"game", interaction.Token,
			)

			return
		},
	}
}
