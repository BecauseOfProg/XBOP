package irregular_verbs

import (
	"fmt"
	"github.com/theovidal/onyxcord"
	"math/rand"
	"time"

	"github.com/theovidal/105quiz/lib"
)

func SendQuestion(bot *onyxcord.Bot, player *lib.Client) {
	rand.Seed(time.Now().UnixNano())
	row := rand.Intn(len(verbs))
	questionColumn := rand.Intn(4)
	verbColumn := questionColumn

	for verbColumn == questionColumn {
		verbColumn = rand.Intn(4)
	}

	question := verbs[row][questionColumn]
	verb := verbs[row][verbColumn]
	player.Props["verb"] = verb

	player.Props["answers"] = player.Props["answers"].(int) + 1
	player.Props["succeeded"] = true
	bot.Client.ChannelMessageSend(
		player.Context.ChannelID,
		fmt.Sprintf("**#%d** `%s` - Indiquer %s", player.Props["answers"], question, categories[verbColumn]),
	)
}
