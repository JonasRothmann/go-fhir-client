package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Cache struct {
	db   *sql.DB
	lock sync.RWMutex
}

func NewCache() *Cache {
	var err error
	db, err := sql.Open("sqlite3", "temp.db")
	if err != nil {
		log.Fatalf("Failed to open SQLite DB: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS code_systems (
			system_url TEXT PRIMARY KEY,
			data BLOB,
			last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return &Cache{
		db: db,
	}
}

var errCacheMiss = errors.New("cache miss")

func (db *Cache) Get(systemURL string) (result CodeSystem, err error) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	var blobData []byte
	err = db.db.QueryRow("SELECT data FROM code_systems WHERE system_url = ?", systemURL).Scan(&blobData)
	if err == sql.ErrNoRows {
		return result, errCacheMiss
	} else if err != nil {
		return result, err
	}

	err = json.Unmarshal(blobData, &result)
	return result, err
}

func (db *Cache) Set(systemURL string, value CodeSystem) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	blobData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = db.db.Exec(
		"INSERT INTO code_systems (system_url, data) VALUES (?, ?) ON CONFLICT(system_url) DO UPDATE SET data = ?",
		systemURL, blobData, blobData,
	)
	return err
}

func (db *Cache) Close() {
	db.db.Close()
}
