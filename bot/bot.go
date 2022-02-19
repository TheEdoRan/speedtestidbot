package bot

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Start starts the Telegram bot.
func Start(done <-chan os.Signal) {
	token := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatalln("error creating bot:", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

	updates := bot.GetUpdatesChan(u)

	fmt.Println("bot started")

	go func() {
		<-done
		bot.StopReceivingUpdates()
		fmt.Println("gracefully stopping bot")
	}()

	for update := range updates {
		msg := update.Message
		inline := update.InlineQuery
		chosen := update.ChosenInlineResult

		if msg != nil && msg.Command() == "start" {
			LogRequest(msg)
			HandleStartCommand(bot, msg)
		}

		if inline != nil {
			LogRequest(inline)
			HandleInlineQuery(bot, inline)
		}

		if chosen != nil {
			LogRequest(chosen)
		}
	}
}
