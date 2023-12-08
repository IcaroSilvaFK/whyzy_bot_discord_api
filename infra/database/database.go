package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDatabaseConn() *sql.DB {

	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	updateDatabase(db)

	return db
}

func updateDatabase(db *sql.DB) {

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS latest_animes (mal_id TEXT, url TEXT, title TEXT, rank INTEGER, image_url TEXT)")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS recommendations (mal_id INTEGER,content TEXT, title TEXT, image_url TEXT, url TEXT) ")

	if err != nil {
		panic(err)
	}
}
