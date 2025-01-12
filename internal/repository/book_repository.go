package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func (r *SQLiteRepository) InitDB() {
	var err error
	r.db, err = sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatalln("Could not connnect to the database: " + err.Error())
	}

	// Set connection pool configurations
	r.db.SetMaxOpenConns(10)
	r.db.SetMaxIdleConns(5)

	r.createTable()
}

func (r *SQLiteRepository) createTable() {
	createBooksTable := `
	CREATE TABLE IF NOT EXISTS books (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 title TEXT NOT NULL,
	 author TEXT NOT NULL,
	 isbn TEXT,
	 status TEXT NOT NULL,
	 rating INTEGER,
	 notes TEXT,
	 created_at DATETIME NOT NULL,
	 updated_at DATETIME NOT NULL
	);
	 CREATE INDEX IF NOT EXISTS idx_books_status ON books(status);
	 CREATE INDEX IF NOT EXISTS idx_books_rating ON books(rating);
	`

	_, err := r.db.Exec(createBooksTable)
	if err != nil {
		log.Fatalln("Could not create books table. " + err.Error())
	}
}

func (r *SQLiteRepository) CloseDB() {
	if r.db != nil {
		r.db.Close()
	}
}
