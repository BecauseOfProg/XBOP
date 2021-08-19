package connect_four

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func handleTurn(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, cacheID string, args []string) (err error) {
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

	if playingMember.User.ID != interaction.Member.User.ID {
		token, _ := strconv.Atoi(playingIndex)
		editMessage(bot, interaction, playingMember, token, columns)
		return
	}

	columnIndex, _ := strconv.Atoi(args[1])

	columns[columnIndex] = strings.Replace(columns[columnIndex], "0", playingIndex, 1)
	rowIndex := 0
	for i, char := range columns[columnIndex] {
		if string(char) == "0" {
			rowIndex = i - 1
			break
		}
	}

	bot.Cache.HSet(context.Background(), cacheID, "playing", waitingIndex)
	bot.Cache.LSet(context.Background(), cacheID+"/columns", int64(columnIndex), columns[columnIndex])

	if isVictorious(columns, columnIndex, rowIndex) {
		fmt.Println("putain")
		return stopGame(bot, interaction, fmt.Sprintf(":tada: %s remporte la partie!", playingMember.Mention()))
	}

	editMessage(bot, interaction, waitingMember, waitingIndex, columns)
	return
}

func editMessage(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, player *discordgo.Member, token int, columns []string) {
	bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    generateTurnMessage(player.User, token) + generateGrid(columns),
			Components: components(columns),
		},
	})
}

var directions = [][]int{
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
}

func isVictorious(columns []string, column, row int) bool {
	columnDiff := 0
	rowDiff := 0
	token := string(columns[column][row])

	directionIndex := 0
	directionCounts := []int{1, 1, 1, 1}

	for {
		if directionIndex == 8 {
			return false
		}

		direction := directions[directionIndex]

		columnDiff += direction[0]
		rowDiff += direction[1]

		nextColumn := column + columnDiff
		nextRow := row + rowDiff

		if nextColumn < 0 || nextColumn > 6 || nextRow < 0 || nextRow > 5 {
			directionIndex += 1
			columnDiff = 0
			rowDiff = 0
			continue
		}
		nextToken := string(columns[column+columnDiff][row+rowDiff])

		if nextToken == token {
			directionCounts[directionIndex%4] += 1
		} else {
			directionIndex += 1
			columnDiff = 0
			rowDiff = 0
			continue
		}

		for _, count := range directionCounts {
			if count >= 4 {
				return true
			}
		}
	}
}
