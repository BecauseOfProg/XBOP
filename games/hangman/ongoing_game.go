package hangman

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func HandleOngoingGame(bot *onyxcord.Bot, message *discordgo.Message) {
	hangmanPlayer := bot.Cache.Exists(context.Background(), "hangman:"+message.ChannelID).Val()
	if hangmanPlayer == 1 {
		handleAttempt(bot, message, "hangman:"+message.ChannelID)
	}
}

func StopGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) error {
	cacheID := "hangman:" + interaction.ChannelID
	word := bot.Cache.HGet(context.Background(), cacheID, "word").Val()
	bot.Cache.Del(context.Background(), cacheID)

	bot.Client.ChannelMessageSend(
		interaction.ChannelID,
		fmt.Sprintf("**:stop_sign: Arrêt de la partie prononcé.**\n\nLe mot à trouver était **%s**.", word),
	)
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
