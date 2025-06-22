package services

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	TGStatusCreated = "Telegram Service ready"
)

type Telegram struct {
	bot    *tgbotapi.BotAPI
	chatID int64
	Status string
}

func NewTelegram(token string, chatID int64) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Telegram{
		bot: bot, chatID: chatID, Status: TGStatusCreated,
	}, nil
}

func (tg *Telegram) SendFrenzy(ctx context.Context, frenzy Frenzy) error {

	if FileExists(frenzy.ImagePath) {
		return tg.sendPhotoMessage(frenzy)
	}

	return tg.sendTextMessage(frenzy)
}

func (tg *Telegram) sendPhotoMessage(frenzy Frenzy) error {
	text := bimboFrenzy(frenzy)
	file := tgbotapi.FilePath(frenzy.ImagePath)

	msg := tgbotapi.NewPhoto(tg.chatID, file)
	msg.Caption = text
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	_, err := tg.bot.Send(msg)

	return err
}

func (tg *Telegram) sendTextMessage(frenzy Frenzy) error {
	text := bimboFrenzy(frenzy)

	msg := tgbotapi.NewMessage(tg.chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	_, err := tg.bot.Send(msg)

	return err
}
