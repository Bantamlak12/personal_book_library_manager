package handlers

import (
	"log"
	"net/http"

	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
	"github.com/Bantamlak12/personal_book_library_manager/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo *repository.SQLiteRepository
}

func NewHandler(repo *repository.SQLiteRepository) *Handler {
	return &Handler{Repo: repo}
}

func GetAllBooks(context *gin.Context) {}

func GetBookById(context *gin.Context) {}

func (h *Handler) CreateBook(context *gin.Context) {
	var data models.CreateBook
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the data
	err := h.Repo.Create(&data)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the book: " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Book created!", "Book": data})
}

func UpdateBook(context *gin.Context) {}

func DeleteBook(context *gin.Context) {}

func UpdateBookStatus(context *gin.Context) {}

func UpdateBookRating(context *gin.Context) {}
