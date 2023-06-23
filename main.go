package main

import (
	"github.com/ShivangShandilya/golang-api/handler"
	"github.com/gin-gonic/gin"
)

var HOST = ":8080"

func main() {
	router := gin.Default()
	router.GET("/", handler.HomeRoute)
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.PostAlbums)

	router.Run(HOST)
}
