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
		"verbs":        irregular_verbs.Command(),
		"connect-four": connect_four.Command(),
		"hangman":      hangman.Command(),
		"tic-tac-toe":  tic_tac_toe.Command(),
		"about":        commands.About(),
	}

	bot.Components = map[string]onyxcord.Component{
		"connectfour": connect_four.StopGame,
		"hangman":     hangman.StopGame,
		"tictactoe":   tic_tac_toe.HandleOngoingGame,
	}

	bot.Client.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		connect_four.HandleOngoingGame(&bot, message.Message)
		hangman.HandleOngoingGame(&bot, message.Message)
		irregular_verbs.HandleOngoingGame(&bot, message.Message)
	})

	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	bot.Run(true)
}
