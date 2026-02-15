package config

import (
	"os"
)

// Config holds the application configuration.
type Config struct {
	RTMPPort string
	HTTPPort string
}

// Load loads the configuration from environment variables.
// Default values are used if environment variables are not set.
func Load() *Config {
	rtmpPort := os.Getenv("RTMP_PORT")
	if rtmpPort == "" {
		rtmpPort = "1935"
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	return &Config{
		RTMPPort: rtmpPort,
		HTTPPort: httpPort,
	}
}
