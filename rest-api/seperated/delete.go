package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	_, err := fetchAlbumFromDatabase(id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := deleteAlbumFromDatabase(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}
