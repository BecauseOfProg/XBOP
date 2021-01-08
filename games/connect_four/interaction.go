package connect_four

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func HandleInteraction(bot *onyxcord.Bot, message *discordgo.Message) {
	connectFourPlayer := bot.Cache.Exists(context.Background(), "connectfour:"+message.ChannelID).Val()
	if connectFourPlayer == 1 {
		handlePlay(bot, message, "connectfour:"+message.ChannelID)
	}
}
