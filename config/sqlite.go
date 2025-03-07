package config

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitilizeSqlite() (*sql.DB, error) {
	DB_PATH := os.Getenv("DB_PATH")
	db, err := sql.Open("sqlite3", DB_PATH)

	return db, err
}
