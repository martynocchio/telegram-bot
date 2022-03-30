package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"telegram-bot/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	pocketClient, err := pocket.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocketClient, "http://localhost")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
