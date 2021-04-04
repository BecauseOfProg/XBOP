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
			if len(interaction.Data.Options) == 0 {
				maxErrors = 7
			} else {
				maxErrors = int(interaction.Data.Options[0].IntValue())
			}

			_ = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: "**:chains: Et c'est parti pour un jeu du pendu !**\nTous les utilisateurs ayant accès au salon peuvent participer.\nPour arrêter la partie, envoyez `stop`.",
				},
			})

			letters := string(word[0])
			game, _ := bot.Client.ChannelMessageSend(interaction.ChannelID, formatMessage(word, letters, "", maxErrors))

			bot.Cache.HMSet(context.Background(), "hangman:"+interaction.ChannelID,
				"word", word,
				"letters", letters,
				"falseLetters", "",
				"maxErrors", maxErrors,
				"message", game.ID,
			)

			return
		},
	}
}
