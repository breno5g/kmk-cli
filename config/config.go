package config

import (
	"database/sql"
	"fmt"
)

var (
	logger *Logger
	db     *sql.DB
)

func Init() error {
	var err error
	db, err = InitilizeSqlite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	err = LoadEnv()
	if err != nil {
		return fmt.Errorf("error loading env: %v", err)
	}

	return nil

}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

func GetDB() *sql.DB {
	return db
}
