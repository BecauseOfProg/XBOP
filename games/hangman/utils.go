package hangman

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func stopButton(disabled bool) discordgo.Button {
	return discordgo.Button{
		Label:    "Arrêter la partie",
		Style:    discordgo.DangerButton,
		CustomID: "hangman_stop",
		Disabled: disabled,
	}
}

func hideWord(word string, letters string) string {
	regex := regexp.MustCompile(fmt.Sprintf("[^%c%s]", word[0], letters))
	return string(regex.ReplaceAll([]byte(word), []byte("_ ")))
}

func formatMessage(word string, letters string, falseLetters string, errors int) string {
	format := fmt.Sprintf(":arrow_forward: `%s`\nErreurs restantes : %d", hideWord(word, letters), errors-len(falseLetters))
	if falseLetters != "" {
		format += fmt.Sprintf("\nUtilisées : %s", falseLetters)
	}
	return format
}
