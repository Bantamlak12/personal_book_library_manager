package api

import (
	"github.com/Bantamlak12/personal_book_library_manager/internal/api/handlers"
	"github.com/Bantamlak12/personal_book_library_manager/internal/repository"
	"github.com/Bantamlak12/personal_book_library_manager/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repo *repository.SQLiteRepository) *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	// Create the book service and pass the repository
	bookService := service.NewBookService(repo)

	// Create BookHandler and pass the bookService
	h := handlers.NewBookHandler(bookService)

	// Book endpoints
	bookRoutes := router.Group("/api/v1/books")
	{
		bookRoutes.GET("", h.GetAllBooks)
		bookRoutes.GET("/:id", h.GetBookById)
		bookRoutes.POST("", h.CreateBook)
		bookRoutes.PUT("/:id", h.UpdateBook)
		bookRoutes.DELETE("/:id", h.DeleteBook)
		bookRoutes.PATCH("/:id/status", h.UpdateBookStatus)
		bookRoutes.PATCH("/:id/rating", h.UpdateBookRating)
		bookRoutes.GET("/search", h.SearchBook) // Search from openlibrary
	}

	return router
}
