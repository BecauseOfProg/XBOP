package tic_tac_toe

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func HandleOngoingGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, args []string) (err error) {
	connectFourPlayer := bot.Cache.Exists(context.Background(), "tictactoe:"+interaction.ChannelID).Val()
	if connectFourPlayer == 1 {
		if args[0] == "stop" {
			err = stopGame(bot, interaction, fmt.Sprintf("L'arrêt de la partie a été prononcé par %s.", interaction.Member.Mention()))
		} else {
			err = handleTurn(bot, interaction, args, "tictactoe:"+interaction.ChannelID)
		}
	}
	return
}

func stopGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, reason string) error {
	cacheID := "tictactoe:" + interaction.ChannelID
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
			Content:    fmt.Sprintf("**:stop_sign: %s**", reason),
			Components: generateGrid(columns, true),
		},
	})
}
