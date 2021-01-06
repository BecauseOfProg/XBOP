package lib

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Client struct {
	Context *discordgo.MessageCreate
	Props   map[string]interface{}
}

func NewClient(context *discordgo.MessageCreate) Client {
	return Client{
		Context: context,
		Props:   make(map[string]interface{}),
	}
}

type Players struct {
	sync.RWMutex
	clients map[string]*Client
}

func (p *Players) Initialize() {
	p.clients = make(map[string]*Client)
}

func (p *Players) AddPlayer(client *Client) {
	p.Lock()
	p.clients[client.Context.ChannelID] = client
	p.Unlock()
}

func (p *Players) RemovePlayer(channel string) {
	p.Lock()
	delete(p.clients, channel)
	p.Unlock()
}

func (p *Players) GetPlayer(channel string) (player *Client, exists bool) {
	player, exists = p.clients[channel]
	return
}
