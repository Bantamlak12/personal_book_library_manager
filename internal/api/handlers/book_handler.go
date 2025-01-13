package handlers

import (
	"errors"
	"net/http"

	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
	"github.com/Bantamlak12/personal_book_library_manager/internal/service"
	"github.com/Bantamlak12/personal_book_library_manager/internal/utils"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService service.BookService
}

// Accept the BookService instance and assign it to the BookHandler
func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func GetAllBooks(context *gin.Context) {}

func GetBookById(context *gin.Context) {}

func (h *BookHandler) CreateBook(context *gin.Context) {
	var body models.CreateBook
	if err := context.ShouldBindJSON(&body); err != nil {
		utils.NewErrorResponse(context, http.StatusBadRequest, "INVALID_REQUEST", "Invalid input data", err.Error())
		return

	}

	// Save the body
	data, err := h.bookService.CreateBK(&body)
	if err != nil {
		if errors.Is(err, service.ErrDuplicate) {
			utils.NewErrorResponse(context, http.StatusConflict, "DUPLICATE_RESOURCE", "A book with this ISBN already exists", "")
		} else {
			utils.NewErrorResponse(context, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to save the book", err.Error())
		}
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Book created!", "Book": data})
}

func UpdateBook(context *gin.Context) {}

func DeleteBook(context *gin.Context) {}

func UpdateBookStatus(context *gin.Context) {}

func UpdateBookRating(context *gin.Context) {}
