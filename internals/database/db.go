package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() error {
	log.Println("Trying to connect to the Database")

	godotenv.Load(".env")

	// db_url := os.Getenv("DATABASE_URL")
	db_url := "postgres://kishu:kishu@db/blogdb?sslmode=disable"

	if db_url == "" {
		return fmt.Errorf("Database_URL environment variable is not set")
	}

	var err error
	db, err = sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("Not able to open database connection", err)
	}
	// defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Ping to check the connection
	retries := 5
	for retries > 0 {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Database ping failed, retrying... (%d retries left)", retries)
		time.Sleep(2 * time.Second)
		retries--
	}

	if err != nil {
		return fmt.Errorf("could not connect to the database after retries: %w", err)
	}

	log.Println("Successfully connected to the Database")
	return nil
}

func GetDB() *sql.DB {
	return db
}
