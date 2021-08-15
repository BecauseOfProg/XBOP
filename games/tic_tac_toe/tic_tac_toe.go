package tic_tac_toe

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
			var player2 *discordgo.User
			if interaction.ApplicationCommandData().TargetID != "" {
				member, _ := bot.Client.GuildMember(interaction.GuildID, interaction.ApplicationCommandData().TargetID)
				player2 = member.User
			} else {
				player2 = interaction.ApplicationCommandData().Options[0].UserValue(bot.Client)
			}
			if player2.Bot {
				return bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: ":x: Vous ne pouvez pas jouer avec un robot (en plus c'est super glauque).",
						Flags:   1 << 6,
					},
				})
			}

			return startGame(bot, interaction, player1, player2)
		},
	}
}

func HandleInteraction(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, args []string) (err error) {
	if isGameOngoing := bot.Cache.Exists(context.Background(), "tictactoe:"+interaction.ChannelID).Val(); isGameOngoing == 1 {
		if args[0] == "stop" {
			err = stopGame(bot, interaction, fmt.Sprintf(":stop_sign: Arrêt de la partie prononcé par %s.", interaction.Member.Mention()))
		} else {
			err = handleTurn(bot, interaction, args, "tictactoe:"+interaction.ChannelID)
		}
	}
	switch args[0] {
	case "restart":
		player1, _ := bot.Client.GuildMember(interaction.GuildID, args[1])
		player2, _ := bot.Client.GuildMember(interaction.GuildID, args[2])
		if interaction.Member.User.ID != player1.User.ID && interaction.Member.User.ID != player2.User.ID {
			return nil
		}
		return startGame(bot, interaction, player1.User, player2.User)
	}
	return
}
