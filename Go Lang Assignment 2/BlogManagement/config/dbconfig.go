package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // SQLite driver
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./myblogs.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		timestamp TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return DB
}
