package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/egotch/gotkobbler/internal/database"
	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check requests.
type HealthHandler struct {
	db *database.DB
}

// NewHealthHandler creates a new HealthHandler with the given database connection.
func NewHealthHandler(db *database.DB) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

func (h *HealthHandler) Check(c *gin.Context) {

	type response struct {
		Status  string `json:"status"`
		Database string `json:"database"`
		Message string `json:"message,omitempty"`
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := h.db.Ping(ctx)

	if err != nil {
		// Database is uwell
		c.JSON(http.StatusServiceUnavailable, response{
			Status: "unhealthy",
			Database: "disconnected",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response{
		Status:	"healthy",
		Database: "connected",
	})

}
