package internal

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Database DatabaseConfig `envPrefix:"DATABASE_"`
}

type DatabaseConfig struct {
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	Host     string `env:"HOST,required"`
	Port     string `env:"PORT,required"`
	Name     string `env:"NAME,required"`
	SSL      string `env:"SSL,required"`
}

func LoadConfig() (Config, error) {
	config := Config{}

	return config, env.Parse(&config)
}
