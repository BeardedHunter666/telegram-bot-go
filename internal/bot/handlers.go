package bot

import (
	tele "gopkg.in/telebot.v3"
)

// RegisterHandlers adds all command handlers
func RegisterHandlers(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("👋 Hi there! I'm a Telegram bot written in Go. Type something!")
	})

	b.Handle("/help", func(c tele.Context) error {
		return c.Send("🤖 Commands:\n/start - greeting\n/help - list of commands\n/echo <text> - repeats your message")
	})

	b.Handle("/echo", func(c tele.Context) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Send("⚠️ Usage: /echo <text>")
		}
		return c.Send("🪞 " + c.Message().Payload)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("📨 You wrote: " + c.Text())
	})
}