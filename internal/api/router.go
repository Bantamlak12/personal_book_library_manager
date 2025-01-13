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
		// bookRoutes.GET("", handlers.GetAllBooks)
		// bookRoutes.GET("/:id", handlers.GetBookById)
		// bookRoutes.GET("/:isbn", handlers.GetBookByISBN)
		// bookRoutes.GET("/search", handlers.SearchBooks)
		bookRoutes.POST("", h.CreateBook)
		// bookRoutes.PUT("/:id", handlers.UpdateBook)
		// bookRoutes.DELETE("/:id", handlers.DeleteBook)
		// bookRoutes.PATCH("/:id/status", handlers.UpdateBookStatus)
		// bookRoutes.PATCH("/:id/rating", handlers.UpdateBookRating)
	}

	return router
}
