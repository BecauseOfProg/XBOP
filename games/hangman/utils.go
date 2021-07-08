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

func hideWord(word string, letters string) string {
	regex := regexp.MustCompile(fmt.Sprintf("[^%c%s]", word[0], letters))
	return string(regex.ReplaceAll([]byte(word), []byte("_ ")))
}

func stopButton(disabled bool) discordgo.Button {
	return discordgo.Button{
		Label:    "Arrêter la partie",
		Style:    discordgo.DangerButton,
		CustomID: "hangman_stop",
		Disabled: disabled,
	}
}

func formatMessage(word string, letters string, falseLetters string, errors int) string {
	format := fmt.Sprintf(":arrow_forward: `%s`\nErreurs restantes : %d", hideWord(word, letters), errors-len(falseLetters))
	if falseLetters != "" {
		format += fmt.Sprintf("\nUtilisées : %s", falseLetters)
	}
	return format
}
