package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleStartCommand handles /start in chats.
func HandleStartCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	firstName := message.From.FirstName
	txt := fmt.Sprintf(StartCommandReply, tgbotapi.EscapeText(tgbotapi.ModeHTML, firstName))
	m := tgbotapi.NewMessage(message.Chat.ID, txt)
	m.ParseMode = tgbotapi.ModeHTML

	m.ReplyToMessageID = message.MessageID

	if _, err := bot.Send(m); err != nil {
		fmt.Println("error trying to send /start command:", err)
	}
}
