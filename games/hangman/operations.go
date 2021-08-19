package hangman

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func startGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, maxErrors int) (err error) {
	rand.Seed(time.Now().UnixNano())
	word := strings.TrimRight(words[rand.Intn(len(words))], "\r")

	letters := string(word[0])

	_ = bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    formatMessage(word, letters, "", maxErrors, ""),
			Components: stopButton(),
		},
	})

	bot.Cache.HSet(context.Background(), "hangman:"+interaction.ChannelID,
		"word", word,
		"letters", letters,
		"wrongLetters", "",
		"maxErrors", maxErrors,
		"game", interaction.Token,
	)

	return
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
