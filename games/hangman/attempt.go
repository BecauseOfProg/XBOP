package hangman

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func handleAttempt(bot *onyxcord.Bot, message *discordgo.Message, cacheID string) (err error) {
	if message.Content[0] == '!' {
		return
	}
	word := bot.Cache.HGet(context.Background(), cacheID, "word").Val()

	attemptLetter := strings.ToUpper(string(message.Content[0]))
	_ = bot.Client.ChannelMessageDelete(message.ChannelID, message.ID)

	letters := bot.Cache.HGet(context.Background(), cacheID, "letters").Val()
	wrongLetters := bot.Cache.HGet(context.Background(), cacheID, "wrongLetters").Val()
	maxErrors, _ := bot.Cache.HGet(context.Background(), cacheID, "maxErrors").Int()

	if strings.Contains(wrongLetters+letters, attemptLetter) {
		return
	}

	currentWord := hideWord(word, letters)
	attemptWord := hideWord(word, letters+attemptLetter)

	if currentWord == attemptWord {
		wrongLetters += attemptLetter
	} else {
		letters += attemptLetter
	}

	bot.Cache.HMSet(
		context.Background(), cacheID,
		"letters", letters,
		"wrongLetters", wrongLetters,
	)

	if len(wrongLetters) >= maxErrors {
		return stopGame(bot, nil, message.ChannelID, fmt.Sprintf("**:cry: C'est perdu ! Retentez votre chance !**\n\nLe mot était **%s** !", word))
	}
	if attemptWord == word {
		return stopGame(bot, nil, message.ChannelID, fmt.Sprintf("**:tada: Bravo ! Vous avez trouvé le mot %s !**", word))
	} else {
		_ = editMessage(bot, nil, message.ChannelID, word, letters, wrongLetters, maxErrors, defaultMessage, false)
	}

	return
}
