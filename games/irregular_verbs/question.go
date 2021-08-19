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
	part := bot.Cache.HGet(context.Background(), cacheID, "part").Val()

	rand.Seed(time.Now().UnixNano())
	row := rand.Intn(len(verbs[part]))
	questionColumn := rand.Intn(4)
	verbColumn := questionColumn

	for verbColumn == questionColumn {
		verbColumn = rand.Intn(4)
	}

	question := verbs[part][row][questionColumn]
	verb := verbs[part][row][verbColumn]
	bot.Cache.HSet(context.Background(), cacheID, "verb", verb)
	bot.Cache.HIncrBy(context.Background(), cacheID, "answers", 1)
	bot.Cache.HSet(context.Background(), cacheID, "succeeded", "true")

	bot.ChannelMessageSend(
		channel,
		fmt.Sprintf(
			"**#%s** `%s` - Indiquer %s",
			bot.Cache.HGet(context.Background(), cacheID, "answers").Val(),
			question,
			categories[verbColumn]),
	)
}
