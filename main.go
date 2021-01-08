package main

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"

	"github.com/BecauseOfProg/xbop/games/connect_four"
	"github.com/BecauseOfProg/xbop/games/hangman"
	"github.com/BecauseOfProg/xbop/games/irregular_verbs"
)

func main() {
	bot := onyxcord.RegisterBot("XBOP", false)

	bot.RegisterCommand("verbs", irregular_verbs.Command())

	bot.RegisterCommand("connectfour", connect_four.Command())
	bot.RegisterCommand("hangman", hangman.Command())

	bot.Client.AddHandler(func(session *discordgo.Session, event *discordgo.Event) {
		if event.Type == "INTERACTION_CREATE" {
			data := make(map[string]interface{})
			json.Unmarshal(event.RawData, &data)
			command := data["data"].(map[string]interface{})["name"].(string)
			if command == "verbs" {
				channelID := data["channel_id"].(string)
				event := discordgo.MessageCreate{
					Message: &discordgo.Message{
						ChannelID: channelID,
						Author:    &discordgo.User{},
					},
				}
				irregular_verbs.Command().Execute([]string{}, bot, &event)
			}
		}
	})

	bot.Client.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		connect_four.HandleInteraction(&bot, message.Message)
		hangman.HandleInteraction(&bot, message.Message)
		irregular_verbs.HandleInteraction(&bot, message.Message)
		bot.OnCommand(session, message)
	})

	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	bot.Run()
}
