package bot

import (
	tele "gopkg.in/telebot.v3"
)

// RegisterHandlers додає всі обробники команд
func RegisterHandlers(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("👋 Привіт! Я Telegram-бот на Go. Напиши мені щось!")
	})

	b.Handle("/help", func(c tele.Context) error {
		return c.Send("🤖 Команди доступні для використання:\n/start - привітання\n/help - список команд\n/echo <текст> - повторюю ваше повідомлення")
	})

	b.Handle("/echo", func(c tele.Context) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Send("⚠️ Використання: /echo <текст>")
		}
		return c.Send("🪞 " + c.Message().Payload)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("📨 Ви написали: " + c.Text())
	})
}