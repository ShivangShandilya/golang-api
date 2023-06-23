package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ShivangShandilya/golang-api/handler"
	"github.com/ShivangShandilya/golang-api/src"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {

	router := gin.Default()
	router.GET("/", handler.HomeRoute)

	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// get message key from the response body

	var data map[string]string
	err := json.NewDecoder(resp.Body).Decode(&data)
	assert.NoError(t, err)

	assert.Equal(t, "API is working fine", data["message"])

}

// For getting the list of albums:
func TestGetAlbums(t *testing.T) {

	// gin.Default - creates a router with the default middleware:
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)

	// httptest.NewRequest - creates a new http request
	req := httptest.NewRequest("GET", "/albums", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// check that the response code in "resp" matches http.StatusOK (which is 200)
	assert.Equal(t, http.StatusOK, resp.Code)

	// json.NewDecoder - creates a new JSON decoder that will read from the response body in "resp"
	// Decode - converts the JSON response back to the slice (deserialized)
	var albums []src.Album
	err := json.NewDecoder(resp.Body).Decode(&albums)
	assert.NoError(t, err)

	// check that the response contains the expected number of albums
	assert.Len(t, albums, 4)

}

func TestPostAlbums(t *testing.T) {

	// gin.Default - creates a router with the default middleware:
	router := gin.Default()
	router.POST("/albums", handler.PostAlbums)

	newAlbum := src.Album{
		ID:     "5",
		Title:  "Test",
		Artist: "Shivang Shandilya",
		Price:  10.00,
	}

	// json.Marshal - to serialize the newAlbum object into JSON format
	newAlbumJSON, err := json.Marshal(newAlbum)
	assert.NoError(t, err)

	// httptest.NewRequest - creates a new http request
	req := httptest.NewRequest("POST", "/albums", bytes.NewBuffer(newAlbumJSON))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// check that the response code in "resp" matches http.StatusOK (which is 200)
	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdAlbum src.Album
	err = json.NewDecoder(resp.Body).Decode(&createdAlbum)
	assert.NoError(t, err)

	// check that the created album matches the input:
	assert.Equal(t, newAlbum.Title, createdAlbum.Title)
	assert.Equal(t, newAlbum.Artist, createdAlbum.Artist)
	assert.Equal(t, newAlbum.Price, createdAlbum.Price)

}
