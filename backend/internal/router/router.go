package router

import (
	"github.com/egotch/gotkobbler/internal/database"
	"github.com/egotch/gotkobbler/internal/handlers"
	"github.com/egotch/gotkobbler/internal/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter initializes and returns a Gin router with applied middlewares.
func New(db *database.DB) *gin.Engine {

	// Init the router
	router := gin.Default()

	// Apply middlewares
	router.Use(middleware.CORSMiddleware())

	// Register routes
	registerRoutes(router, db)

	return router
}

// registerRoutes sets up the API routes and their corresponding handlers.
func registerRoutes(router *gin.Engine, db *database.DB) {

	//Handlers
	healthHandler := handlers.NewHealthHandler(db)

	api := router.Group("/api")
	{
		api.GET("/health", healthHandler.Check)
	}
}
