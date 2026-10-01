package internal

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

type telegramConn struct {
	bot    *telebot.Bot
	chatId int64
}

func NewTelegramConnection(token string, chatId int64) (*telegramConn, error) {
	pref := telebot.Settings{
		Token: token,
	}
	// TODO: If error, try again (connection reset by peer)
	bot, err := telebot.NewBot(pref)
	if err != nil {
		return nil, err
	}
	return &telegramConn{
		bot:    bot,
		chatId: chatId,
	}, nil
}

func (listener *telegramConn) SendFile(filename string, message string) error {
	file := telebot.FromDisk(filename)
	if !file.OnDisk() {
		return fmt.Errorf("Arquivo %s não encontrado!", filename)
	}
	// TODO: Check if video is ok
	video := &telebot.Video{File: file, Caption: message}
	chat, err := listener.bot.ChatByID(listener.chatId)
	if err != nil {
		return err
	}

	_, err = video.Send(listener.bot, chat, nil)
	return err
}
