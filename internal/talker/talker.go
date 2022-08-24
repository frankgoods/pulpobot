package talker

import (
	"log"
	"math/rand"
	textbuilder "pulpobot/internal/textBuilder"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Repository textbuilder.Repository

type Talker struct {
	chatID int64
	sendID chan<- int64
	bot    *tgbotapi.BotAPI
	tb     *textbuilder.TextBuilder
}

func NewTalker(bot *tgbotapi.BotAPI, repo Repository) *Talker {
	return &Talker{
		bot: bot,
		tb:  textbuilder.NewTextBuilder(repo),
	}
}

func (t *Talker) Talk(chatID int64) {
	if t.chatID == chatID {
		return
	}

	t.chatID = chatID

	if t.sendID == nil {
		ch := make(chan int64)
		t.sendID = ch

		go func(recieveID <-chan int64, bot *tgbotapi.BotAPI, chatID int64, tb *textbuilder.TextBuilder) {
			timer := time.NewTimer(randTime())

			for {
				select {
				case <-timer.C:
					say := t.tb.Say()
					log.Println("Sending message:", say)
					t.Send(chatID, say)
					timer.Reset(randTime())
				case newID := <-recieveID:
					chatID = newID
				}
			}
		}(ch, t.bot, chatID, t.tb)

	} else {
		t.sendID <- chatID
	}
}

func (t *Talker) Reply(chatID int64) {
	reply := t.tb.Reply()
	log.Println("Replying:", reply)
	t.Send(chatID, reply)
}

func randTime() time.Duration {
	if h := time.Now().Hour(); h == 23 || h < 9 {
		return time.Duration(9-h%23) * time.Hour
	}
	return time.Duration(rand.Intn(30)+35) * time.Minute
}

func (t *Talker) Send(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := t.bot.Send(msg)
	if err != nil {
		log.Println("Error sending message", err)
	}
}
