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
			return stopGame(bot, interaction)
		case "turn":
			handleTurn(bot, interaction, "connectfour:"+interaction.ChannelID)
		}
	}
	return nil
}

func stopGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) error {
	cacheID := "connectfour:" + interaction.ChannelID
	columns := bot.Cache.LRange(context.Background(), cacheID+"/columns", 0, -1).Val()
	bot.Cache.Del(context.Background(), cacheID, cacheID+"/columns")

	return bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    fmt.Sprintf("**:stop_sign: Arrêt de la partie prononcé par %s.**", interaction.Member.User.Mention()) + generateGrid(columns),
			Components: components(columns, true),
		},
	})
}
