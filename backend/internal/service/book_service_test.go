package service

import (
	"errors"
	"testing"
	"time"

	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
	"github.com/Bantamlak12/personal_book_library_manager/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the book repository
type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) IsISBNExists(isbn string) (bool, string, error) {
	args := m.Called(isbn)
	return args.Bool(0), args.String(1), args.Error(2)
}

func (m *MockBookRepository) Create(book *models.CreateBook) error {
	args := m.Called(book)
	return args.Error(0)
}

func TestCreateBK(t *testing.T) {
	mockRepo := new(MockBookRepository)
	bookService := NewBookService(mockRepo)

	// Sample book data
	book := &models.CreateBook{
		Title:  "Test title",
		Author: "Test author",
		ISBN:   "0396796424",
		Status: "unread",
		Rating: 4.5,
		Notes:  "Test note",
	}

	t.Run("Success", func(t *testing.T) {
		// Mock repository behavior
		mockRepo.On("IsISBNExists", book.ISBN).Return(false, "", nil)
		mockRepo.On("Create", mock.AnythingOfType("*models.CreateBook")).Return(nil)

		// Call the method
		result, err := bookService.CreateBK(book)

		// Assert the method
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.Id)
		assert.WithinDuration(t, time.Now(), result.CreatedAt, time.Second)
		assert.WithinDuration(t, time.Now(), result.UpdatedAt, time.Second)

		// Verify mocks
		mockRepo.AssertExpectations(t)
	})

	t.Run("Duplicate ISBN", func(t *testing.T) {
		// Mock repository behaviour
		mockRepo.On("IsISBNExists", book.ISBN).Return(true, "", nil)

		// Call the method
		result, err := bookService.CreateBK(book)

		// Assert results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, repository.ErrDuplicate, err)

		// Verify mocks
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		// Mock repository behaviour
		mockRepo.On("IsISBNExists", book.ISBN).Return(false, "", errors.New("database error"))

		// Call the method
		result, err := bookService.CreateBK(book)

		// Assert results
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err, err.Error(), "database error")

		// Verify mocks
		mockRepo.AssertExpectations(t)
	})
}
