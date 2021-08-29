package irregular_verbs

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

// Duration in seconds for the game to automatically stop if nobody interacts
const expireTime = time.Minute * 10

var categories = []string{"la **base verbale**", "le **prétérit**", "le **participe passé**", "la **traduction**"}

var verbsPartOne = openVerbs("1")
var verbsPartTwo = openVerbs("2")

var verbs = map[string][][]string{
	"1":   verbsPartOne,
	"2":   verbsPartTwo,
	"all": mergeVerbs(verbsPartOne, verbsPartTwo),
}

func openVerbs(part string) [][]string {
	file, err := os.Open("assets/irregular_verbs_" + part + ".csv")
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

func mergeVerbs(verbs1, verbs2 [][]string) (verbs [][]string) {
	verbs = verbs1
	for _, verb := range verbs2 {
		verbs = append(verbs, verb)
	}
	return
}
