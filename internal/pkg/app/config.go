package app

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	// Application holds core application settings
	App struct {

		// CORS holds CORS configuration
		CORS struct {
			// AllowedOrigins is a list of origins a cross-domain request can be executed from.
			// In production, this should be your specific domain(s).
			// Example: https://app.autopilot.domain
			AllowedOrigins []string `env:"CORS_ALLOWED_ORIGINS" envDefault:"http://localhost:2999,http://localhost:3000"`
		}

		Database struct {
			// Worker is the worker database URL
			Worker string `env:"WORKER_DB_URL" envDefault:"postgres://postgres:postgres@localhost:5432/worker?sslmode=disable&search_path=public&pool_max_conns=25&pool_min_conns=2&pool_max_conn_lifetime=1h&pool_max_conn_idle_time=30m&pool_health_check_period=1m"`
		}

		Environment string `env:"APP_ENV" envDefault:"development"`

		// Service holds the service name
		Service string `env:"APP_SERVICE" envDefault:"api"`
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