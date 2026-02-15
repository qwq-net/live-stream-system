package config

import (
	"os"
)

type Config struct {
	RTMPPort string
	HTTPPort string
}

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
