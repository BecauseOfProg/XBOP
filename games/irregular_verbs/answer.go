package irregular_verbs

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/105quiz/lib"
	"github.com/theovidal/onyxcord"
)

func HandleAnswer(bot *onyxcord.Bot, message *discordgo.Message, player *lib.Client) {
	if message.Author.Bot {
		return
	}
	trial := lib.TrimNonLetters(message.Content)

	if lib.Contains(stopSentences, trial) {
		player.Props["answers"] = player.Props["answers"].(int) - 1
		bot.Client.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf(
				":stop_sign: **Arrêt du quiz en cours!** Vous avez réussi %d questions sur %d (note de %.2f/20)",
				player.Props["successfulAnswers"].(int),
				player.Props["answers"].(int),
				(float64(player.Props["successfulAnswers"].(int))/float64(player.Props["answers"].(int)))*20.0,
			),
		)
		VerbsPlayers.RemovePlayer(message.ChannelID)
		return
	}

	if lib.Contains(skipSentences, trial) {
		bot.Client.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf(":fast_forward: Le mot recherché était **%s**", player.Props["unknownVerb"]),
		)
		SendQuestion(bot, player)
		return
	}

	println(trial, lib.TrimNonLetters(player.Props["unknownVerb"].(string)))

	if trial == lib.TrimNonLetters(player.Props["unknownVerb"].(string)) {
		bot.Client.MessageReactionAdd(message.ChannelID, message.ID, "✅")
		if player.Props["succeeded"].(bool) {
			player.Props["successfulAnswers"] = player.Props["successfulAnswers"].(int) + 1
		}
		SendQuestion(bot, player)
	} else {
		bot.Client.MessageReactionAdd(message.ChannelID, message.ID, "❌")
		player.Props["succeeded"] = false
	}
}
