package service

import (
	"errors"
	"time"

	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
	"github.com/Bantamlak12/personal_book_library_manager/internal/repository"
	"github.com/google/uuid"
)

// Defines the methods that are used to manage books
type BookService interface {
	CreateBK(book *models.CreateBook) (*models.CreateBook, error)
}

type bookService struct {
	bookRepo *repository.SQLiteRepository
}

// Initializes a new bookService instance
func NewBookService(repo *repository.SQLiteRepository) *bookService {
	return &bookService{
		bookRepo: repo,
	}
}

var ErrDuplicate = errors.New("a book with this ISBN already exists")

func (s *bookService) CreateBK(book *models.CreateBook) (*models.CreateBook, error) {
	// Check if the ISBN already exists
	exists, err := s.bookRepo.IsISBNExists(book.ISBN)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrDuplicate
	}

	book.Id = uuid.NewString()
	if book.CreatedAt.IsZero() {
		book.CreatedAt = time.Now()
	}
	if book.UpdatedAt.IsZero() {
		book.UpdatedAt = time.Now()
	}

	// Save the book to the repository
	err = s.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
