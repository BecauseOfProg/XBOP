package connect_four

import (
	"context"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func handleTurn(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, cacheID string) {
	playingIndex := bot.Cache.HGet(context.Background(), cacheID, "playing").Val()
	playingUserID := bot.Cache.HGet(context.Background(), cacheID, playingIndex).Val()
	playingMember, _ := bot.Client.GuildMember(interaction.GuildID, playingUserID)

	waitingIndex := 1
	if playingIndex == "1" {
		waitingIndex = 2
	}

	waitingUserID := bot.Cache.HGet(context.Background(), cacheID, strconv.Itoa(waitingIndex)).Val()
	waitingMember, _ := bot.Client.GuildMember(interaction.GuildID, waitingUserID)

	columns := bot.Cache.LRange(context.Background(), cacheID+"/columns", 0, -1).Val()

	if playingMember.User.ID != interaction.Member.User.ID {
		token, _ := strconv.Atoi(playingIndex)
		editMessage(bot, interaction, playingMember.User, token, columns)
		return
	}

	columnIndex, _ := strconv.Atoi(interaction.MessageComponentData().Values[0])

	columns[columnIndex] = strings.Replace(columns[columnIndex], "0", playingIndex, 1)

	bot.Cache.HSet(context.Background(), cacheID, "playing", waitingIndex)
	bot.Cache.LSet(context.Background(), cacheID+"/columns", int64(columnIndex), columns[columnIndex])

	editMessage(bot, interaction, waitingMember.User, waitingIndex, columns)
}

func editMessage(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, player *discordgo.User, token int, columns []string) {
	bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    generateTurnMessage(player, token) + generateGrid(columns),
			Components: components(columns, false),
		},
	})
}
