package main

import (
	"log"
	"os"

	"pulpobot/internal/router"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	token, err := loadToken()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	router := router.NewRouter(bot)

	for update := range updates {
		router.HandleUpdate(update)
	}
}

func loadToken() (string, error) {
	token, err := os.ReadFile("priv/token")
	if err != nil {
		return "", err
	}
	return string(token), nil
}
