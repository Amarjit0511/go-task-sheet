package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Data structure for songs albums
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db *sql.DB

func main() {
	// Loading the environment variable from config.env
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Getting the database connection string from the env file
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	// Starting the database connection using the obtained connection string
	db, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	// Enable CORS for all routes
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	router.GET("/albums/get", getAlbums)
	router.GET("/albums/get/:id", getAlbumByID)
	router.POST("/albums/post", postAlbums)
	router.PUT("/albums/put/:id", putAlbumByID)
	router.DELETE("/albums/delete/:id", deleteAlbumByID)

	sslCertPath := "/Users/amarjitkumar/Desktop/key/ssl_cert.pem"
	sslKeyPath := "/Users/amarjitkumar/Desktop/key/ssl_key.pem"

	router.RunTLS("localhost:8443", sslCertPath, sslKeyPath)
}

// This will return all the albums from the pg database
func getAlbums(c *gin.Context) {
	albums := fetchAlbumsFromDatabase()
	c.IndentedJSON(http.StatusOK, albums)
}

// This will add an album to the pg database
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Binding the received json to newAlbum for structure validation
	if err := c.BindJSON(&newAlbum); err != nil {
		log.Printf("Error binding JSON for new album: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Sorry": "invalid data"})
		return
	}

	// Finally inserting the new album in the database
	if err := insertAlbumIntoDatabase(newAlbum); err != nil {
		log.Printf("Error inserting new album into database: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Sorry": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Getting albums by id value
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Getting the album from the database for a particular id
	a, err := fetchAlbumFromDatabase(id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Sorry": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Sorry": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, a)
}

// Putting an album by id
func putAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	// Binding the JSON received in the request with the actually defined struct
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Sorry": "invalid data"})
		return
	}

	// First checking if the album with the given ID exists or not
	_, err := fetchAlbumFromDatabase(id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Sorry": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Sorry": "internal server error"})
		return
	}

	// If everything is right, updating the album in the database
	if err := updateAlbumInDatabase(updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Sorry": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

// Deleting album by id
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Checking if the album with a particular id exists or not
	_, err := fetchAlbumFromDatabase(id)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Sorry": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Sorry": "internal server error"})
		return
	}

	// If found in the database, deleting it
	if err := deleteAlbumFromDatabase(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Sorry": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Successfull": "album deleted"})
}

// Functions for interacting with the database
func fetchAlbumsFromDatabase() []album {
	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		log.Fatalf("Error finding albums in database: %v", err)
	}
	defer rows.Close()

	var albums []album
	for rows.Next() {
		var a album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			log.Fatalf("Error searching row: %v", err)
		}
		albums = append(albums, a)
	}

	return albums
}

func fetchAlbumFromDatabase(id string) (album, error) {
	var a album
	err := db.QueryRow("SELECT * FROM albums WHERE id=$1", id).Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
	if err != nil {
		log.Printf("Error finding album with ID %s in database: %v", id, err)
		return album{}, err
	}
	return a, nil
}

func insertAlbumIntoDatabase(a album) error {
	_, err := db.Exec("INSERT INTO albums (id, title, artist, price) VALUES ($1, $2, $3, $4)", a.ID, a.Title, a.Artist, a.Price)
	if err != nil {
		log.Printf("Error inserting album into the database: %v", err)
	}
	return err
}

func updateAlbumInDatabase(a album) error {
	_, err := db.Exec("UPDATE albums SET title=$1, artist=$2, price=$3 WHERE id=$4", a.Title, a.Artist, a.Price, a.ID)
	if err != nil {
		log.Printf("Error updating album in the database: %v", err)
	}
	return err
}

func deleteAlbumFromDatabase(id string) error {
	_, err := db.Exec("DELETE FROM albums WHERE id=$1", id)
	if err != nil {
		log.Printf("Error deleting album from the database: %v", err)
	}
	return err
}
