package main

import (
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

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
    router.PUT("/albums/:id", putAlbumByID)
    router.PATCH("/albums/:id", patchAlbumByID)
    router.DELETE("/albums/:id", deleteAlbumByID)

    router.Run("localhost:8082")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}


// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// putAlbumByID updates an album from JSON received in the request body.
func putAlbumByID(c *gin.Context) {
    id := c.Param("id")
    var updatedAlbum album

    // Call BindJSON to bind the received JSON to
    // updatedAlbum.
    if err := c.BindJSON(&updatedAlbum); err != nil {
        return
    }

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, a := range albums {
        if a.ID == id {
            albums[i] = updatedAlbum
            c.IndentedJSON(http.StatusOK, updatedAlbum)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// patchAlbumByID partially updates an album from JSON received in the request body.
func patchAlbumByID(c *gin.Context) {
    id := c.Param("id")
    var updatedAlbum album

    // Call BindJSON to bind the received JSON to
    // updatedAlbum.
    if err := c.BindJSON(&updatedAlbum); err != nil {
        return
    }

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, a := range albums {
        if a.ID == id {
            if updatedAlbum.Title != "" {
                albums[i].Title = updatedAlbum.Title
            }
            if updatedAlbum.Artist != "" {
                albums[i].Artist = updatedAlbum.Artist
            }
            if updatedAlbum.Price != 0 {
                albums[i].Price = updatedAlbum.Price
            }
            c.IndentedJSON(http.StatusOK, albums[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// deleteAlbumByID deletes an album whose ID value matches the id
// parameter sent by the client.
func deleteAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, a := range albums {
        if a.ID == id {
            // Remove the album from the slice.
            albums = append(albums[:i], albums[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
