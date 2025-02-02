package irregular_verbs

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/BecauseOfProg/xbop/lib"
)

func handleAnswer(bot *onyxcord.Bot, message *discordgo.Message, cacheID string) {
	if message.Author.Bot {
		return
	}
	bot.Cache.Expire(context.Background(), cacheID, expireTime)
	trial := lib.TrimNonLetters(message.Content)
	verb := bot.Cache.HGet(context.Background(), cacheID, "verb").Val()

	if lib.Contains(lib.StopSentences, trial) {
		bot.Cache.HIncrBy(context.Background(), cacheID, "answers", -1)
		successfulAnswers, _ := bot.Cache.HGet(context.Background(), cacheID, "successfulAnswers").Int()
		answers, _ := bot.Cache.HGet(context.Background(), cacheID, "answers").Int()

		bot.ChannelMessageSend(
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

	if lib.Contains(lib.SkipSentences, trial) {
		bot.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf(":fast_forward: Le mot recherché était **%s**", verb),
		)
		sendQuestion(bot, message.ChannelID)
		return
	}

	if trial == lib.TrimNonLetters(verb) {
		bot.MessageReactionAdd(message.ChannelID, message.ID, "✅")
		if bot.Cache.HGet(context.Background(), cacheID, "succeeded").Val() == "true" {
			bot.Cache.HIncrBy(context.Background(), cacheID, "successfulAnswers", 1)
		}
		sendQuestion(bot, message.ChannelID)
	} else {
		bot.MessageReactionAdd(message.ChannelID, message.ID, "❌")
		bot.Cache.HSet(context.Background(), cacheID, "suceeded", "false")
	}
}
