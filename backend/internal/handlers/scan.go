package handlers

import (
	"net/http"

	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/ariveratij40-lab/skia/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

// ScanHandler maneja operaciones de escaneos
type ScanHandler struct {
	scanRepo *repository.ScanRepository
	nodeRepo *repository.NodeRepository
}

// NewScanHandler crea un nuevo handler de escaneos
func NewScanHandler(scanRepo *repository.ScanRepository, nodeRepo *repository.NodeRepository) *ScanHandler {
	return &ScanHandler{
		scanRepo: scanRepo,
		nodeRepo: nodeRepo,
	}
}

// List maneja la lista de escaneos
func (h *ScanHandler) List(c *gin.Context) {
	// TODO: Implementar
	c.JSON(http.StatusOK, models.ListResponse{
		Data: []models.Scan{},
		Pagination: models.Pagination{
			Page:       1,
			Limit:      20,
			Total:      0,
			TotalPages: 0,
		},
	})
}

// Create maneja la creación de un escaneo
func (h *ScanHandler) Create(c *gin.Context) {
	var req models.CreateScanRequest
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
	// 1. Buscar nodo por RFID
	// 2. Crear registro de escaneo
	// 3. Actualizar last_scan del nodo
	// 4. Retornar información del nodo

	c.JSON(http.StatusCreated, models.ScanResponse{
		Scan: models.Scan{
			ScanType: req.ScanType,
			Notes:    req.Notes,
		},
	})
}

// Dashboard maneja las métricas del dashboard
func (h *ScanHandler) Dashboard(c *gin.Context) {
	// TODO: Implementar
	c.JSON(http.StatusOK, models.DashboardResponse{})
}

// Activity maneja el reporte de actividad
func (h *ScanHandler) Activity(c *gin.Context) {
	// TODO: Implementar
	c.JSON(http.StatusOK, []map[string]interface{}{})
}
