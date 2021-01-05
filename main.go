package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/theovidal/105quiz/commands"
)

func main() {
	bot := onyxcord.RegisterBot("105quiz")

	bot.RegisterCommand("verbs", commands.IrregularVerbs())

	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	bot.Run()
}
