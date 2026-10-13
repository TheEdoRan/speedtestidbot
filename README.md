# SpeedtestIDBot

![Go Version](https://img.shields.io/github/go-mod/go-version/theedoran/speedtestidbot)
[![Go Report Card](https://goreportcard.com/badge/github.com/theedoran/speedtestidbot)](https://goreportcard.com/report/github.com/theedoran/speedtestidbot)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A Telegram inline bot that finds the ID of a [Speedtest](https://speedtest.net) server from its name.

**Try it:** [@speedtestidbot](https://t.me/speedtestidbot)

## Usage

The bot works in inline mode, so you can use it in any chat. Type its handle followed by a server name, sponsor, or city:

```
@speedtestidbot Vodafone
```

Pick a server from the list (up to 10 results) and the bot sends a message with:

- sponsor, location, and country
- the server **ID**
- the server host (useful for traceroutes)
- a ready-to-copy CLI command: `speedtest -s <id>`
- a button to run the test in the browser

## Self-hosting

Create a bot with [@BotFather](https://t.me/BotFather), enable inline mode, then copy the env file and set your token:

```sh
cp .env.example .env
```

| Variable    | Description                                             |
| ----------- | ------------------------------------------------------- |
| `BOT_TOKEN` | Telegram bot token (required)                           |
| `DEV`       | `true` disables the 4-hour inline result cache          |
| `DEBUG`     | `true` logs incoming queries and chosen results         |

### Docker

```sh
docker compose up -d --build
```

### From source

Requires Go 1.27+.

```sh
go run .
go test ./...
```

## License

[MIT](LICENSE)
