package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func tick() {
	log.Println("Running cleanup for downloads older than", int(config.MaxDownloadAge), "days")
	downloads := GetDownloads([]string{})

	for _, download := range downloads {
		path := filepath.Join(config.OutputPath, download.Name)
		info, err := os.Stat(path)

		if err != nil {
			log.Println(err)
			continue
		}

		if info.ModTime().Add(config.MaxDownloadAge * 24 * time.Hour).Before(time.Now()) {
			err := os.RemoveAll(path)

			if err != nil {
				log.Println(err)
				continue
			}

			log.Println("Deleted:", download.Name)
		}
	}
}

func cleanup() {
	interval := 6 * time.Hour
	tick()

	for range time.Tick(interval) {
		tick()
	}
}
