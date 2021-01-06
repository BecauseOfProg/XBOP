package irregular_verbs

import (
	"encoding/csv"
	"log"
	"os"
)

var categories = []string{"la **base verbale**", "le **prétérit**", "le **participe passé**", "la **traduction**"}
var skipSentences = []string{"jepasse", "passe", "passer", "suivant", "skip", "jsp", "jesaispas", "jesaipa", "jesaispa", "jesaipas", "aucuneidee"}
var stopSentences = []string{"stop", "arret", "arreter", "tg", "areter"}

var verbs = OpenVerbs()

func OpenVerbs() [][]string {
	file, err := os.Open("assets/irregular_verbs.csv")
	if err != nil {
		log.Panicln(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Panicln(err)
	}

	return records
}
