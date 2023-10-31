package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/", IndexHandler)
	r.GET("/downloads", DownloadsHandler)
	r.POST("/add", AddHandler)
	r.DELETE("/remove/:hash", RemoveHandler)
	r.DELETE("/remove", RemoveTitleHandler)
	r.GET("/query", QueryHandler)
}

func ServeApi() {
	r := gin.Default()
	setupRoutes(r)
	log.Print("Starting server on port " + config.ApiPort)
	r.Run(":" + config.ApiPort)
}
