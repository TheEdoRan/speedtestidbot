import dotenv from "dotenv";
import { Telegraf, Markup } from "telegraf";
import escape from "html-escape";
import { searchByName } from "./api.js";

dotenv.config();

const bot = new Telegraf(process.env.BOT_TOKEN);

bot.start((ctx) =>
  ctx
    .reply(
      `Hi <b>${escape(
        ctx.from.first_name,
      )}</b>!\n\nYou can invoke me in any chat you want, providing the name of the <b>Speedtest server</b> you're interested to.\n\nExample:\n<pre>@speedtestidbot Vodafone</pre>\n\nThen, you just have to select an item from the list, and you'll get the ID back.`,
      { parse_mode: "HTML" },
    )
    .catch((e) => console.error(e)),
);

bot.on("inline_query", async (ctx) => {
  try {
    let { data: servers } = await searchByName(ctx.inlineQuery.query);

    // Process first 10 results.
    servers = servers.map(({ name, country, sponsor, id, url }) => ({
      type: "article",
      id,
      title: sponsor,
      description: `${name}, ${country}`,
      input_message_content: {
        message_text: `<b>${sponsor}</b> - <i>${name}, ${country}</i>\n\nID: <b>${id}</b>\nURL: ${url.replace(
          ":8080/speedtest/upload.php",
          "",
        )}`,
        parse_mode: "HTML",
        disable_web_page_preview: true,
      },
      ...Markup.inlineKeyboard([
        Markup.button.url(
          "👉 Test with this server 👈",
          `https://speedtest.net/server/${id}`,
        ),
      ]),
    }));

    // Cache for 6 hours.
    return await ctx.answerInlineQuery(servers, { cache_time: 21600 });
  } catch (e) {
    console.error(e.message);
  }
});

bot.launch();

// Enable graceful stop
process.once("SIGINT", () => bot.stop("SIGINT"));
process.once("SIGTERM", () => bot.stop("SIGTERM"));
