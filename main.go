package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/BecauseOfProg/xbop/commands"
	"github.com/BecauseOfProg/xbop/games/connect_four"
	"github.com/BecauseOfProg/xbop/games/hangman"
	"github.com/BecauseOfProg/xbop/games/irregular_verbs"
)

func main() {
	bot := onyxcord.RegisterBot("XBOP")

	bot.RegisterCommand("verbs", irregular_verbs.Command())

	bot.RegisterCommand("connect-four", connect_four.Command())
	bot.RegisterCommand("hangman", hangman.Command())

	bot.RegisterCommand("about", commands.About())

	bot.Client.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		connect_four.HandleOngoingGame(&bot, message.Message)
		hangman.HandleOngoingGame(&bot, message.Message)
		irregular_verbs.HandleOngoingGame(&bot, message.Message)
	})

	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	bot.Run(true)
}
