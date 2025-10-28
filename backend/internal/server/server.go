package server

import (
	"context"
	"log"
	"net/http"

	"github.com/egotch/gotkobbler/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
	router	 *gin.Engine
}

func New(cfg *config.Config, router *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: router,
		},
		router: router,
	}
}

func (s *Server) Start(cfg *config.Config) {

	go func() {
		log.Printf("Starting server on port %s...", cfg.Port)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
		return err
	}

	log.Println("👋 Server exited")
	return nil
}
