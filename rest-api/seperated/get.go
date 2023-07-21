package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	albums := fetchAlbumsFromDatabase()
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

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
