package config

import (
	"github.com/sakirsensoy/genv"
)

type Config struct {
	Env       string
	LAPIKey   string
	LSecret   string
	RID       string
	RSecret   string
	RUserName string
	RPassword string
}

func New() *Config {
	cfg := Config{
		Env:       genv.Key("ENV").String(),
		LAPIKey:   genv.Key("LAST_FM_API_KEY").String(),
		LSecret:   genv.Key("LAST_FM_SECRET").String(),
		RID:       genv.Key("REDDIT_APP_ID").String(),
		RSecret:   genv.Key("REDDIT_APP_SECRET").String(),
		RUserName: genv.Key("REDDIT_USER_NAME").String(),
		RPassword: genv.Key("REDDIT_PASSWORD").String(),
	}

	return &cfg
}
