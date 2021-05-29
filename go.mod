module github.com/BecauseOfProg/xbop

go 1.16

require (
	github.com/bwmarrin/discordgo v0.22.0
	github.com/theovidal/onyxcord v0.1.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	go.mongodb.org/mongo-driver v1.4.0 // indirect
	golang.org/x/text v0.3.3
)

replace github.com/theovidal/onyxcord => ../../theovidal/onyxcord

replace github.com/bwmarrin/discordgo => github.com/FedorLap2006/discordgo v0.22.1-0.20210526221316-e7fb87fa3c1b
