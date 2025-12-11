package config

import (
    "log/slog"

    "github.com/caarlos0/env/v10"
    "github.com/joho/godotenv"
)

func LoadConfig() (*Config, error) {
    if err := godotenv.Load(); err != nil {
        slog.Warn(".env file not found — using environment variables only")
    }

    var cfg Config
    if err := env.Parse(&cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}
