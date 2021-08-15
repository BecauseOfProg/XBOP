package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/BecauseOfProg/xbop/commands"
	"github.com/BecauseOfProg/xbop/games/connect_four"
	"github.com/BecauseOfProg/xbop/games/hangman"
	"github.com/BecauseOfProg/xbop/games/irregular_verbs"
	"github.com/BecauseOfProg/xbop/games/tic_tac_toe"
)

func main() {
	bot := onyxcord.RegisterBot("XBOP")
	bot.Commands = map[string]*onyxcord.Command{
		"verbs":   irregular_verbs.Command(),
		"hangman": hangman.Command(),
		"about":   commands.About(),

		"connect-four":          connect_four.Command(),
		"Défier au Puissance 4": connect_four.Command(),

		"tic-tac-toe":       tic_tac_toe.Command(),
		"Défier au morpion": tic_tac_toe.Command(),
	}

	bot.Components = map[string]onyxcord.Component{
		"connectfour": connect_four.HandleOngoingGame,
		"hangman":     hangman.HandleInteraction,
		"tictactoe":   tic_tac_toe.HandleInteraction,
	}

	bot.Client.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		_ = hangman.HandleMessage(&bot, message.Message)
		irregular_verbs.HandleOngoingGame(&bot, message.Message)
	})

	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	bot.Run(true)
}
