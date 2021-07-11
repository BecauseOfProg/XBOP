package connect_four

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func startGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, player1, player2 *discordgo.User) (err error) {
	var columns []string
	for i := 0; i < 7; i++ {
		columns = append(columns, "000000")
		bot.Cache.LPush(context.Background(), "connectfour:"+interaction.ChannelID+"/columns", "000000")
	}

	err = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    generateTurnMessage(player1, 1) + generateGrid(columns),
			Components: components(columns),
		},
	})
	if err != nil {
		return
	}

	bot.Cache.HSet(context.Background(), "connectfour:"+interaction.ChannelID,
		"1", player1.ID,
		"2", player2.ID,
		"playing", "1",
	)

	return
}

func stopGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, reason string) error {
	cacheID := "connectfour:" + interaction.ChannelID
	player1ID := bot.Cache.HGet(context.Background(), cacheID, "1").Val()
	player2ID := bot.Cache.HGet(context.Background(), cacheID, "2").Val()
	if interaction.Member.User.ID != player1ID && interaction.Member.User.ID != player2ID {
		return nil
	}

	columns := bot.Cache.LRange(context.Background(), cacheID+"/columns", 0, -1).Val()
	bot.Cache.Del(context.Background(), cacheID, cacheID+"/columns")

	return bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("**%s**\n", reason) + generateGrid(columns),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Relancer (mêmes adversaires)",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🔄",
							},
							CustomID: fmt.Sprintf("connectfour_restart_%s_%s", player2ID, player1ID),
						},
					},
				},
			},
		},
	})
}
