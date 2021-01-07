package irregular_verbs

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		Description:    "Lancer un quiz sur les verbes irréguliers en anglais",
		Category:       "quizzes",
		Show:           true,
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(arguments []string, bot onyxcord.Bot, message *discordgo.MessageCreate) (err error) {
			bot.Cache.HMSet(context.Background(), "verbs:"+message.ChannelID,
				"answers", 0,
				"successfulAnswers", 0,
			)

			bot.Client.ChannelMessageSend(message.ChannelID, ":flag_gb: **Quiz sur les verbes irréguliers**")
			SendQuestion(&bot, message.ChannelID)

			return
		},
	}
}
