package config

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address         string        `env:"APP_ADDRESS" envDefault:"0.0.0.0:8080"`
	LogLevel        slog.Level    `env:"APP_LOG_LEVEL" envDefault:"INFO"`
	ShutdownTimeout time.Duration `env:"APP_SHUTDOWN_TIMEOUT" envDefault:"10s"`
	StorageType     string        `env:"STORAGE_TYPE" envDefault:"INMEMORY"`
	DB              *DBConfig
}

type DBConfig struct {
	Host string `env:"DB_HOST"`
	Port int    `env:"DB_PORT"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASSWORD"`
	Name string `env:"DB_NAME"`
}

func Load() (Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse base config: %w", err)
	}

	if cfg.StorageType == "POSTGRES" {
		var dbCfg DBConfig
		if err := env.Parse(&dbCfg); err != nil {
			return Config{}, fmt.Errorf("failed to parse DB config: %w", err)
		}
		cfg.DB = &dbCfg
	}

	return cfg, nil
}
