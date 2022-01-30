import "dotenv/config";
import escape from "escape-html";
import { createServer } from "http";
import { Markup, Telegraf } from "telegraf";
import type { InlineQueryResultArticle } from "typegram";
import { searchByName } from "./api";
import type { SpeedtestServer } from "./api/types";
import { logger } from "./telegram/logger";

const bot = new Telegraf(process.env.BOT_TOKEN as string);

// Only for debugging purposes.
if (process.env.DEBUG === "true") {
  bot.use(logger());
}

// /start command.
bot.start((ctx) =>
  ctx
    .reply(
      `Hi <b>${escape(
        ctx.from.first_name
      )}</b>!\n\nYou can invoke me in any chat you want, providing the name of the <b>Speedtest server</b> you're interested to.\n\nExample:\n<pre>@speedtestidbot Vodafone</pre>\n\nThen, you just have to select an item from the list, and you'll get ID and other useful data back.`,
      { parse_mode: "HTML" }
    )
    .catch((e: any) => console.error(`/start error: ${e.message}`))
);

// Handle inline queries.
bot.on("inline_query", async (ctx) => {
  try {
    const { data: servers } = await searchByName(ctx.inlineQuery.query);

    const results: InlineQueryResultArticle[] = servers.map(
      ({ name, country, sponsor, id, url }: SpeedtestServer) => ({
        type: "article",
        id,
        title: sponsor,
        description: `${name}, ${country}`,
        input_message_content: {
          message_text: `<b>${sponsor}</b> - <i>${name}, ${country}</i>\n\nID: <b>${id}</b>\nURL: ${url.replace(
            ":8080/speedtest/upload.php",
            ""
          )}`,
          parse_mode: "HTML",
          disable_web_page_preview: true,
        },
        ...Markup.inlineKeyboard([
          Markup.button.url(
            "↗️  Test with this server",
            `https://speedtest.net/server/${id}`
          ),
        ]),
      })
    );

    // Cache for 6 hours.
    return await ctx
      .answerInlineQuery(results, { cache_time: 21600 })
      .catch((e: any) =>
        console.error(`answerInlineQuery error: ${e.message}`)
      );
  } catch (e: any) {
    console.error(e?.message);
  }
});

bot.launch();

// fly healthcheck.
const server = createServer((_, res) => {
  res.writeHead(200);
  res.end("ok");
});

server.listen(8080);

// Enable graceful stop.
process.once("SIGINT", () => {
  bot.stop("SIGINT");
  server.close();
});
process.once("SIGTERM", () => {
  bot.stop("SIGTERM");
  server.close();
});
