package main

import (
	"github.com/gin-gonic/gin"
	//"time"
)

var torrentClient *Torrentclient
var config *Config

func main() {
	config = MakeConfig()
	go cleanup()
	gin.SetMode(gin.ReleaseMode)
	torrentClient = MakeTorrentclient()
	ServeApi()
}
