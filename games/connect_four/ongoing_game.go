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

func StopGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) error {
	cacheID := "connectfour:" + interaction.ChannelID
	bot.Cache.Del(context.Background(), cacheID, cacheID+"/grid")

	bot.Client.ChannelMessageSend(interaction.ChannelID, "**:stop_sign: Arrêt de la partie prononcé.**")
	return bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						stopButton(true),
					},
				},
			},
		},
	})
}
