package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	// Replace "your/package/directory" with the actual package directory path
	"github.com/your/package/directory"
)

// Initialize and set up the Gin router for testing
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums/get", main.GetAlbums)
	router.GET("/albums/get/:id", main.GetAlbumByID)
	router.POST("/albums/post", main.PostAlbums)
	router.PUT("/albums/put/:id", main.PutAlbumByID)
	router.DELETE("/albums/delete/:id", main.DeleteAlbumByID)
	return router
}

func TestGetAlbums(t *testing.T) {
	// Set up a test server with the router
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/get", nil)
	router.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body and unmarshal it into a slice of albums
	var albums []main.Album
	err := json.Unmarshal(w.Body.Bytes(), &albums)

	// Assert that there are no errors and the response contains at least one album
	assert.NoError(t, err)
	assert.True(t, len(albums) > 0)
}

func TestPostAlbums(t *testing.T) {
	// Set up a test server with the router
	router := setupRouter()

	// Create a new album to be posted
	newAlbum := main.Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  9.99,
	}

	// Convert the newAlbum struct to JSON
	jsonData, err := json.Marshal(newAlbum)
	assert.NoError(t, err)

	// Set up a POST request with the JSON data
	req, _ := http.NewRequest("POST", "/albums/post", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform the POST request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response status code is 201 Created
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse the response body and unmarshal it into an album struct
	var responseAlbum main.Album
	err = json.Unmarshal(w.Body.Bytes(), &responseAlbum)

	// Assert that there are no errors and the response matches the posted album
	assert.NoError(t, err)
	assert.Equal(t, newAlbum, responseAlbum)
}

func TestPutAlbumByID(t *testing.T) {
	// Set up a test server with the router
	router := setupRouter()

	// Create a new album to be updated
	newAlbum := main.Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  9.99,
	}

	// Insert the new album into the database for testing
	err := main.InsertAlbumIntoDatabase(newAlbum)
	assert.NoError(t, err)

	// Modify the new album
	updatedAlbum := main.Album{
		ID:     "1",
		Title:  "Updated Album",
		Artist: "Updated Artist",
		Price:  19.99,
	}

	// Convert the updatedAlbum struct to JSON
	jsonData, err := json.Marshal(updatedAlbum)
	assert.NoError(t, err)

	// Set up a PUT request with the JSON data
	req, _ := http.NewRequest("PUT", "/albums/put/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform the PUT request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body and unmarshal it into an album struct
	var responseAlbum main.Album
	err = json.Unmarshal(w.Body.Bytes(), &responseAlbum)

	// Assert that there are no errors and the response matches the updated album
	assert.NoError(t, err)
	assert.Equal(t, updatedAlbum, responseAlbum)
}

func TestDeleteAlbumByID(t *testing.T) {
	// Set up a test server with the router
	router := setupRouter()

	// Create a new album to be deleted
	newAlbum := main.Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  9.99,
	}

	// Insert the new album into the database for testing
	err := main.InsertAlbumIntoDatabase(newAlbum)
	assert.NoError(t, err)

	// Set up a DELETE request
	req, _ := http.NewRequest("DELETE", "/albums/delete/1", nil)

	// Perform the DELETE request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body and unmarshal it into a map
	var responseMap map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &responseMap)

	// Assert that there are no errors and the response indicates successful deletion
	assert.NoError(t, err)
	assert.Equal(t, "album deleted", responseMap["Successfull"])
}

