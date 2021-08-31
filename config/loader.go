package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

// Loader is a utility to load config data into a config struct.
type Loader interface {
	// Load loads the configuration into the given conf struct, which ought therefore be passed by reference.
	Load(conf interface{}) error
}

type loader struct {
	envPrefix string // The env prefix to use when parsing environment variables into config fields.
}

// NewLoader returns a new config loader.
func NewLoader(envPrefix string) Loader {
	return &loader{
		envPrefix: envPrefix,
	}
}

// Load loads the configuration into the given conf struct, which ought therefore be passed by reference.
func (l *loader) Load(conf interface{}) error {
	confDir := os.Getenv(fmt.Sprintf("%s_CONFIG_DIRECTORY", l.envPrefix))

	if confDir == "" {
		confDir = "."
	}

	confFilePath := fmt.Sprintf("%s/config.yml", confDir)

	if _, err := os.Stat(confFilePath); os.IsNotExist(err) {
		return fmt.Errorf("config file cannot be found at %s", confFilePath)
	}

	return configor.
		New(&configor.Config{
			ENVPrefix:   l.envPrefix,
			Environment: os.Getenv("ENVIRONMENT"),
		}).
		Load(conf, confFilePath)
}
