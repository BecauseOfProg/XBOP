package hangman

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func handleAttempt(bot *onyxcord.Bot, message *discordgo.Message, cacheID string) {
	if message.Content[0] == '!' {
		return
	}
	word := bot.Cache.HGet(context.Background(), cacheID, "word").Val()

	attemptLetter := strings.ToUpper(string(message.Content[0]))
	bot.Client.ChannelMessageDelete(message.ChannelID, message.ID)

	letters := bot.Cache.HGet(context.Background(), cacheID, "letters").Val()
	falseLetters := bot.Cache.HGet(context.Background(), cacheID, "falseLetters").Val()
	gameMessage := bot.Cache.HGet(context.Background(), cacheID, "message").Val()
	maxErrors, _ := bot.Cache.HGet(context.Background(), cacheID, "maxErrors").Int()

	if strings.Contains(falseLetters+letters, attemptLetter) {
		return
	}

	currentWord := hideWord(word, letters)
	attemptWord := hideWord(word, letters+attemptLetter)

	if currentWord == attemptWord {
		falseLetters += attemptLetter
	} else {
		letters += attemptLetter
	}

	bot.Client.ChannelMessageEdit(message.ChannelID, gameMessage, formatMessage(word, letters, falseLetters, maxErrors))

	if len(falseLetters) >= maxErrors {
		bot.Cache.Del(context.Background(), cacheID)
		bot.Client.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf("**:cry: C'est perdu ! Retentez votre chance !**\n\nLe mot était **%s** !", word),
		)
		return
	}
	if attemptWord == word {
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
