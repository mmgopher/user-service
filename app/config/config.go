package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config represents aplication config object.
type Config struct {
}

// New creates new instance of Config object.
func New() (*Config, error) {

	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, errors.Wrap(err, "impossible to parse app configuration ")
	}

	return &c, nil
}
