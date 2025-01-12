package handlers

import (
	"net/http"
	"time"

	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(context *gin.Context) {}

func GetBookById(context *gin.Context) {}

func CreateBook(context *gin.Context) {
	var data models.CreateBook
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.Id = "1"
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	context.JSON(http.StatusCreated, gin.H{"message": "Book created!", "Book": data})
}

func UpdateBook(context *gin.Context) {}

func DeleteBook(context *gin.Context) {}

func UpdateBookStatus(context *gin.Context) {}

func UpdateBookRating(context *gin.Context) {}
