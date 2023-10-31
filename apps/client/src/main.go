package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leona/freeflix/client/src/scraper"
)

var torrentClient *Torrentclient
var config *Config
var scrapeClient *scraper.Scraper

func main() {
	scrapeClient = scraper.MakeScraper()
	config = MakeConfig()
	go cleanup()
	gin.SetMode(gin.ReleaseMode)
	torrentClient = MakeTorrentclient()
	ServeApi()
}
