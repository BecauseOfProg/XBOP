package tic_tac_toe

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			player1 := interaction.Member.User
			player2 := interaction.ApplicationCommandData().Options[0].UserValue(bot.Client)

			var columns []string
			for i := 0; i < 3; i++ {
				columns = append(columns, "000")
				bot.Cache.LPush(context.Background(), "tictactoe:"+interaction.ChannelID+"/columns", "000")
			}

			err = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    generateTurnMessage(player1, 1),
					Components: generateGrid(columns, false),
				},
			})
			if err != nil {
				return
			}

			bot.Cache.HMSet(context.Background(), "tictactoe:"+interaction.ChannelID,
				"1", player1.ID,
				"2", player2.ID,
				"playing", "1",
				"turn", 1,
			)

			return
		},
	}
}
