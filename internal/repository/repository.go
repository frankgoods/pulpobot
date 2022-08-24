package repository

import (
	"os"
	"strconv"
)

type Repository struct {
	path string
}

func NewRepository(path string) Repository {
	return Repository{path: path}
}

func (r Repository) LoadToken() (string, error) {
	token, err := os.ReadFile(r.path + "/token")
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (r Repository) LoadChatID() (int64, error) {
	var chatID int64
	idText, err := os.ReadFile(r.path + "/chatid")
	if err == nil {
		if id, err := strconv.ParseInt(string(idText), 10, 64); err == nil {
			chatID = id
		}
	}
	return chatID, err
}

func (r Repository) SaveChatID(chatID int64) error {
	return os.WriteFile(r.path+"/chatid", []byte(strconv.FormatInt(chatID, 10)), 0644)
}

func (r Repository) LoadText() ([]byte, error) {
	return os.ReadFile(r.path + "/talk")
}

func (r Repository) LoadPass() (string, error) {
	p, err := os.ReadFile(r.path + "/password")
	if err != nil {
		return "", err
	}
	return string(p), nil
}
