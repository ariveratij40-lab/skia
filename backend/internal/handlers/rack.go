package handlers

import (
	"net/http"

	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/ariveratij40-lab/skia/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

// RackHandler maneja operaciones de racks
type RackHandler struct {
	rackRepo *repository.RackRepository
}

// NewRackHandler crea un nuevo handler de racks
func NewRackHandler(trackRepo *repository.RackRepository) *RackHandler {
	return &RackHandler{trackRepo: trackRepo}
}

// List maneja la lista de racks
func (h *RackHandler) List(c *gin.Context) {
	// TODO: Implementar
	c.JSON(http.StatusOK, models.ListResponse{
		Data: []models.Rack{},
		Pagination: models.Pagination{
			Page:       1,
			Limit:      20,
			Total:      0,
			TotalPages: 0,
		},
	})
}

// Get maneja la obtención de un rack
func (h *RackHandler) Get(c *gin.Context) {
	id := c.Param("id")
	_ = id // TODO: Implementar

	c.JSON(http.StatusOK, models.Rack{})
}

// Create maneja la creación de un rack
func (h *RackHandler) Create(c *gin.Context) {
	// TODO: Implementar
	c.JSON(http.StatusCreated, models.Rack{})
}

// Update maneja la actualización de un rack
func (h *RackHandler) Update(c *gin.Context) {
	id := c.Param("id")
	_ = id // TODO: Implementar

	c.JSON(http.StatusOK, models.Rack{})
}

// Delete maneja la eliminación de un rack
func (h *RackHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	_ = id // TODO: Implementar

	c.Status(http.StatusNoContent)
}
