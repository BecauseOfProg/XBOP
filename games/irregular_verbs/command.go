package irregular_verbs

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

var parts = map[string]string{
	"all": "Les deux parties",
	"1":   "Partie 1 (Décembre 2020)",
	"2":   "Partie 2 (Avril 2021)",
}

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			var part string
			options := interaction.ApplicationCommandData().Options
			if len(options) < 1 {
				part = "all"
			} else {
				part = options[0].StringValue()
			}

			err = bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						":flag_gb: **Quiz sur les verbes irréguliers** - %s\nTous les utilisateurs ayant accès au salon peuvent participer.",
						parts[part],
					),
				},
			})
			if err != nil {
				return
			}

			bot.Cache.HSet(context.Background(), "verbs:"+interaction.ChannelID,
				"answers", 0,
				"part", part,
				"successfulAnswers", 0,
			)

			sendQuestion(bot, interaction.ChannelID)
			return
		},
	}
}
