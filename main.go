package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/theovidal/105quiz/games/irregular_verbs"
)

func main() {
	bot := onyxcord.RegisterBot("105quiz", false)

	bot.RegisterCommand("verbs", irregular_verbs.Command())
	irregular_verbs.VerbsPlayers.Initialize()

	bot.Client.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		verbsPlayer, exists := irregular_verbs.VerbsPlayers.GetPlayer(message.ChannelID)
		if exists {
			irregular_verbs.HandleAnswer(&bot, message.Message, verbsPlayer)
		}
		bot.OnCommand(session, message)
	})

	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	bot.Run()
}
