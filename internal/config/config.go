package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Server   Server   `envPrefix:"SERVER_"`
	Database Database `envPrefix:"DATABASE_"`
}

type Server struct {
	BindAddressPublic string `env:"BIND_ADDRESS_PUBLIC" envDefault:"0.0.0.0:8080"`
}

type Database struct {
	Host     string `env:"HOST"`
	Name     string `env:"NAME"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
}

func New() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
