package main

import "github.com/bwmarrin/discordgo"

func discordCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		// --------------------------------
		//          INFORMATION
		// --------------------------------
		{
			Name:        "about",
			Type:        discordgo.ChatApplicationCommand,
			Description: "Informations pratiques sur le robot",
		},
		// --------------------------------
		//            GAMES
		// --------------------------------
		{
			Name:        "connect-four",
			Type:        discordgo.ChatApplicationCommand,
			Description: "Jouer au morpion contre un membre du serveur",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "opponent",
					Description: "L'utilisateur à affronter",
					Required:    true,
				},
			},
		},
		{
			Name: "Défier au morpion",
			Type: discordgo.UserApplicationCommand,
		},
		{
			Name:        "tic-tac-toe",
			Type:        discordgo.ChatApplicationCommand,
			Description: "Jouer au Puissance 4 contre un membre du serveur",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "opponent",
					Description: "L'utilisateur à affronter",
					Required:    true,
				},
			},
		},
		{
			Name: "Défier au Puissance 4",
			Type: discordgo.UserApplicationCommand,
		},
		{
			Name:        "hangman",
			Type:        discordgo.ChatApplicationCommand,
			Description: "Jouer au jeu du pendu",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "max-errors",
					Description: "Nombre maximum d'erreurs autorisées",
				},
			},
		},
		{
			Name:        "verbs",
			Type:        discordgo.ChatApplicationCommand,
			Description: "Lancer un quiz sur les verbes irréguliers",
		},
	}
}
