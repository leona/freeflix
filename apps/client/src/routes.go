package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "1",
	})
}

type AddRequest struct {
	Magnet string `json:"magnet"`
}

func AddHandler(c *gin.Context) {
	var request AddRequest
	c.BindJSON(&request)
	err := torrentClient.Add(request.Magnet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func RemoveHandler(c *gin.Context) {
	torrentClient.Remove(c.Param("hash"))
	c.JSON(http.StatusOK, gin.H{})
}

type RemoveRequest struct {
	Title string `json:"title"`
}

func RemoveTitleHandler(c *gin.Context) {
	log.Println("Request to delete torrent")
	var request RemoveRequest
	c.BindJSON(&request)

	if len(request.Title) < 4 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid title",
		})
		return
	}

	path := filepath.Join(config.OutputPath, request.Title)
	log.Println("Deleting", path)
	os.RemoveAll(path)
	c.JSON(http.StatusOK, gin.H{})
}

func DownloadsHandler(c *gin.Context) {
	torrents := torrentClient.List()
	names := []string{}

	for _, torrent := range torrents {
		names = append(names, torrent.Name)
	}

	downloads := GetDownloads(names)

	c.JSON(http.StatusOK, gin.H{
		"complete": downloads,
		"active":   torrents,
	})
}

func QueryHandler(c *gin.Context) {
	query := c.Query("q")
	torrents, err := scrapeClient.Query(query)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, torrents)
}
