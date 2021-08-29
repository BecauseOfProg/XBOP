package main

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"

	"github.com/theovidal/onyxcord"
)

var expiringGames = map[string]string{
	"connectfour": "Le puissance 4",
	"hangman":     "Le jeu du pendu",
	"verbs":       "Le quiz Verbes irréguliers",
	"tic_tac_toe": "Le morpion",
}

func subscribe(bot *onyxcord.Bot) {
	time.Sleep(10 * time.Second)
	pubsub := bot.Cache.Subscribe(context.Background(), "__keyevent@0__:expired")

	for message := range pubsub.Channel() {
		parts := strings.Split(message.Payload, ":")
		gameString, channelID := parts[0], parts[1]
		bot.ChannelMessageSendEmbed(channelID, bot.MakeEmbed(&discordgo.MessageEmbed{
			Color:       bot.Config.Bot.ErrorColor,
			Title:       "⏳ Temps écoulé",
			Description: fmt.Sprintf(" %s a expiré dans ce salon, soit car aucun membre n'a participé ou car il s'agit de la durée maximale!", expiringGames[gameString]),
		}))
	}
}
