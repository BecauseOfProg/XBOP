package irregular_verbs

import (
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/theovidal/105quiz/lib"
)

var VerbsPlayers lib.Players

func Command() *onyxcord.Command {
	return &onyxcord.Command{
		Description:    "Lancer un quiz sur les verbes irréguliers en anglais",
		Show:           true,
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(arguments []string, bot onyxcord.Bot, message *discordgo.MessageCreate) (err error) {
			player := lib.NewClient(message)
			player.Props["answers"] = 0
			player.Props["successfulAnswers"] = 0
			VerbsPlayers.AddPlayer(&player)

			bot.Client.ChannelMessageSend(message.ChannelID, ":flag_gb: **Quiz sur les verbes irréguliers**")
			SendQuestion(&bot, &player)

			return
		},
	}
}
