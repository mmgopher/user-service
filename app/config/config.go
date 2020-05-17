package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config represents aplication config object.
type Config struct {
	DBUser   string `envconfig:"DB_USER" required:"true"`
	DBPass   string `envconfig:"DB_PASS" required:"true"`
	DBType   string `envconfig:"DB_TYPE" required:"true"`
	DBName   string `envconfig:"DB_NAME" required:"true"`
	DBHost   string `envconfig:"DB_HOST" required:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"error"`
}

// New creates new instance of Config object.
func New() (*Config, error) {

	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, errors.Wrap(err, "impossible to parse app configuration ")
	}

	return &c, nil
}
