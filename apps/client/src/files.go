package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	bolt "go.etcd.io/bbolt"
)

type Download struct {
	Name string `json:"name"`
}

func GetDownloads(filter []string) []Download {
	directories, err := ioutil.ReadDir(config.OutputPath)

	if err != nil {
		log.Panic(err)
	}

	downloads := []Download{}

	for _, path := range directories {
		if !stringInSlice(path.Name(), filter) && !strings.HasPrefix(path.Name(), ".") {
			downloads = append(downloads, Download{
				Name: path.Name(),
			})
		}
	}

	return downloads
}

func RemoveDownload(name string) {
	log.Println("Removing torrent:", name)
	torrents := torrentClient.Client.Torrents()

	for _, torrent := range torrents {
		if torrent.Name() == name {
			torrent.Drop()
			log.Println("Removed torrent:", torrent.Name())
			break
		}
	}

	path := filepath.Join(config.OutputPath, name)
	err := os.RemoveAll(path)

	if err != nil {
		log.Println("Failed to delete files:", err)
	}

	log.Println("Files deleted:", path)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("torrents.magnet.index"))
		err := b.Delete([]byte(name))
		return err
	})

	if err != nil {
		log.Println("Failed to delete database index:", err)
		return
	}
}
