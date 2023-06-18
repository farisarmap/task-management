package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func DatabaseConnection() {
	dbType := os.Getenv("DB_TYPE")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbType, dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbType, dbSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to database: ", dbType)

	DB = db
}
