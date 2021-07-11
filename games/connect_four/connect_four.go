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
			player2 := interaction.ApplicationCommandData().Options[0].UserValue(bot.Client)
			return startGame(bot, interaction, player1, player2)
		},
	}
}

func HandleOngoingGame(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, args []string) error {
	if isGameOngoing := bot.Cache.Exists(context.Background(), "connectfour:"+interaction.ChannelID).Val(); isGameOngoing == 1 {
		switch args[0] {
		case "stop":
			return stopGame(bot, interaction, fmt.Sprintf(":stop_sign: Arrêt de la partie prononcé par %s.", interaction.Member.Mention()))
		case "turn":
			return handleTurn(bot, interaction, "connectfour:"+interaction.ChannelID, args)
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
	return nil
}
