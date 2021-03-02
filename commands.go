package main

import "github.com/bwmarrin/discordgo"

func commands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "connect-four",
			Description: "Affronter un utilisateur au Puissance 4",
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
			Name:        "hangman",
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
			Description: "Lancer un quiz sur les verbes irréguliers",
		},
	}
}
