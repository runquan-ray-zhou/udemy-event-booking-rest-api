package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// set database connection
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // should end with .db  this opens the database and go can make connections

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

// create table
func createTables() {
	createEventsTables := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	_, err := DB.Exec(createEventsTables) // executes sql query statement

	if err != nil {
		panic("Could not create events table.")
	}
}