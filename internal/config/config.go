package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	ChatID int64
	Schedule string
	AiKey string
}

const (
	configPath1 = ".env" // todo: set as parameter

	ErrBotTokenMissing = "missing FRENZY_BOT_TOKEN"
	ErrChatIdMissing = "missing FRENZY_CHAT_ID"
	ErrAiKeyMissing = "missing AI key"
	ErrChatIdWrongValue = "can't convert CHAT_ID to int"

	CronSchedule = "0 5,7,11,13,17,19 * * *"
)

func LoadConfig() (*Config, error) {

	err := godotenv.Load(configPath1)
	if err != nil {
		return nil, err
	}

	token := os.Getenv("FRENZY_BOT_TOKEN")
	chatIdStr := os.Getenv("FRENZY_CHAT_ID")
	
	if token == "" {
		return nil, errors.New(ErrBotTokenMissing)
	}
	if chatIdStr == "" {
		return nil, errors.New(ErrChatIdMissing)
	}

	chatId, err := strconv.ParseInt(chatIdStr, 10, 64)
	if err != nil {
		return nil, errors.New(ErrChatIdWrongValue)
	}

	aiKey := os.Getenv("AI_STUDIO_KEY")
	if aiKey == "" {
		return nil, errors.New(ErrAiKeyMissing)
	}

	return &Config{
		BotToken: token,
		ChatID: chatId,
		AiKey: aiKey,
		Schedule: CronSchedule,
	}, nil
}