package handler

import (
	"net/http"

	"github.com/ShivangShandilya/golang-api/src"
	"github.com/gin-gonic/gin"
)

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, src.Albums)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum src.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	src.Albums = append(src.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func HomeRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is working fine",
	})
}
