package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// connectDB initializes the database connection.
func connectDB() error {
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	conn, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		conn.Close()
		return fmt.Errorf("failed to ping the database: %w", err)
	}

	db = conn
	return nil
}

// fetchAlbumsFromDatabase fetches all albums from the database.
func fetchAlbumsFromDatabase() ([]album, error) {
	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch albums from the database: %w", err)
	}
	defer rows.Close()

	var albums []album
	for rows.Next() {
		var a album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan album row: %w", err)
		}
		albums = append(albums, a)
	}

	return albums, nil
}

// fetchAlbumFromDatabase retrieves a specific album by its ID from the database.
func fetchAlbumFromDatabase(id string) (album, error) {
	var a album
	err := db.QueryRow("SELECT * FROM albums WHERE id=$1", id).Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return album{}, fmt.Errorf("album not found")
		}
		return album{}, fmt.Errorf("failed to fetch album from the database: %w", err)
	}

	return a, nil
}

// insertAlbumIntoDatabase inserts a new album into the database.
func insertAlbumIntoDatabase(a album) error {
	_, err := db.Exec("INSERT INTO albums (id, title, artist, price) VALUES ($1, $2, $3, $4)", a.ID, a.Title, a.Artist, a.Price)
	if err != nil {
		return fmt.Errorf("failed to insert album into the database: %w", err)
	}
	return nil
}

// updateAlbumInDatabase updates an existing album in the database.
func updateAlbumInDatabase(id string, a album) error {
	_, err := db.Exec("UPDATE albums SET title=$1, artist=$2, price=$3 WHERE id=$4", a.Title, a.Artist, a.Price, id)
	if err != nil {
		return fmt.Errorf("failed to update album in the database: %w", err)
	}
	return nil
}


// updateAlbumFieldsInDatabase updates specific fields of an existing album in the database.
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
	if err != nil {
		return fmt.Errorf("failed to update album fields in the database: %w", err)
	}
	return nil
}

// deleteAlbumFromDatabase deletes an album from the database.
func deleteAlbumFromDatabase(id string) error {
	_, err := db.Exec("DELETE FROM albums WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to delete album from the database: %w", err)
	}
	return nil
}
