package config

import (
	"fmt"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		logger   *slog.Logger
		API      API        `envconfig:"API"`
		LogLevel slog.Level `envconfig:"LOG_LEVEL"`
	}

	API struct {
		Host            string `envconfig:"HOST"`
		Port            int    `envconfig:"PORT"`
		ShutdownTimeout int    `envconfig:"SHUTDOWN_TIMEOUT_SECONDS"`
	}
)

func NewConfig(logger *slog.Logger) *Config {
	return &Config{
		logger: logger,
	}
}

func (c *Config) Load() error {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		c.logger.Debug("failed to load .env file", "error", err)
	}

	// Parse environment variables
	if err := envconfig.Process("", c); err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	return nil
}
