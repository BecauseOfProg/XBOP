package connect_four

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/BecauseOfProg/xbop/lib"
)

func handleTurn(bot *onyxcord.Bot, message *discordgo.Message, cacheID string) {
	turn := bot.Cache.HGet(context.Background(), cacheID, "turn").Val()
	player := bot.Cache.HGet(context.Background(), cacheID, turn).Val()

	if lib.Contains(lib.StopSentences, lib.TrimNonLetters(message.Content)) {
		bot.Client.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf("**:stop_sign: Arrêt de la partie prononcé par %s**", message.Author.Mention()),
		)
		bot.Cache.Del(context.Background(), cacheID, cacheID+"/grid")
		return
	}
	if message.Author.Bot || player != message.Author.ID {
		return
	}

	turnMessage := bot.Cache.HGet(context.Background(), cacheID, "turnMessage").Val()
	bot.Client.ChannelMessageDelete(message.ChannelID, turnMessage)
	bot.Client.ChannelMessageDelete(message.ChannelID, message.ID)

	column, err := strconv.Atoi(string(message.Content[0]))
	if err != nil || column < 1 || column > 7 {
		bot.Client.ChannelMessageSend(message.ChannelID, "*:arrows_counterclockwise: Rangée invalide.*")
		return
	}

	column -= 1
	grid := bot.Cache.LRange(context.Background(), cacheID+"/grid", 0, -1).Val()
	oldRow := grid[column]
	grid[column] = strings.Replace(grid[column], "0", turn, 1)
	if oldRow == grid[column] {
		bot.Client.ChannelMessageSend(message.ChannelID, "*:arrows_counterclockwise: Rangée invalide.*")
		return
	}

	game := bot.Cache.HGet(context.Background(), cacheID, "message").Val()
	bot.Client.ChannelMessageEdit(message.ChannelID, game, generateGrid(grid))

	var nowTurn int
	if turn == "1" {
		nowTurn = 2
	} else {
		nowTurn = 1
	}
	bot.Cache.HSet(context.Background(), cacheID, "turn", nowTurn)
	bot.Cache.LSet(context.Background(), cacheID+"/grid", int64(column), grid[column])

	nowPlayerID := bot.Cache.HGet(context.Background(), cacheID, strconv.Itoa(nowTurn)).Val()
	nowPlayer, _ := bot.Client.GuildMember(message.GuildID, nowPlayerID)
	turnMessage = sendTurnMessage(bot, nowPlayer.User, message.ChannelID, nowTurn)
	bot.Cache.HSet(context.Background(), cacheID, "turnMessage", turnMessage)
}
