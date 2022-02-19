package bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/theedoran/speedtestidbot/api"
)

// HandleInlineQuery handles user query (search), forwarding text to
// SearchByName (Speedtest API). It displays server info and an inline keyboard
// containing an URL button to test with the chosen server in the browser.
func HandleInlineQuery(bot *tgbotapi.BotAPI, inline *tgbotapi.InlineQuery) {
	servers, err := api.SearchByName(inline.Query)

	if err != nil {
		fmt.Println("error reaching Speedtest api:", err)
		return
	}

	var articles []interface{}

	for _, s := range servers {
		url := strings.Replace(s.URL, ":8080/speedtest/upload.php", "", 1)
		txt := fmt.Sprintf(InlineResponseText, s.Sponsor, s.Name, s.Country, s.ID, url)
		art := tgbotapi.NewInlineQueryResultArticleHTML(s.ID, s.Sponsor, txt)
		art.Description = fmt.Sprintf("%s, %s", s.Name, s.Country)

		kb := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("🌐  Test with this server", fmt.Sprintf("https://speedtest.net/server/%s", s.ID)),
			),
		)
		art.ReplyMarkup = &kb

		articles = append(articles, art)
	}

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: inline.ID,
		IsPersonal:    false,
		CacheTime:     1,
		Results:       articles,
	}

	if _, err := bot.Request(inlineConf); err != nil {
		fmt.Println("error trying to send inline response:", err)
	}
}
