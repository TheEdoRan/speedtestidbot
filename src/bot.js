import dotenv from "dotenv";
import { Telegraf, Markup } from "telegraf";
import escape from "html-escape";
import { memoSearchByName } from "./api.js";

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
    let { data: servers } = await memoSearchByName(ctx.inlineQuery.query);

    // Process first 10 results.
    servers = servers
      .slice(0, 10)
      .map(({ name, country, sponsor, id, host }) => ({
        type: "article",
        id,
        title: sponsor,
        description: `${name}, ${country}`,
        input_message_content: {
          message_text: `${sponsor} - ${name}, ${country}\n\nID: ${id}\nHOST: ${host}`,
        },
        ...Markup.inlineKeyboard([
          Markup.button.url(
            "🌐  Test with this server",
            `https://speedtest.net/server/${id}`,
          ),
        ]),
      }));

    return await ctx.answerInlineQuery(servers);
  } catch (e) {
    console.error(e.message);
  }
});

bot.launch();

// Enable graceful stop
process.once("SIGINT", () => bot.stop("SIGINT"));
process.once("SIGTERM", () => bot.stop("SIGTERM"));
