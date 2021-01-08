package hangman

import (
	"io/ioutil"
	"log"
	"strings"
)

var words = openWords()

func openWords() []string {
	content, err := ioutil.ReadFile("assets/wordslist_fr.txt")
	if err != nil {
		log.Panicf("‼ Error opening words list for hangman: %s", err.Error())
	}
	return strings.Split(string(content), "\n")
}
