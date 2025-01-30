package api

import (
	"github.com/Bantamlak12/personal_book_library_manager/internal/api/handlers"
	"github.com/Bantamlak12/personal_book_library_manager/internal/repository"
	"github.com/Bantamlak12/personal_book_library_manager/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repo *repository.SQLiteRepository) *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	// CORS configuration
	config := cors.Config{
		AllowOrigins:     []string{"https://personal-book-library-manager-sooty.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	// Apply CORS middleware
	router.Use(cors.New(config))

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
		bookRoutes.GET("/search", h.SearchBook) // Search from openlibrary
	}

	return router
}
