package main

import (
	"os"
	"time"
)

type Config struct {
	OutputPath     string
	ApiPort        string
	MaxDownloadAge time.Duration
}

func MakeConfig() *Config {
	config := &Config{
		OutputPath:     "/data",
		ApiPort:        DefaultString(os.Getenv("API_PORT"), "80"),
		MaxDownloadAge: time.Duration(DefaultInt(os.Getenv("MAX_DOWNLOAD_AGE"), 21)),
	}

	return config
}
