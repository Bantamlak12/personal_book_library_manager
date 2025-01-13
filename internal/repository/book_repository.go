package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
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

// Check for exisitng book by its ISBN
func (r *SQLiteRepository) IsISBNExists(isbn string) (bool, string, error) {
	var existingISBN string
	err := r.db.QueryRow(`SELECT isbn FROM books WHERE isbn = ?`, isbn).Scan(&existingISBN)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// ISBN not found
			return false, "", nil
		}
		// Return any other error
		return false, "", err
	}
	return true, existingISBN, nil
}

// Create a new Book
func (r *SQLiteRepository) Create(book *models.CreateBook) error {
	query := `
	INSERT INTO books (id, title, author, isbn, status, rating, notes, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Closes the statement after use

	_, err = stmt.Exec(book.Id, book.Title, book.Author, book.ISBN, book.Status, book.Rating, book.Notes, book.CreatedAt, book.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to execute insert statement: %w", err)
	}

	return nil
}

// Get all Books
func (r *SQLiteRepository) SearchOrFilter(page, limit int, rating float64, title, author, status string) (*models.PaginatedResponse, error) {
	var books []*models.Book
	var conditions []string
	var args []interface{}

	// Base query
	query := "SELECT * FROM books"

	// Dynamic Conditions based on input
	if title != "" {
		conditions = append(conditions, "LOWER(title) LIKE ?")
		args = append(args, "%"+strings.ToLower(title)+"%")
	}
	if author != "" {
		conditions = append(conditions, "LOWER(author) LIKE ?")
		args = append(args, "%"+strings.ToLower(author)+"%")
	}
	if status != "" {
		conditions = append(conditions, "LOWER(status) = ?")
		args = append(args, status)
	}
	if rating > 0 {
		conditions = append(conditions, "rating >= ?")
		args = append(args, rating)
	}

	// Combine conditions
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Add pagination
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, (page-1)*limit)

	// Execute the query
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process the result
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.ISBN, &book.Status, &book.Rating, &book.Notes, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	// Count total records
	totalQuery := "SELECT COUNT(*) FROM books"
	if len(conditions) > 0 {
		totalQuery += " WHERE " + strings.Join(conditions, " AND ")
	}
	var total int
	err = r.db.QueryRow(totalQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, err
	}

	// Return paginated response
	return &models.PaginatedResponse{
		Status: http.StatusOK,
		Data:   books,
		Metadata: models.Metadata{
			Total: total,
			Page:  page,
			Limit: limit,
		},
	}, nil
}

// Get Book its Id

// Get Book by its ISBN
