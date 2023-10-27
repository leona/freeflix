package main

import (
	"os"
	"time"
	"log"
)

type Config struct {
	WireguardEnable     bool
	WireguardConfigPath string
	OutputPath          string
	ApiPort             string
	MullvadAccount      string
	MaxDownloadAge      time.Duration
}

func MakeConfig() *Config {
	config := &Config{
		WireguardEnable:     os.Getenv("WIREGUARD_ENABLED") == "true",
		WireguardConfigPath: DefaultString(os.Getenv("WIREGUARD_CONFIG_PATH"), "/config/1.conf"),
		OutputPath:          DefaultString(os.Getenv("OUTPUT_PATH"), "/data"),
		ApiPort:             DefaultString(os.Getenv("API_PORT"), "80"),
		MullvadAccount:      DefaultString(os.Getenv("MULLVAD_ACCOUNT"), ""),
		MaxDownloadAge:      time.Duration(DefaultInt(os.Getenv("MAX_DOWNLOAD_AGE"), 7)),
	}

	if !config.WireguardEnable {
		log.Println("WARNING: Wireguard is disabled", os.Getenv("WIREGUARD_ENABLE"))
	}

	return config
}
