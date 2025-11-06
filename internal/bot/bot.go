package bot

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

// Start create and run Telegram-bot
func Start(token string) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Initializing handlers
	RegisterHandlers(b)

	log.Println("✅ Bot is running...")
	b.Start()
}
