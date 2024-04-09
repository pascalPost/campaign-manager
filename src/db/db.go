package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	// run create statements
	CreateProjectsTable(db)
	CreateSettingsTable(db)

	return db
}

func DisconnectDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
