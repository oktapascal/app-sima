package bootstraps

import (
	"github.com/joho/godotenv"
	"os"
)

// Config is an interface for accessing configuration values.
type Config interface {
	// Get returns the value of the configuration option with the given key.
	Get(key string) string
}

// ConfigImpl is a concrete implementation of the Config interface.
type ConfigImpl struct {
}

// Get returns the value of the configuration option with the given key.
func (config *ConfigImpl) Get(key string) string {
	// Look up the value of the given key in the environment.
	return os.Getenv(key)
}

// New creates a new ConfigImpl object and loads environment variables from one or more files.
func New(filenames ...string) Config {
	// Load the environment variables from the specified files.
	err := godotenv.Load(filenames...)

	// If there is an error, panic.
	if err != nil {
		panic(err)
	}

	// Return a new ConfigImpl object.
	return &ConfigImpl{}
}
