package router

import (
	"pulpobot/internal/commander"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander interface {
	HandleCommand(callback *tgbotapi.Message)
	ShowHint(callback *tgbotapi.Message)
}

type Router struct {
	bot *tgbotapi.BotAPI
	com Commander
}

func NewRouter(bot *tgbotapi.BotAPI) *Router {
	return &Router{
		bot: bot,
		com: commander.NewCommander(bot),
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	if msg := update.Message; msg != nil {
		if !msg.IsCommand() {
			r.com.ShowHint(msg)
			return
		}
		r.com.HandleCommand(update.Message)
	}
}
