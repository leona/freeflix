package main

import (
	"log"
	"os"
	"time"
)

type Config struct {
	WireguardEnable  bool
	OutputPath       string
	ApiPort          string
	MullvadAccount   string
	MullvadCountries []string
	MaxDownloadAge   time.Duration
}

func MakeConfig() *Config {
	config := &Config{
		WireguardEnable:  os.Getenv("WIREGUARD_ENABLED") == "true",
		OutputPath:       DefaultString(os.Getenv("OUTPUT_PATH"), "/data"),
		ApiPort:          DefaultString(os.Getenv("API_PORT"), "80"),
		MullvadAccount:   DefaultString(os.Getenv("MULLVAD_ACCOUNT"), ""),
		MullvadCountries: DefaultSlice(os.Getenv("MULLVAD_COUNTRIES"), []string{"nl"}),
		MaxDownloadAge:   time.Duration(DefaultInt(os.Getenv("MAX_DOWNLOAD_AGE"), 7)),
	}

	if !config.WireguardEnable {
		log.Println("WARNING: Wireguard is disabled", os.Getenv("WIREGUARD_ENABLE"))
	}

	return config
}
