package cm

import (
	"database/sql"
	log2 "github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log2.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	return db
}
