package commands

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

var categories = []string{"la **base verbale**", "le **prétérit**", "le **participe passé**", "la **traduction**"}
var skipSentences = []string{"jepasse", "passe", "passer", "suivant", "skip"}
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

// TODO: create a centralized handler that stores all games so the bot doesn't freeze
func IrregularVerbs() *onyxcord.Command {
	return &onyxcord.Command{
		Description:    "Lancer un quiz sur les verbes irréguliers en anglais",
		Show:           true,
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(arguments []string, bot onyxcord.Bot, message *discordgo.MessageCreate) (err error) {
			channel := make(chan *discordgo.Message)

			bot.Client.ChannelMessageSend(message.ChannelID, ":flag_gb: **Quiz sur les verbes irréguliers**")
			answers := 0
			successfulAnswers := 0

		game:
			for {
				rand.Seed(time.Now().UnixNano())
				row := rand.Intn(len(verbs))
				knownColumn := rand.Intn(4)
				unknownColumn := knownColumn

				for unknownColumn == knownColumn {
					unknownColumn = rand.Intn(4)
				}

				knownVerb := verbs[row][knownColumn]
				unknownVerb := verbs[row][unknownColumn]

				answers += 1
				succeeded := true
				bot.Client.ChannelMessageSend(
					message.ChannelID,
					fmt.Sprintf("**#%d** `%s` - Indiquer %s", answers, knownVerb, categories[unknownColumn]),
				)

			question:
				for {
					bot.Client.AddHandlerOnce(func(_ *discordgo.Session, event *discordgo.MessageCreate) {
						if event.ChannelID != message.ChannelID {
							event.Author.Bot = true
						}
						channel <- event.Message
					})
					msg := <-channel
					if msg.Author.Bot {
						continue question
					}
					trial := trimNonLetters(msg.Content)

					if contains(stopSentences, trial) {
						answers -= 1
						bot.Client.ChannelMessageSend(
							msg.ChannelID,
							fmt.Sprintf(
								":stop_sign: **Arrêt du quiz en cours!** Vous avez réussi %d questions sur %d (note de %.2f/20)",
								successfulAnswers,
								answers,
								(float64(successfulAnswers)/float64(answers))*20.0,
							),
						)
						break game
					}

					if contains(skipSentences, trial) {
						bot.Client.ChannelMessageSend(
							msg.ChannelID,
							fmt.Sprintf(":fast_forward: Le mot recherché était **%s**", unknownVerb),
						)
						continue game
					}

					if trial == trimNonLetters(unknownVerb) {
						bot.Client.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
						if succeeded {
							successfulAnswers += 1
						}
						break question
					} else {
						bot.Client.MessageReactionAdd(msg.ChannelID, msg.ID, "❌")
						succeeded = false
					}
				}
			}
			return nil
		},
	}
}
