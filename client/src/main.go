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
		_ = MakeMullvad(config.MullvadAccount)
	}

	if config.WireguardEnable {
		wireguardClient = MakeWireguard()
	}

	torrentClient = MakeTorrentclient()
	ServeApi()
}
