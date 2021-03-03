package connect_four

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func HandleOngoingGame(bot *onyxcord.Bot, message *discordgo.Message) {
	connectFourPlayer := bot.Cache.Exists(context.Background(), "connectfour:"+message.ChannelID).Val()
	if connectFourPlayer == 1 {
		handleTurn(bot, message, "connectfour:"+message.ChannelID)
	}
}
