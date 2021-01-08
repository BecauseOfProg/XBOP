package irregular_verbs

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func HandleInteraction(bot *onyxcord.Bot, message *discordgo.Message) {
	verbsPlayer := bot.Cache.Exists(context.Background(), "verbs:"+message.ChannelID).Val()
	if verbsPlayer == 1 {
		handleAnswer(bot, message, "verbs:"+message.ChannelID)
	}
}
