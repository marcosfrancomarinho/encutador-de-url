package database

import (
	"database/sql"
	"fmt"
	_"modernc.org/sqlite"
	"log"
	"os"
	"sync"
)

var (
	DB   *sql.DB
	once sync.Once
)

func InitDataBase() *sql.DB {
	once.Do(func() {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		db, err := sql.Open("sqlite", fmt.Sprintf("%s/../../database.sqlite", path))
		if err != nil {
			log.Fatal(err)
		}

		DB = db

		if err := CreateTable(); err != nil {
			log.Fatal(err)
		}
	})
	return DB
}

func CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id TEXT PRIMARY KEY,
		short_id TEXT UNIQUE,
		long_url TEXT
	);`
	_, err := DB.Exec(query)
	return err
}
