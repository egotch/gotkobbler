package app

import (
	"context"
	"fmt"
	"log"

	"github.com/egotch/gotkobbler/backend/internal/config"
	"github.com/egotch/gotkobbler/backend/internal/database"
	"github.com/egotch/gotkobbler/backend/internal/router"
	"github.com/egotch/gotkobbler/backend/internal/server"
)

type App struct {
	Config *config.Config
	DB *database.DB
	Server *server.Server
}

// New initializes the application with its components
func New(cfg *config.Config) (*App, error) {
    // Initialize database
    db, err := database.New(cfg)
    if err != nil {
        return nil, fmt.Errorf("database init failed: %w", err)
    }
    
    // Create router with db
    router := router.New(db)
    
    // Create server
    srv := server.New(cfg, router)
    
    return &App{
        Config: cfg,
        DB:     db,
        Server: srv,
    }, nil
}

// Run starts the application Server
func (a *App) Run() {
	a.Server.Start(a.Config)

}

// Shutdown gracefully shuts down the application
func (a *App) Shutdown(ctx context.Context) error {
    log.Println("Shutting down application...")
    
    // Shutdown server first
    if err := a.Server.Shutdown(ctx); err != nil {
        return fmt.Errorf("server shutdown failed: %w", err)
    }
    
    // Then close database
    a.DB.Close()
    
    log.Println("Application shutdown complete")
    return nil
}
