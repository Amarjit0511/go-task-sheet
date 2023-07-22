package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbumByID)
	router.PATCH("/albums/:id", patchAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	return router
}


// getAlbums responds with the list of all albums from the database as JSON.
func getAlbums(c *gin.Context) {
	albums := fetchAlbumsFromDatabase()
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body to the database.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	// Insert the new album into the database.
	if err := insertAlbumIntoDatabase(newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Fetch the album from the database.
	a, err := fetchAlbumFromDatabase(id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, a)
}

// putAlbumByID updates an album from JSON received in the request body.
func putAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	// Call BindJSON to bind the received JSON to updatedAlbum.
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	// Check if the album with the given ID exists.
	_, err := fetchAlbumFromDatabase(id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Update the album in the database.
	if err := updateAlbumInDatabase(updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

// patchAlbumByID partially updates an album from JSON received in the request body.
func patchAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	// Call BindJSON to bind the received JSON to updatedAlbum.
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	// Check if the album with the given ID exists.
	_, err := fetchAlbumFromDatabase(id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Update the album in the database.
	if err := updateAlbumFieldsInDatabase(updatedAlbum, id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

// deleteAlbumByID deletes an album whose ID value matches the id
// parameter sent by the client.
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Check if the album with the given ID exists.
	_, err := fetchAlbumFromDatabase(id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Delete the album from the database.
	if err := deleteAlbumFromDatabase(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}

// Helper functions for database interactions

func fetchAlbumsFromDatabase() []album {
	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var albums []album
	for rows.Next() {
		var a album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			panic(err)
		}
		albums = append(albums, a)
	}

	return albums
}

func fetchAlbumFromDatabase(id string) (album, error) {
	var a album
	err := db.QueryRow("SELECT * FROM albums WHERE id=$1", id).Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
	return a, err
}

func insertAlbumIntoDatabase(a album) error {
	_, err := db.Exec("INSERT INTO albums (id, title, artist, price) VALUES ($1, $2, $3, $4)", a.ID, a.Title, a.Artist, a.Price)
	return err
}

func updateAlbumInDatabase(a album) error {
	_, err := db.Exec("UPDATE albums SET title=$1, artist=$2, price=$3 WHERE id=$4", a.Title, a.Artist, a.Price, a.ID)
	return err
}

func updateAlbumFieldsInDatabase(a album, id string) error {
	query := "UPDATE albums SET"
	params := []interface{}{}

	if a.Title != "" {
		query += " title=$1,"
		params = append(params, a.Title)
	}
	if a.Artist != "" {
		query += " artist=$2,"
		params = append(params, a.Artist)
	}
	if a.Price != 0 {
		query += " price=$3,"
		params = append(params, a.Price)
	}

	// Remove the trailing comma from the query.
	query = query[:len(query)-1]

	query += " WHERE id=$4"
	params = append(params, id)

	_, err := db.Exec(query, params...)
	return err
}

func deleteAlbumFromDatabase(id string) error {
	_, err := db.Exec("DELETE FROM albums WHERE id=$1", id)
	return err
}
