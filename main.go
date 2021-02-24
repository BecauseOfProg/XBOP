package main

import (
	"github.com/BecauseOfProg/xbop/games/connect_four"
	"github.com/BecauseOfProg/xbop/games/hangman"
	"github.com/BecauseOfProg/xbop/games/irregular_verbs"
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func main() {
	bot := onyxcord.RegisterBot("XBOP")

	bot.RegisterCommand("verbs", irregular_verbs.Command())

	bot.RegisterCommand("connect-four", connect_four.Command())
	bot.RegisterCommand("hangman", hangman.Command())

	bot.Client.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		connect_four.HandleInteraction(&bot, message.Message)
		hangman.HandleInteraction(&bot, message.Message)
		irregular_verbs.HandleInteraction(&bot, message.Message)
	})

	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	bot.Run(true)
}
