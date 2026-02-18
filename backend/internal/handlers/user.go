package handlers

import (
	"net/http"

	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/ariveratij40-lab/skia/backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserHandler maneja operaciones de usuarios
type UserHandler struct {
	userRepo *repository.UserRepository
}

// NewUserHandler crea un nuevo handler de usuarios
func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// List maneja la lista de usuarios
func (h *UserHandler) List(c *gin.Context) {
	// TODO: Implementar
	c.JSON(http.StatusOK, models.ListResponse{
		Data: []models.UserResponse{},
		Pagination: models.Pagination{
			Page:       1,
			Limit:      20,
			Total:      0,
			TotalPages: 0,
		},
	})
}

// Get maneja la obtención de un usuario
func (h *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")
	_ = id // TODO: Validar UUID y buscar usuario

	c.JSON(http.StatusOK, models.UserResponse{})
}

// Create maneja la creación de un usuario
func (h *UserHandler) Create(c *gin.Context) {
	var req models.CreateUserRequest
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

	// TODO: Implementar creación

	c.JSON(http.StatusCreated, models.UserResponse{
		Email: req.Email,
		Role:  req.Role,
	})
}

// Update maneja la actualización de un usuario
func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	_ = id // TODO: Validar UUID

	var req models.UpdateUserRequest
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

	// TODO: Implementar actualización

	c.JSON(http.StatusOK, models.UserResponse{})
}

// Delete maneja la eliminación de un usuario
func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	_ = id // TODO: Validar UUID y eliminar

	c.Status(http.StatusNoContent)
}

// parseUUID convierte string a UUID
func parseUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}
