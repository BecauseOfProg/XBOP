package hangman

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			maxErrors := 7
			if options := interaction.ApplicationCommandData().Options; len(options) != 0 {
				maxErrors = int(options[0].IntValue())
			}

			return startGame(bot, interaction, maxErrors)
		},
	}
}

func HandleMessage(bot *onyxcord.Bot, message *discordgo.Message) error {
	if isGameOngoing := bot.Cache.Exists(context.Background(), "hangman:"+message.ChannelID).Val(); isGameOngoing == 1 {
		return handleAttempt(bot, message, "hangman:"+message.ChannelID)
	}
	return nil
}

func HandleInteraction(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate, args []string) (err error) {
	if isGameOngoing := bot.Cache.Exists(context.Background(), "hangman:"+interaction.ChannelID).Val(); isGameOngoing == 1 {
		switch args[0] {
		case "stop":
			word := bot.Cache.HGet(context.Background(), "hangman:"+interaction.ChannelID, "word").Val()
			return stopGame(bot, interaction, interaction.ChannelID, fmt.Sprintf("**:stop_sign: Arrêt de la partie prononcé par %s.**\n\nLe mot à trouver était **%s**.", interaction.Member.Mention(), word))
		}
	}
	switch args[0] {
	case "restart":
		if maxErrors, convErr := strconv.Atoi(args[1]); convErr != nil {
			return convErr
		} else {
			return startGame(bot, interaction, maxErrors)
		}
	}
	return
}
