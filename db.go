package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func createDb() {
	db, _ := sql.Open("sqlite3", "./database.db")
	drop, _ := db.Prepare("DROP TABLE IF EXISTS games;")
	create, _ := db.Prepare("CREATE TABLE IF NOT EXISTS games (name TEXT, cover TEXT);")
	drop.Exec()
	create.Exec()
	drop.Close()
	create.Close()
	db.Close()
}

func insertIntoDb(name string, cover string) {
	db, _ := sql.Open("sqlite3", "./database.db")
	statement, _ := db.Prepare("INSERT INTO games (name, cover) VALUES (?, ?);")
	statement.Exec(name, cover)
	statement.Close()
	db.Close()
}
