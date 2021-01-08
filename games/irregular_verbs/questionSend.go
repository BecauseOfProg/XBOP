package irregular_verbs

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/theovidal/onyxcord"
)

func sendQuestion(bot *onyxcord.Bot, channel string) {
	cacheID := "verbs:" + channel

	rand.Seed(time.Now().UnixNano())
	row := rand.Intn(len(verbs))
	questionColumn := rand.Intn(4)
	verbColumn := questionColumn

	for verbColumn == questionColumn {
		verbColumn = rand.Intn(4)
	}

	question := verbs[row][questionColumn]
	verb := verbs[row][verbColumn]
	bot.Cache.HSet(context.Background(), cacheID, "verb", verb)
	bot.Cache.HIncrBy(context.Background(), cacheID, "answers", 1)
	bot.Cache.HSet(context.Background(), cacheID, "succeeded", "true")

	bot.Client.ChannelMessageSend(
		channel,
		fmt.Sprintf(
			"**#%s** `%s` - Indiquer %s",
			bot.Cache.HGet(context.Background(), cacheID, "answers").Val(),
			question,
			categories[verbColumn]),
	)
}
