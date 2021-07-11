package irregular_verbs

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func HandleOngoingGame(bot *onyxcord.Bot, message *discordgo.Message) {
	if isGameOngoing := bot.Cache.Exists(context.Background(), "verbs:"+message.ChannelID).Val(); isGameOngoing == 1 {
		handleAnswer(bot, message, "verbs:"+message.ChannelID)
	}
}
