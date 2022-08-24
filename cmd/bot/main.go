package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"pulpobot/internal/repository"
	"pulpobot/internal/router"
	"pulpobot/internal/talker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	minDelay, randDelay := getFlags()

	rand.Seed(time.Now().Unix())

	repo := repository.NewRepository("priv")
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

	talker := talker.NewTalker(bot, repo, minDelay, randDelay)
	router := router.NewRouter(bot, talker, repo)

	for update := range updates {
		router.HandleUpdate(update)
	}
}

func getFlags() (int, int) {
	const (
		minDelayName  = "min"
		randDelayName = "rand"
	)
	var minDelay = flag.Int(minDelayName, 35, "Minimum time period befor a new message is sent to the subcriber")
	var randDelay = flag.Int(randDelayName, 30, "Additional random delay that is added to minimum delay")

	flag.Parse()

	if *minDelay <= 0 {
		log.Fatalf("Wrong value for '%s' specified", minDelayName)
	}

	if *randDelay <= 0 {
		log.Fatalf("Wrong value for '%s' specified", randDelayName)
	}

	return *minDelay, *randDelay
}
