package deps

import (
	"log"

	"github.com/gabezeck/test-api/internal/config"
	"github.com/shkh/lastfm-go/lastfm"
	"github.com/thecsw/mira"
	"go.uber.org/zap"
)

type Deps struct {
	LClient *lastfm.Api
	Logger  *zap.SugaredLogger
	RClient *mira.Reddit
}

func New(cfg *config.Config) *Deps {
	return &Deps{
		LClient: CreateLClient(cfg),
		Logger:  CreateLogger(),
		RClient: CreateRClient(cfg),
	}
}

func CreateLClient(cfg *config.Config) *lastfm.Api {
	return lastfm.New(cfg.LAPIKey, cfg.LSecret)
}

func CreateLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	return logger.Sugar()
}

func CreateRClient(cfg *config.Config) *mira.Reddit {
	client, err := mira.Init(mira.Credentials{
		ClientId:     cfg.RID,
		ClientSecret: cfg.RSecret,
		Username:     cfg.RUserName,
		Password:     cfg.RPassword,
		UserAgent:    "Apollon API",
	})
	if err != nil {
		log.Fatal(err)
	}

	return client
}
