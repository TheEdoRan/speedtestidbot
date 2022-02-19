package bot

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func printInfo(user *tgbotapi.User, text string) {
	fmt.Printf("%v %v [%v]: %v\n", user.FirstName, user.LastName, user.UserName, text)
}

// LogRequest prints information regarding incoming update.
// Only with env var DEBUG set to true.
func LogRequest(update interface{}) {
	if os.Getenv("DEBUG") != "true" {
		return
	}

	switch upd := update.(type) {
	case *tgbotapi.Message:
		if upd.Text == "" {
			return
		}
		printInfo(upd.From, upd.Text)
	case *tgbotapi.InlineQuery:
		printInfo(upd.From, upd.Query)
	case *tgbotapi.ChosenInlineResult:
		printInfo(upd.From, "chosen -> "+upd.ResultID)
	}
}
