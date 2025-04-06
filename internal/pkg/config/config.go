package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"github.com/joho/godotenv"
)

type Config struct {
	Env      string `env:"ENV" `
	Logger   string `env:"LOGGER_LEVEL"`
	APP_PORT string `env:"APP_PORT"`
	PISTON_HOST string `env:"PISTON_HOST"`
	DB_CONNECTION_STRING string `env:"DB_CONNECTION_STRING"`
}

func CreateConfig() (*Config, error) {
	godotenv.Load()
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config from environment variables")
	}

	return &cfg, nil
}