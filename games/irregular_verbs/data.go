package irregular_verbs

import (
	"encoding/csv"
	"log"
	"os"
)

var categories = []string{"la **base verbale**", "le **prétérit**", "le **participe passé**", "la **traduction**"}

var verbs = openVerbs()

func openVerbs() [][]string {
	file, err := os.Open("assets/irregular_verbs.csv")
	if err != nil {
		log.Panicf("‼ Error opening verbs list for irregular verbs quiz: %s", err.Error())
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Panicf("‼ Error opening verbs list for irregular verbs quiz: %s", err.Error())
	}

	return records
}
