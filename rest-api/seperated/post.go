package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	if err := insertAlbumIntoDatabase(newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
