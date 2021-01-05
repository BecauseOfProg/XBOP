package commands

import (
	"log"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func trimNonLetters(input string) (output string) {
	output = strings.ToLower(input)
	output = trimAccents(input)

	regex := regexp.MustCompile(`[^a-z]`)
	output = string(regex.ReplaceAll([]byte(output), []byte{}))

	return
}

func trimAccents(input string) string {
	transformer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(transformer, input)
	if err != nil {
		log.Panicln(err)
	}

	return output
}

// Checks if a specific slice contains a string
func contains(slice []string, text string) bool {
	for _, item := range slice {
		if item == text {
			return true
		}
	}

	return false
}
