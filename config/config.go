package config

import (
	"errors"
	"github.com/companieshouse/gofigure"
	log "github.com/sirupsen/logrus"
	"sync"
)

// Config holds configuration details set by the environment
type Config struct {
	MongoDBURL      string `env:"MONGODB_URL"       flag:"mongodb-url"       flagDesc:"MongoDB server URL"`
	MongoDBDatabase string `env:"MONGODB_DATABASE"  flag:"mongodb-database"  flagDesc:"MongoDB database for data"`
	LogLevel        string `env:"LOG_LEVEL"         flag:"log-level"         flagDesc:"Logging level of the application"`
}

var cfg *Config
var mtx sync.Mutex

// Get returns a pointer to a Config instance
// populated with values from environment or command-line flags
func Get() (*Config, error) {

	mtx.Lock()
	defer mtx.Unlock()

	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{}

	err := gofigure.Gofigure(cfg)
	if err != nil {
		return nil, err
	}

	mandatoryConfigsMissing := false

	if cfg.MongoDBURL == "" {
		log.Warn("MONGODB_URL not set in environment")
		mandatoryConfigsMissing = true
	}

	if cfg.MongoDBDatabase == "" {
		log.Warn("MONGODB_DATABASE not set in environment")
		mandatoryConfigsMissing = true
	}

	if mandatoryConfigsMissing {
		return nil, errors.New("mandatory configs missing from environment")
	}

	return cfg, nil
}
