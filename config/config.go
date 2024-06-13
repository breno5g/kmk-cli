package config

import (
	"database/sql"
	"fmt"
)

var (
	logger *Logger
	db     *sql.DB
	paths  *Paths
)

func Init() error {
	var err error

	err = LoadEnv()
	if err != nil {
		return fmt.Errorf("error loading env: %v", err)
	}

	db, err = InitilizeSqlite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
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

func GetPaths() *Paths {
	return paths
}
