package hangman

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
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

func stopButton(disabled bool) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Arrêter la partie",
					Style:    discordgo.DangerButton,
					CustomID: "hangman_stop",
					Disabled: disabled,
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
