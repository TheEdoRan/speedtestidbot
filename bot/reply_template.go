package bot

// StartCommandReply is the template for /start command.
// Params:
// 1. user first name
const StartCommandReply = "Hi <b>%s</b>! 👋\n" +
	"You can invoke me in any chat you want, providing the name of the Speedtest server you're interested to.\n\n" +
	"Example:\n" +
	"<pre>@speedtestidbot Vodafone</pre>\n\n" +
	"Then, you just have to select an item from the list, and you'll get ID and other useful data back.\n\n" +
	"This bot is licensed under <a href=\"https://www.gnu.org/licenses/gpl-3.0.html\">GNU GPLv3</a>.\n" +
	"Source code can be found <a href=\"https://github.com/theedoran/speedtestidbot\">here</a>."

// InlineResponseText is the template for inline queries replies (chosen server).
// Params:
// 1: Server sponsor
// 2: Server name
// 3: Server country
// 4: Server ID
// 5: Server URL (without port)
const InlineResponseText = "<b>%s</b> - <i>%s, %s</i>\n\nID: <b>%s</b>\nURL: %v"
