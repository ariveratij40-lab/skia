package handlers

import (
	"net/http"

	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/ariveratij40-lab/skia/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

// NodeHandler maneja operaciones de nodos
type NodeHandler struct {
	nodeRepo *repository.NodeRepository
}

// NewNodeHandler crea un nuevo handler de nodos
func NewNodeHandler(nodeRepo *repository.NodeRepository) *NodeHandler {
	return &NodeHandler{nodeRepo: nodeRepo}
}

// List maneja la lista de nodos
func (h *NodeHandler) List(c *gin.Context) {
	// TODO: Implementar
	c.JSON(http.StatusOK, models.ListResponse{
		Data: []models.Node{},
		Pagination: models.Pagination{
			Page:       1,
			Limit:      20,
			Total:      0,
			TotalPages: 0,
		},
	})
}

// Get maneja la obtención de un nodo
func (h *NodeHandler) Get(c *gin.Context) {
	id := c.Param("id")
	_ = id // TODO: Implementar

	c.JSON(http.StatusOK, models.Node{})
}

// GetByRFID maneja la búsqueda por RFID
func (h *NodeHandler) GetByRFID(c *gin.Context) {
	rfid := c.Param("rfid")
	_ = rfid // TODO: Implementar

	c.JSON(http.StatusOK, models.Node{})
}

// Create maneja la creación de un nodo
func (h *NodeHandler) Create(c *gin.Context) {
	var req models.CreateNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: struct {
				Code    string      `json:"code"`
				Message string      `json:"message"`
				Details interface{} `json:"details,omitempty"`
			}{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	// TODO: Implementar

	c.JSON(http.StatusCreated, models.Node{
		RFIDUID:    req.RFIDUID,
		PortNumber: req.PortNumber,
	})
}

// Update maneja la actualización de un nodo
func (h *NodeHandler) Update(c *gin.Context) {
	id := c.Param("id")
	_ = id // TODO: Implementar

	c.JSON(http.StatusOK, models.Node{})
}

// Delete maneja la eliminación de un nodo
func (h *NodeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	_ = id // TODO: Implementar

	c.Status(http.StatusNoContent)
}
