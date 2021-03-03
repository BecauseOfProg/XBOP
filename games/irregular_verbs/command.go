package irregular_verbs

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			err = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseACKWithSource,
			})
			if err != nil {
				return
			}

			var part string
			if len(interaction.Data.Options) < 1 {
				part = "all"
			} else {
				part = interaction.Data.Options[0].StringValue()
			}

			bot.Cache.HMSet(context.Background(), "verbs:"+interaction.ChannelID,
				"answers", 0,
				"part", part,
				"successfulAnswers", 0,
			)

			sendQuestion(bot, interaction.ChannelID)
			return
		},
	}
}
