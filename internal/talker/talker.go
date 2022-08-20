package talker

import (
	"log"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Talker struct {
	chatID int64
	sendID chan<- int64
	bot    *tgbotapi.BotAPI
}

func NewTalker(bot *tgbotapi.BotAPI) *Talker {
	var chatID int64
	idText, err := os.ReadFile("priv/chatid")
	if err == nil {
		if id, err := strconv.ParseInt(string(idText), 10, 64); err == nil {
			chatID = id
		}
	}
	t := &Talker{
		bot: bot,
	}
	t.Talk(chatID)
	return t
}

func (t *Talker) Talk(chatID int64) {
	if t.chatID == chatID {
		return
	}

	t.chatID = chatID

	if t.sendID == nil {
		ch := make(chan int64)
		t.sendID = ch
		go func(recieveID <-chan int64, bot *tgbotapi.BotAPI, chatID int64) {
			t := time.NewTimer(10 * time.Second)
			for {
				select {
				case <-t.C:
					msg := tgbotapi.NewMessage(chatID, "random message")
					_, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message", err)
					}
					t.Reset(10 * time.Second)
				case newID := <-recieveID:
					chatID = newID
				}
			}
		}(ch, t.bot, chatID)
	} else {
		t.sendID <- chatID
	}
}
