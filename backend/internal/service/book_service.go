package service

import (
	"time"

	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
	"github.com/Bantamlak12/personal_book_library_manager/internal/repository"
	"github.com/google/uuid"
)

// Defines the methods that are used to manage books
type BookService interface {
	CreateBK(book *models.CreateBook) (*models.CreateBook, error)
	GetBooks(page, limit int, rating float64, title, author, status string) (*models.PaginatedResponse, error)
	GetBookById(id string) (*models.Book, error)
	UpdateBook(id string, book *models.Book) (*models.Book, error)
	DeleteBook(id string) error
	SearchFromOpenLibrary(page, limit int, title, author, isbn string) (*OpenLibrarySearchResponse, error)
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

func (s *bookService) CreateBK(book *models.CreateBook) (*models.CreateBook, error) {
	// Check if the ISBN already exists
	exists, _, err := s.bookRepo.IsISBNExists(book.ISBN)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, repository.ErrDuplicate
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

// page and limit: For pagination
// rating and status: For filtration
// tile, author, and ISBN: For searching
func (s *bookService) GetBooks(page, limit int, rating float64, title, author, status string) (*models.PaginatedResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 20 {
		limit = 12
	}

	book, err := s.bookRepo.SearchOrFilter(page, limit, rating, title, author, status)
	return book, err
}

func (s *bookService) GetBookById(id string) (*models.Book, error) {
	book, err := s.bookRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *bookService) UpdateBook(id string, book *models.Book) (*models.Book, error) {
	updatedBook, err := s.bookRepo.UpdateBK(id, book)
	if err != nil {
		return nil, err
	}

	return updatedBook, nil
}

func (s *bookService) DeleteBook(id string) error {
	err := s.bookRepo.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}
