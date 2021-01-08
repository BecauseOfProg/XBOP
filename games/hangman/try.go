package hangman

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/BecauseOfProg/xbop/lib"
)

func Try(bot *onyxcord.Bot, message *discordgo.Message, cacheID string) {
	word := bot.Cache.HGet(context.Background(), cacheID, "word").Val()
	if lib.Contains(lib.StopSentences, lib.TrimNonLetters(message.Content)) {
		bot.Cache.Del(context.Background(), cacheID)
		bot.Client.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf("**:stop_sign: Arrêt de la partie prononcé.**\n\nLe mot à trouver était **%s**.", word),
		)
		return
	}

	trialLetter := strings.ToUpper(string(message.Content[0]))
	bot.Client.ChannelMessageDelete(message.ChannelID, message.ID)

	letters := bot.Cache.HGet(context.Background(), cacheID, "letters").Val()
	falseLetters := bot.Cache.HGet(context.Background(), cacheID, "falseLetters").Val()
	gameMessage := bot.Cache.HGet(context.Background(), cacheID, "message").Val()
	maxErrors, _ := bot.Cache.HGet(context.Background(), cacheID, "maxErrors").Int()

	if strings.Contains(falseLetters+letters, trialLetter) {
		return
	}

	currentWord := hideWord(word, letters)
	trialWord := hideWord(word, letters+trialLetter)

	if currentWord == trialWord {
		falseLetters += trialLetter
	} else {
		letters += trialLetter
	}

	if len(falseLetters) >= maxErrors {
		bot.Cache.Del(context.Background(), cacheID)
		bot.Client.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf("**:cry: C'est perdu ! Retentez votre chance !**\n\nLe mot était **%s** !", word),
		)
		return
	}

	bot.Client.ChannelMessageEdit(message.ChannelID, gameMessage, formatMessage(word, letters, falseLetters, maxErrors))

	if trialWord == word {
		bot.Cache.Del(context.Background(), cacheID)
		bot.Client.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf("**:tada: Bravo ! Vous avez trouvé le mot %s !**", word),
		)
		return
	}

	bot.Cache.HMSet(
		context.Background(), cacheID,
		"letters", letters,
		"falseLetters", falseLetters,
	)
}