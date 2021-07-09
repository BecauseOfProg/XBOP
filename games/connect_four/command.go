package connect_four

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			player1 := interaction.Member
			player2 := interaction.ApplicationCommandData().Options[0].UserValue(bot.Client)

			var columns []string
			for i := 0; i < 7; i++ {
				columns = append(columns, "000000")
				bot.Cache.LPush(context.Background(), "connectfour:"+interaction.ChannelID+"/columns", "000000")
			}

			err = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    generateTurnMessage(player1, 1) + generateGrid(columns),
					Components: components(columns, false),
				},
			})
			if err != nil {
				return
			}

			bot.Cache.HSet(context.Background(), "connectfour:"+interaction.ChannelID,
				"1", player1.User.ID,
				"2", player2.ID,
				"playing", "1",
			)

			return
		},
	}
}
