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
			bot.Cache.HMSet(context.Background(), "verbs:"+interaction.ChannelID,
				"answers", 0,
				"successfulAnswers", 0,
			)

			sendQuestion(bot, interaction.ChannelID)

			return
		},
	}
}
