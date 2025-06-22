package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/robfig/cron/v3"
	"github.com/rooslunn/frenzy/internal/config"
	"github.com/rooslunn/frenzy/internal/services"
)

const (
	GameStarted = "Frenzy Helper started"
)

// done: image for day

// todo: test frenzy (mock, fakes)

// todo: frenzy simple meaning (for ai prompt misunderstanding)
// todo: lang pair
// todo: two different message per a day (change getting frenzy)

func main() {

	log := setupLogger()
	log.Info("Log set up")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	telegram, err := services.NewTelegram(cfg.BotToken, cfg.ChatID)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info(telegram.Status)


	// >>> test one

	// ctx := context.Background()

	// frenzy, err := services.FetchFrenzy(ctx, cfg.AiKey)
	// if err != nil {
	// 	log.Error(err.Error())
	// 	os.Exit(1)
	// }

	// err = telegram.SendFrenzy(ctx, frenzy)
	// if err != nil {
	// 	log.Error(err.Error())
	// } else {
	// 	log.Info("frenzy sent")
	// }
	// os.Exit(0)

	// <<< test one

	crono := cron.New()
	defer crono.Stop()


	crono.AddFunc(cfg.Schedule, func() {
		ctx := context.Background()

		frenzy, err := services.FetchFrenzy(ctx, cfg.AiKey)
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		err = telegram.SendFrenzy(ctx, frenzy)
		if err != nil {
			log.Error(err.Error())
		} else {
			log.Info("frenzy sent")
		}
	})

	crono.Start()
	log.Info(GameStarted)

	select {}

}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
