package cmd

import (
	"fmt"
	"os"

	"github.com/BeardedHunter666/telegram-bot-go/internal/bot"
	"github.com/BeardedHunter666/telegram-bot-go/internal/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "telegram-bot-go",
	Short: "Simple Telegram bot powered by Go and Telebot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting Telegram Bot...")
		token := config.GetToken()
		bot.Start(token)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
