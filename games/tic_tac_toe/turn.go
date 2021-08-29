package tic_tac_toe

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func handleTurn(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, args []string, cacheID string) (err error) {
	bot.Cache.Expire(context.Background(), cacheID, expireTime)
	playingIndex := bot.Cache.HGet(context.Background(), cacheID, "playing").Val()
	playingUserID := bot.Cache.HGet(context.Background(), cacheID, playingIndex).Val()
	playingMember, _ := bot.GuildMember(interaction.GuildID, playingUserID)

	waitingIndex := 1
	if playingIndex == "1" {
		waitingIndex = 2
	}
	waitingUserID := bot.Cache.HGet(context.Background(), cacheID, strconv.Itoa(waitingIndex)).Val()
	waitingMember, _ := bot.GuildMember(interaction.GuildID, waitingUserID)

	columns := bot.Cache.LRange(context.Background(), cacheID+"/columns", 0, -1).Val()

	columnIndex, _ := strconv.Atoi(args[0])
	rowIndex, _ := strconv.Atoi(args[1])
	column := columns[columnIndex]
	if column[rowIndex] != '0' || playingUserID != interaction.Member.User.ID {
		token, _ := strconv.Atoi(playingIndex)
		editMessage(bot, interaction, playingMember, token, columns)
		return
	}

	columns[columnIndex] = column[:rowIndex] + playingIndex + column[rowIndex+1:]
	bot.Cache.HSet(context.Background(), cacheID, "playing", waitingIndex)
	bot.Cache.LSet(context.Background(), cacheID+"/columns", int64(columnIndex), columns[columnIndex])

	for _, config := range winningGrids {
		k := 0
		for _, pos := range config {
			if string(columns[pos[0]][pos[1]]) == playingIndex {
				k++
			}
		}
		if k == 3 {
			return stopGame(bot, interaction, fmt.Sprintf(":tada: %s remporte la partie!", playingMember.User.Mention()))
		}
	}

	if turn, _ := bot.Cache.HGet(context.Background(), cacheID, "turn").Int(); turn > 8 {
		return stopGame(bot, interaction, "La partie n'a pas de gagnant.")
	} else {
		bot.Cache.HSet(context.Background(), cacheID, "turn", turn+1)
	}

	editMessage(bot, interaction, waitingMember, waitingIndex, columns)

	return
}

func editMessage(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, player *discordgo.Member, token int, columns []string) {
	bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    generateTurnMessage(player.User, token),
			Components: generateGrid(columns, false),
		},
	})
}
