package config

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	Port string
	LogLevel string
	APIVersion string
	APIName string
}

var (
	cfg *Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{
				Port:       os.Getenv("PORT"),
				LogLevel:   os.Getenv("LOG_LEVEL"),
				APIVersion: os.Getenv("API_VERSION"),
				APIName:    os.Getenv("API_NAME"),
				// Initialize other fields
		}

		// Validate required environment variables
		missingVars := cfg.validate()
		if len(missingVars) > 0 {
				fmt.Fprintf(os.Stderr, "Missing required environment variables: %v\n", missingVars)
				os.Exit(1)
		}
})
return cfg
}

func (c *Config) validate() []string {
	var missing []string
	if c.Port == "" {
			missing = append(missing, "PORT")
	}
	if c.LogLevel == "" {
			missing = append(missing, "LOG_LEVEL")
	}
	if c.APIVersion == "" {
			missing = append(missing, "API_VERSION")
	}
	if c.APIName == "" {
			missing = append(missing, "API_NAME")
	}
	return missing
}