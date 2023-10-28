package main

import (
	"io/ioutil"
	"log"
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

	for _, dir := range directories {
		if dir.IsDir() && !stringInSlice(dir.Name(), filter) {
			downloads = append(downloads, Download{
				Name: dir.Name(),
			})
		}
	}

	return downloads
}
