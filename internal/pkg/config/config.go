package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env                  string `env:"ENV" `
	Logger               string `env:"LOGGER_LEVEL"`
	APP_PORT             string `env:"APP_PORT"`
	PISTON_HOST          string `env:"PISTON_HOST"`
	DB_CONNECTION_STRING string `env:"DB_CONNECTION_STRING"`
}

func CreateConfig() (*Config, error) {
	godotenv.Load()
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		slog.Error("Error writing to env", "error", err)
		return nil, errors.Wrap(err, "cannot read config from environment variables")
	}

	return &cfg, nil
}
