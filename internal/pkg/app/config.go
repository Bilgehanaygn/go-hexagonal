package app

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	// Application holds core application settings

	DB struct {
		DbUrl string `env:"DB_URL"`
	}

	App struct {
		CORS struct {
			AllowedOrigins []string `env:"CORS_ALLOWED_ORIGINS" envDefault:"http://localhost:2999,http://localhost:3000"`
		}

		Port string `env:"PORT"`
	}
}

// NewConfig creates a new Config instance with values from environment variables
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}
