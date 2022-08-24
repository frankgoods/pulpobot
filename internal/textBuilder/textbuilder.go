package textbuilder

import (
	"encoding/json"
	"log"
	"math/rand"
)

type Repository interface {
	LoadText() ([]byte, error)
}

type TextBuilder struct {
	Begin     []string
	Middle    []string
	End       []string
	Whole     []string
	ReplyText []string
}

func NewTextBuilder(repo Repository) *TextBuilder {
	tb := &TextBuilder{}

	data, err := repo.LoadText()
	if err != nil {
		log.Fatal("Failed to load text for talker:", err)
	}
	err = json.Unmarshal(data, tb)
	if err != nil {
		log.Fatal("Failed to unmarshal talker text", err)
	}

	return tb
}

func (tb *TextBuilder) Say() string {
	if rand.Intn(100) < 95 {
		return tb.Begin[rand.Intn(len(tb.Begin))] + tb.Middle[rand.Intn(len(tb.Middle))] + tb.End[rand.Intn(len(tb.End))]
	} else {
		return tb.Whole[rand.Intn(len(tb.Whole))]
	}
}

func (tb *TextBuilder) Reply() string {
	return tb.ReplyText[rand.Intn(len(tb.ReplyText))]
}
