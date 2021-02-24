module github.com/BecauseOfProg/xbop

go 1.15

require (
	github.com/bwmarrin/discordgo v0.22.0
	github.com/theovidal/onyxcord v0.1.0
	golang.org/x/text v0.3.3
)

replace github.com/theovidal/onyxcord => ../../theovidal/onyxcord

replace github.com/bwmarrin/discordgo => github.com/FedorLap2006/discordgo v0.22.1-0.20210217184539-8718e2d37898
