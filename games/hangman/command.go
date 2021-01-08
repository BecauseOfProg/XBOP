package hangman

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		Description:    "Lancer une partie de pendu",
		Usage:          "hangman <max_errors:int>",
		Category:       "classic",
		Show:           true,
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(arguments []string, bot onyxcord.Bot, message *discordgo.MessageCreate) (err error) {
			rand.Seed(time.Now().UnixNano())
			word := strings.TrimRight(words[rand.Intn(len(words))], "\r")

			var maxErrors int
			if arguments[0] == "" {
				maxErrors = 7
			} else {
				maxErrors, err = strconv.Atoi(arguments[0])
				if err != nil || maxErrors < 1 {
					return errors.New("Le nombre d'erreurs doit être un nombre entier supérieur à 0")
				}
			}

			letters := string(word[0])
			game, _ := bot.Client.ChannelMessageSend(message.ChannelID, formatMessage(word, letters, "", maxErrors))

			bot.Cache.HMSet(context.Background(), "hangman:"+message.ChannelID,
				"word", word,
				"letters", letters,
				"falseLetters", "",
				"maxErrors", maxErrors,
				"message", game.ID,
			)

			return
		},
	}
}
