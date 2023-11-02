package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leona/freeflix/client/src/scraper"
	bolt "go.etcd.io/bbolt"
)

var torrentClient *Torrentclient
var config *Config
var scrapeClient *scraper.Scraper
var db *bolt.DB

func main() {
	scrapeClient = scraper.MakeScraper()
	config = MakeConfig()
	var err error
	db, err = bolt.Open(config.OutputPath+"/.clientdb", 0600, nil)
	FatalError(err)

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("torrents.magnet.index"))
		FatalError(err)
		return nil
	})

	defer db.Close()
	go cleanup()
	gin.SetMode(gin.ReleaseMode)
	torrentClient = MakeTorrentclient()
	torrentClient.CheckExisting()
	ServeApi()
}
