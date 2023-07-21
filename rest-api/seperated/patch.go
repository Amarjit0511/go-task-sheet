package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func patchAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	_, err := fetchAlbumFromDatabase(id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	if err := updateAlbumFieldsInDatabase(updatedAlbum, id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedAlbum)
}
