package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://amarjit:amarjit@localhost/sendgriddb?sslmode=disable")
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
