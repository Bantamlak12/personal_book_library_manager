package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	repo := &SQLiteRepository{}
	repo.InitDB()
	return repo
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
	 id TEXT PRIMARY KEY,
	 title TEXT NOT NULL,
	 author TEXT NOT NULL,
	 isbn TEXT,
	 status TEXT NOT NULL,
	 rating REAL,
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

func (r *SQLiteRepository) Create(book *models.CreateBook) error {
	query := `
	INSERT INTO books (id, title, author, isbn, status, rating, notes, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close() // Closes the statement after use

	book.Id = uuid.NewString()
	// Add timestamp
	if book.CreatedAt.IsZero() {
		book.CreatedAt = time.Now()
	}
	if book.UpdatedAt.IsZero() {
		book.UpdatedAt = time.Now()
	}

	_, err = stmt.Exec(book.Id, book.Title, book.Author, book.ISBN, book.Status, book.Rating, book.Notes, book.CreatedAt, book.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
