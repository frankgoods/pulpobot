package main

import (
	"log"
	"math/rand"
	"time"

	"pulpobot/internal/repository"
	"pulpobot/internal/router"
	"pulpobot/internal/talker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	rand.Seed(time.Now().Unix())

	repo := repository.NewRepository("priv1")
	token, err := repo.LoadToken()
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

	talker := talker.NewTalker(bot, repo)
	router := router.NewRouter(bot, talker, repo)

	for update := range updates {
		router.HandleUpdate(update)
	}
}
