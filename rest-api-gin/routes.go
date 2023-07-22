package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)

func setupRoutes(router *gin.Engine) {
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbumByID)
	router.PATCH("/albums/:id", patchAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
}

func getAlbums(c *gin.Context) {
	albums, err := fetchAlbumsFromDatabase()
	if err != nil {
		log.Println("Failed to fetch albums:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch albums"})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	a, err := fetchAlbumFromDatabase(id)
	if err != nil {
		log.Println("Album not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}
	c.JSON(http.StatusOK, a)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		log.Println("Invalid data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := insertAlbumIntoDatabase(newAlbum); err != nil {
		log.Println("Failed to insert album:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert album"})
		return
	}
	c.JSON(http.StatusCreated, newAlbum)
}

func putAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		log.Println("Invalid data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := updateAlbumInDatabase(id, updatedAlbum); err != nil {
		log.Println("Failed to update album:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update album"})
		return
	}
	c.JSON(http.StatusOK, updatedAlbum)
}



func patchAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		log.Println("Invalid data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := updateAlbumFieldsInDatabase(updatedAlbum, id); err != nil {
		log.Println("Failed to update album:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update album"})
		return
	}
	c.JSON(http.StatusOK, updatedAlbum)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	if err := deleteAlbumFromDatabase(id); err != nil {
		log.Println("Failed to delete album:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete album"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Album deleted"})
}
