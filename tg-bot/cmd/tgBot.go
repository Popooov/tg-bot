/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

var (
	// TbToken
	TbotToken = os.Getenv("TELE_TOKEN")
)

// tgBotCmd represents the tgBot command
var tgBotCmd = &cobra.Command{
	Use:     "tgBot",
	Aliases: []string{"start"},
	Short:   "Telegram Bot",
	Long:    "A Telegram bot built with Golang and telebot.v3.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tgBot %s started", appVersion)
		tgBot, err := tele.NewBot(tele.Settings{
			URL:    "",
			Token:  TbotToken,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TOKEN env variable. %s", err)
			return
		}

		tgBot.Handle(tele.OnText, func(c tele.Context) error {
			handleTextMessage(tgBot, c.Message())
			return err
		})

		fmt.Println("Bot is now running. Press Ctrl+C to exit.")
		tgBot.Start()
	},
}

func handleTextMessage(tgBot *tele.Bot, m *tele.Message) {
	switch m.Text {
	case "/start":
		tgBot.Send(m.Sender, "Вітаю! Я тут, щоб вам допомогти. Відправте /help, щоб побачити доступні команди.")
	case "/time":
		tgBot.Send(m.Sender, fmt.Sprintf("Поточний час: %s", time.Now().Format("15:04:05")))
	case "/version":
		tgBot.Send(m.Sender, fmt.Sprintf("tgBot %s", appVersion))
	case "/help":
		tgBot.Send(m.Sender,
			`Доступні команди:
			/start - Початок роботи з ботом
			/time - Поточний час
			/version - Версія бота
			/help - Виведення списку команд
		`)
	default:
		tgBot.Send(m.Sender, fmt.Sprintf("Ти написав: %s", m.Text))
	}
}

func init() {
	rootCmd.AddCommand(tgBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tgBotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tgBotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
