package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	HTTPPort        string `env:"http_port,required"`
	PGUrl           string `env:"pg_url,required"`
	LeetCodeApiURL  string `env:"leetcode_api_url,required"`
	RateLimitWindow int    `env:"rate_limit_window,default=5"`
	RateBlockTime   int    `env:"rate_block_time,default=25"`
}

func Get() (Config, error) {
	return parseConfig()
}

func parseConfig() (cfg Config, err error) {
	godotenv.Load()

	err = envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
