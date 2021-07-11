package hangman

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

var words = openWords()

func openWords() []string {
	content, err := ioutil.ReadFile("assets/wordslist_fr.txt")
	if err != nil {
		log.Panicf("‼ Error opening words list for hangman: %s", err.Error())
	}
	return strings.Split(string(content), "\n")
}

func hideWord(word, letters string) string {
	regex := regexp.MustCompile(fmt.Sprintf("[^%c%s]", word[0], letters))
	return string(regex.ReplaceAll([]byte(word), []byte("_ ")))
}

func stopButton() []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Arrêter la partie",
					Style:    discordgo.DangerButton,
					CustomID: "hangman_stop",
				},
			},
		},
	}
}

const defaultMessage = "**:chains: Un jeu du pendu est en cours dans le salon!**\nTous les utilisateurs ayant accès au salon peuvent participer. Pour que votre message ne soit pas compté, précédez-le d'un point d'exclamation `!`.\n*La partie est valable pendant 15 minutes après son lancement.*"

func formatMessage(word, letters, wrongLetters string, maxErrors int, message string) string {
	if message == "" {
		message = defaultMessage
	}

	format := fmt.Sprintf("%s\n\n:arrow_forward: `%s`\nErreurs restantes : %d", message, hideWord(word, letters), maxErrors-len(wrongLetters))
	if wrongLetters != "" {
		format += fmt.Sprintf("\nUtilisées : %s", wrongLetters)
	}
	return format
}

func editMessage(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, channelID, word, letters, wrongLetters string, maxErrors int, message string, stop bool) (err error) {
	token := bot.Cache.HGet(context.Background(), "hangman:"+channelID, "game").Val()

	components := stopButton()
	if stop {
		components = []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: fmt.Sprintf("Relancer (%d erreurs autorisées)", maxErrors),
						Style: discordgo.SuccessButton,
						Emoji: discordgo.ComponentEmoji{
							Name: "🔄",
						},
						CustomID: fmt.Sprintf("hangman_restart_%d", maxErrors),
					},
				},
			},
		}
	}

	if interaction == nil {
		_, err = bot.Client.InteractionResponseEdit(bot.Config.Bot.ID, &discordgo.Interaction{Token: token}, &discordgo.WebhookEdit{
			Content:    formatMessage(word, letters, wrongLetters, maxErrors, message),
			Components: components,
		})
	} else {
		err = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content:    formatMessage(word, letters, wrongLetters, maxErrors, message),
				Components: components,
			},
		})
	}
	return
}
