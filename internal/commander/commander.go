package commander

import (
	"log"
	"pulpobot/internal/talker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Talker interface {
	Talk(chatID int64)
}

type Commander struct {
	bot    *tgbotapi.BotAPI
	talker Talker
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot:    bot,
		talker: talker.NewTalker(bot),
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message) {
	cmd := msg.Command()
	if cmd != "start" {
		c.ShowHint(msg)
		return
	}

	if msg.CommandArguments() != "secret_pass_code" {
		c.Reply(msg.Chat.ID, "I dont know you! I won't talk to you!")
		c.ShowHint(msg)
		return
	}

	c.talker.Talk(msg.Chat.ID)
	c.Reply(msg.Chat.ID, "Ok, hi, I recognize you! From now on I will be sending you random messages once in a while!")
}

func (c *Commander) ShowHint(msg *tgbotapi.Message) {
	c.Reply(msg.Chat.ID, "type /start [secret pass that you have]")
}

func (c *Commander) Reply(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("Error sending message", err)
	}
}
