package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	HTTPPort       string `env:"http_port" required:"true"`
	PGUrl          string `env:"pg_url" required:"true"`
	LeetCodeApiURL string `env:"leetcode_api_url" required:"true"`
}

func Get() (Config, error) {
	return parseConfig()
}

func parseConfig() (cfg Config, err error) {
	err = godotenv.Load()
	if err != nil {
		return cfg, err
	}

	err = envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
