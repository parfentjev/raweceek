package config

import "github.com/caarlos0/env/v11"

type Config struct {
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabaseName     string `env:"DATABASE_NAME"`
	DatabaseUser     string `env:"DATABASE_USER"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
}

func New() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
