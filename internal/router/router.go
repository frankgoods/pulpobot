package router

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Talker interface {
	Talk(chatID int64)
	Reply(chatID int64)
}

type Repository interface {
	LoadChatID() (int64, error)
	SaveChatID(chatID int64) error
	LoadPass() (string, error)
}

type Router struct {
	bot    *tgbotapi.BotAPI
	talker Talker
	repo   Repository
	chatID int64
	pass   string
}

func NewRouter(bot *tgbotapi.BotAPI, talker Talker, repo Repository) *Router {
	pass, err := repo.LoadPass()
	if err != nil {
		log.Fatal("Failed to load password", err)
	}

	chatID, err := repo.LoadChatID()
	if err == nil {
		talker.Talk(chatID)
	}

	return &Router{
		bot:    bot,
		talker: talker,
		repo:   repo,
		chatID: chatID,
		pass:   pass,
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	if msg := update.Message; msg != nil {
		if msg.IsCommand() {
			r.HandleCommand(msg)
		} else {
			r.HandleMessage(msg)
		}
	}
}

func (r *Router) HandleCommand(msg *tgbotapi.Message) {
	newChatID := msg.Chat.ID

	if r.chatID == newChatID {
		r.Reply(msg.Chat.ID, "I'm already talking to you, my friend! No need to command me.")
		return
	}

	cmd := msg.Command()
	if cmd != "start" {
		r.ShowHint(msg)
		return
	}

	if msg.CommandArguments() != r.pass {
		r.Reply(newChatID, "I dont know you! I won't talk to you!")
		r.ShowHint(msg)
		return
	}

	r.chatID = newChatID

	if err := r.repo.SaveChatID(newChatID); err != nil {
		log.Println("Falied to store chatID to file:", err)
	}

	r.talker.Talk(newChatID)
	r.Reply(newChatID, "Ok, hi, I recognize you! From now on I will be sending you random messages once in a while!")
	log.Println("Talking now to", newChatID)
}

func (r *Router) HandleMessage(msg *tgbotapi.Message) {
	log.Println("Incomming message:", msg.Text)

	if msg.Chat.ID == r.chatID {
		r.talker.Reply(msg.Chat.ID)
	} else {
		r.ShowHint(msg)
	}
}

func (r *Router) ShowHint(msg *tgbotapi.Message) {
	r.Reply(msg.Chat.ID, "type /start [secret pass that you have]")
}

func (r *Router) Reply(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := r.bot.Send(msg)
	if err != nil {
		log.Println("Error sending message", err)
	}
}
