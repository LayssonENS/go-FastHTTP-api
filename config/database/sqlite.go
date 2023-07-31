package database

import (
	"database/sql"
	"fmt"
	"github.com/LayssonENS/go-FastHTTP-api/config"
	"log"
)

func NewSqliteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.GetEnv().DbConfig.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the SQLite database.")
	return db, nil
}
