package connect_four

import (
	"context"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		Description:    "Lancer une partie de Puissance 4",
		Usage:          "connectfour <@opponent>",
		Category:       "classic",
		Show:           true,
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(arguments []string, bot onyxcord.Bot, message *discordgo.MessageCreate) (err error) {
			if len(message.Mentions) == 0 {
				return errors.New("Vous devez mentionner la personne avec qui vous souhaitez jouer")
			}

			player1 := message.Author
			player2 := message.Mentions[0]

			var columns []string
			for i := 0; i < 7; i++ {
				columns = append(columns, "000000")
				bot.Cache.LPush(context.Background(), "connectfour:"+message.ChannelID+"/grid", "000000")
			}

			game, _ := bot.Client.ChannelMessageSend(message.ChannelID, generateGrid(columns))
			turnMessage := sendTurnMessage(&bot, player1, message.ChannelID, 1)

			bot.Cache.HMSet(context.Background(), "connectfour:"+message.ChannelID,
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
