package hangman

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/BecauseOfProg/xbop/lib"
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func startGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, word string, maxErrors int) (err error) {
	if maxErrors > 99 {
		return bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: ":x: Le nombre maximum d'erreurs autorisées est de 99.",
				Flags:   1 << 6,
			},
		})
	}

	rand.Seed(time.Now().UnixNano())
	if word == "" {
		word = strings.TrimRight(words[rand.Intn(len(words))], "\r")
	}

	word = strings.ToUpper(lib.TrimNonLetters(word))

	if len(word) > 100 {
		return bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: ":x: Votre mot ne peut dépasser les 100 caractères.",
				Flags:   1 << 6,
			},
		})
	}

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
	bot.Cache.Expire(context.Background(), "hangman:"+interaction.ChannelID, availableTime)

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
