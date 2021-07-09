package hangman

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func HandleOngoingGame(bot *onyxcord.Bot, message *discordgo.Message) error {
	hangmanPlayer := bot.Cache.Exists(context.Background(), "hangman:"+message.ChannelID).Val()
	if hangmanPlayer == 1 {
		return handleAttempt(bot, message, "hangman:"+message.ChannelID)
	}
	return nil
}

func HandleInteraction(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, _ []string) error {
	word := bot.Cache.HGet(context.Background(), "hangman:"+interaction.ChannelID, "word").Val()
	return stopGame(bot, interaction, interaction.ChannelID, fmt.Sprintf("**:stop_sign: Arrêt de la partie prononcé par %s.**\n\nLe mot à trouver était **%s**.", interaction.Member.Mention(), word))
}

func stopGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, channelID, reason string) (err error) {
	cacheID := "hangman:" + channelID

	word := bot.Cache.HGet(context.Background(), cacheID, "word").Val()
	letters := bot.Cache.HGet(context.Background(), cacheID, "letters").Val()
	wrongLetters := bot.Cache.HGet(context.Background(), cacheID, "wrongLetters").Val()
	maxErrors, _ := bot.Cache.HGet(context.Background(), cacheID, "maxErrors").Int()

	err = editMessage(bot, interaction, channelID, word, letters, wrongLetters, maxErrors, reason, true)
	bot.Cache.Del(context.Background(), cacheID)
	return
}

func editMessage(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, channelID, word, letters, wrongLetters string, maxErrors int, message string, disabled bool) (err error) {
	token := bot.Cache.HGet(context.Background(), "hangman:"+channelID, "game").Val()

	if interaction == nil {
		_, err = bot.Client.InteractionResponseEdit(bot.Config.Bot.ID, &discordgo.Interaction{Token: token}, &discordgo.WebhookEdit{
			Content:    formatMessage(word, letters, wrongLetters, maxErrors, message),
			Components: stopButton(disabled),
		})
	} else {
		err = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content:    formatMessage(word, letters, wrongLetters, maxErrors, message),
				Components: stopButton(disabled),
			},
		})
	}
	return
}
