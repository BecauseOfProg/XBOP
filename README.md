<div align="center">
    <img src="assets/xbop.png" alt="xbop logo" width="150"/>
    <h1>XBOP</h1>
    <h3>Play various minigames on Discord</h3>
    <a href="https://discord.com/oauth2/authorize?client_id=796457702666534972&permissions=3136&redirect_uri=https%3A%2F%2Fbecauseofprog.fr&response_type=code&scope=bot%20applications.commands%20applications.commands.update">Invite</a> - <a href="https://becauseofprog.fr">Website</a> - <a href="https://discord.becauseofprog.fr">Discord server</a> - <a href="./LICENSE">License</a>
</div>

**⚠ For the moment, only french is available. _onk onk baguette_**

Use our bot on your server using this [awesome link](https://discord.com/oauth2/authorize?client_id=796457702666534972&permissions=3136&redirect_uri=https%3A%2F%2Fbecauseofprog.fr&response_type=code&scope=bot%20applications.commands%20applications.commands.update)! *no paid features I swear*

To get started, type `/` and see the list of commands. Try playing some games such as connect four -amazing- or hangman -wholesome!-.

## 💻 Development

First, check the following requirements:

- Git, for version control
- Golang 1.15 or higher with go-modules for dependencies
- A running instance of [Redis](https://redis.io/) v5 or higher

Clone the project on your local machine:

```bash
git clone https://github.com/theovidal/bacbot  # HTTP
git clone git@github.com:theovidal/bacbot      # SSH
```

Set up some environment variables described in the [.env.example file](./.env.example), either by adding them in the shell or by creating a .env file at the root of the project.

To run and test the bot, simply use `go run .`. To build an executable, use `go build .`.

## 📜 Credits

- Libraries: [discordgo](https://github.com/bwmarrin/discordgo), [onyxcord](https://github.com/theovidal/onyxcord)
- Maintainer: [Théo Vidal](https://github.com/theovidal)

## 🔐 License

[GNU GPL v3]
