package bot

import (
	"fmt"
	"os"
	"regexp"

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
		match := regexp.MustCompile("https?://(.+):8080/speedtest/upload.php")
		trace := match.ReplaceAllString(s.URL, "$1")
		txt := fmt.Sprintf(InlineResponseText, s.Sponsor, s.Name, s.Country, s.ID, trace)
		art := tgbotapi.NewInlineQueryResultArticleHTML(s.ID, s.Sponsor, txt)
		art.Description = s.Name + ", " + s.Country

		kb := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("🔗  Test with this server", "https://speedtest.net/server/"+s.ID),
			),
		)
		art.ReplyMarkup = &kb

		articles = append(articles, art)
	}

	// Cache for 4 hours (in production)
	var cacheTime = 60 * 60 * 4

	// Disable cache in development
	if os.Getenv("DEV") == "true" {
		cacheTime = 0
	}

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: inline.ID,
		IsPersonal:    false,
		CacheTime:     cacheTime,
		Results:       articles,
	}

	if _, err := bot.Request(inlineConf); err != nil {
		fmt.Println("error trying to send inline response:", err)
	}
}
