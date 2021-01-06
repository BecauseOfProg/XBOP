package irregular_verbs

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/BecauseOfProg/xbop/lib"
)

func HandleAnswer(bot *onyxcord.Bot, message *discordgo.Message, cacheID string) {
	if message.Author.Bot {
		return
	}
	trial := lib.TrimNonLetters(message.Content)
	verb := bot.Cache.HGet(context.Background(), cacheID, "verb").Val()

	if lib.Contains(stopSentences, trial) {
		bot.Cache.HIncrBy(context.Background(), cacheID, "answers", -1)
		successfulAnswers, _ := bot.Cache.HGet(context.Background(), cacheID, "successfulAnswers").Int()
		answers, _ := bot.Cache.HGet(context.Background(), cacheID, "answers").Int()

		bot.Client.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf(
				":stop_sign: **Arrêt du quiz en cours!** Vous avez réussi %d questions sur %d (note de %.2f/20)",
				successfulAnswers,
				answers,
				(float64(successfulAnswers)/float64(answers))*20.0,
			),
		)
		bot.Cache.Del(context.Background(), cacheID)
		return
	}

	if lib.Contains(skipSentences, trial) {
		bot.Client.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf(":fast_forward: Le mot recherché était **%s**", verb),
		)
		SendQuestion(bot, message.ChannelID)
		return
	}

	if trial == lib.TrimNonLetters(verb) {
		bot.Client.MessageReactionAdd(message.ChannelID, message.ID, "✅")
		if bot.Cache.HGet(context.Background(), cacheID, "succeeded").Val() == "true" {
			bot.Cache.HIncrBy(context.Background(), cacheID, "successfulAnswers", 1)
		}
		SendQuestion(bot, message.ChannelID)
	} else {
		bot.Client.MessageReactionAdd(message.ChannelID, message.ID, "❌")
		bot.Cache.HSet(context.Background(), cacheID, "suceeded", "false")
	}
}
