package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func DatabaseConnection() {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:5432/tmdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	DB = db
}
