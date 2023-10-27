package main

import (
	"github.com/gin-gonic/gin"
)

var torrentClient *Torrentclient
var wireguardClient *Wireguard
var config *Config

func main() {
	config = MakeConfig()
	go cleanup()
	gin.SetMode(gin.ReleaseMode)

	if config.MullvadAccount != "" {
		mullvad := MakeMullvad(config.MullvadAccount)
		mullvad.GetServers()
	}

	if config.WireguardEnable {
		wireguardClient = MakeWireguard(config.WireguardConfigPath)
		wireguardClient.Connect()
	}

	torrentClient = MakeTorrentclient()
	ServeApi()
}
