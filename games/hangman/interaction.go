package hangman

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func HandleInteraction(bot *onyxcord.Bot, message *discordgo.Message) {
	hangmanPlayer := bot.Cache.Exists(context.Background(), "hangman:"+message.ChannelID).Val()
	if hangmanPlayer == 1 {
		Try(bot, message, "hangman:"+message.ChannelID)
	}
}
