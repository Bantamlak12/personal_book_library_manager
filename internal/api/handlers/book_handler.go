package handlers

import (
	"errors"
	"net/http"
	"strconv"

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

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var body models.CreateBook
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "INVALID_REQUEST", "Invalid input data", err.Error())
		return
	}

	// Save the body
	data, err := h.bookService.CreateBK(&body)
	if err != nil {
		if errors.Is(err, service.ErrDuplicate) {
			utils.NewErrorResponse(c, http.StatusConflict, "DUPLICATE_RESOURCE", "A book with this ISBN already exists", "")
		} else {
			utils.NewErrorResponse(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to save the book", err.Error())
		}
		return
	}

	response := SuccessResponse{
		Status:  http.StatusOK,
		Message: "Book created successfully",
		Data:    data,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	title := c.Query("title")
	author := c.Query("author")
	rating, _ := strconv.ParseFloat(c.Query("rating"), 64)
	status := c.Query("status")

	response, err := h.bookService.GetBooks(page, limit, rating, title, author, status)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get books", err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *BookHandler) GetBookById(c *gin.Context) {}

func (h *BookHandler) GetBookByISBN(context *gin.Context) {}

func (h *BookHandler) UpdateBook(c *gin.Context) {}

func (h *BookHandler) DeleteBook(c *gin.Context) {}

func (h *BookHandler) UpdateBookStatus(c *gin.Context) {}

func (h *BookHandler) UpdateBookRating(c *gin.Context) {}
