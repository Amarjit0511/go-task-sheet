package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id" db:"id"`
	Title  string  `json:"title" db:"title"`
	Artist string  `json:"artist" db:"artist"`
	Price  float64 `json:"price" db:"price"`
}

var db *sqlx.DB

func main() {
	// Open the database connection using the provided connection string.
	var err error
	db, err = sqlx.Connect("postgres", "postgres://amarjit:amarjit@localhost/sendgriddb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbumByID)
	router.PATCH("/albums/:id", patchAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8082")
}

// getAlbums responds with the list of all albums from the database as JSON.
func getAlbums(c *gin.Context) {
	var albums []album
	err := db.Select(&albums, "SELECT * FROM albums")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

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
	_, err := db.NamedExec("INSERT INTO albums (id, title, artist, price) VALUES (:id, :title, :artist, :price)", newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var a album
	err := db.Get(&a, "SELECT * FROM albums WHERE id=$1", id)
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
	var existingAlbum album
	err := db.Get(&existingAlbum, "SELECT * FROM albums WHERE id=$1", id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Update the album in the database.
	_, err = db.NamedExec("UPDATE albums SET title=:title, artist=:artist, price=:price WHERE id=:id", updatedAlbum)
	if err != nil {
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
	var existingAlbum album
	err := db.Get(&existingAlbum, "SELECT * FROM albums WHERE id=$1", id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Update the album in the database.
	query := "UPDATE albums SET"
	params := make(map[string]interface{})

	if updatedAlbum.Title != "" {
		query += " title=:title,"
		params["title"] = updatedAlbum.Title
	}
	if updatedAlbum.Artist != "" {
		query += " artist=:artist,"
		params["artist"] = updatedAlbum.Artist
	}
	if updatedAlbum.Price != 0 {
		query += " price=:price,"
		params["price"] = updatedAlbum.Price
	}

	// Remove the trailing comma from the query.
	query = query[:len(query)-1]

	query += " WHERE id=:id"
	params["id"] = id

	_, err = db.NamedExec(query, params)
	if err != nil {
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
	var existingAlbum album
	err := db.Get(&existingAlbum, "SELECT * FROM albums WHERE id=$1", id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Delete the album from the database.
	_, err = db.Exec("DELETE FROM albums WHERE id=$1", id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}
