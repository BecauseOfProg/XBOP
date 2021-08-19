package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

func About() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			return bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						bot.MakeEmbed(&discordgo.MessageEmbed{
							Title: "ℹ À propos",
							Thumbnail: &discordgo.MessageEmbedThumbnail{
								URL: "https://cdn.becauseofprog.fr/v2/sites/becauseofprog.fr/assets/logos/bop.png",
							},
							Description: fmt.Sprintf("**XBOP** est un robot Discord créé par la **BecauseOfProg** et proposant des jeux variés, solo, en duel ou multijoueur. [Invitez-le](%s) sur votre serveur pour en profiter avec encore plus de monde!\n"+
								"Son code source est ouvert à tous : n'hésitez-pas à contribuer à son développement !\n\n"+

								"🔨 Version : %s\n"+
								"💻 Développeur : [%s](%s)\n\n"+

								"© 2020-présent, BecauseOfProg. Sous licence [GNU GPL v3](%s)",
								bot.Config.Bot.InviteLink,
								bot.Config.Dev.Version,
								bot.Config.Dev.Maintainer.Name,
								bot.Config.Dev.Maintainer.Link,
								bot.Config.Bot.License,
							),
						}),
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Style: discordgo.LinkButton,
									Label: "Site Internet",
									Emoji: discordgo.ComponentEmoji{
										Name: "🌐",
									},
									URL: bot.Config.Bot.Website,
								},
								discordgo.Button{
									Style: discordgo.LinkButton,
									Label: "Serveur Discord",
									Emoji: discordgo.ComponentEmoji{
										Name: "💬",
									},
									URL: bot.Config.Bot.Server,
								},
							},
						},
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Style: discordgo.LinkButton,
									Label: "Supporter financièrement le projet",
									Emoji: discordgo.ComponentEmoji{
										Name: "❤",
									},
									URL: bot.Config.Bot.Donate,
								},
							},
						},
					},
				},
			})
		},
	}
}
