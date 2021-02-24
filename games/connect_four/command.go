package connect_four

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			player1 := interaction.Member.User
			player2 := interaction.Data.Options[0].UserValue(bot.Client)

			var columns []string
			for i := 0; i < 7; i++ {
				columns = append(columns, "000000")
				bot.Cache.LPush(context.Background(), "connectfour:"+interaction.ChannelID+"/grid", "000000")
			}

			_ = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: fmt.Sprintf("**:red_circle: %s affronte désormais :yellow_circle: %s au Puissance 4!**", player1.Mention(), player2.Mention()),
				},
			})

			game, _ := bot.Client.ChannelMessageSend(interaction.ChannelID, generateGrid(columns))
			turnMessage := sendTurnMessage(bot, player1, interaction.ChannelID, 1)

			bot.Cache.HMSet(context.Background(), "connectfour:"+interaction.ChannelID,
				"1", player1.ID,
				"2", player2.ID,
				"turn", "1",
				"message", game.ID,
				"turnMessage", turnMessage,
			)

			return
		},
	}
}
