package connect_four

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func HandleOngoingGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, args []string) error {
	connectFourPlayer := bot.Cache.Exists(context.Background(), "connectfour:"+interaction.ChannelID).Val()
	if connectFourPlayer == 1 {
		switch args[0] {
		case "stop":
			return stopGame(bot, interaction, fmt.Sprintf("Arrêt de la partie prononcé par %s.", interaction.Member.Mention()))
		case "turn":
			return handleTurn(bot, interaction, "connectfour:"+interaction.ChannelID, args)
		}
	}
	return nil
}

func stopGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, reason string) error {
	cacheID := "connectfour:" + interaction.ChannelID
	player1ID := bot.Cache.HGet(context.Background(), cacheID, "1").Val()
	player2ID := bot.Cache.HGet(context.Background(), cacheID, "2").Val()
	if interaction.Member.User.ID != player1ID || interaction.Member.User.ID != player2ID {
		return nil
	}

	columns := bot.Cache.LRange(context.Background(), cacheID+"/columns", 0, -1).Val()
	bot.Cache.Del(context.Background(), cacheID, cacheID+"/columns")

	return bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    fmt.Sprintf("**:stop_sign: %s**\n", reason) + generateGrid(columns),
			Components: components(columns, true),
		},
	})
}
