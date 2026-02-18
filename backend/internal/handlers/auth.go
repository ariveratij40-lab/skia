package handlers

import (
	"net/http"
	"time"

	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/ariveratij40-lab/skia/backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthHandler maneja operaciones de autenticación
type AuthHandler struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

// NewAuthHandler crea un nuevo handler de autenticación
func NewAuthHandler(userRepo *repository.UserRepository, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Login maneja el inicio de sesión
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
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

	// TODO: Implementar lógica de login
	// 1. Buscar usuario por email
	// 2. Verificar contraseña con bcrypt
	// 3. Generar tokens JWT
	// 4. Actualizar last_login
	// 5. Retornar respuesta

	c.JSON(http.StatusOK, models.LoginResponse{
		AccessToken:  "token",
		RefreshToken: "refresh",
		ExpiresIn:    900,
	})
}

// Refresh maneja la renovación de tokens
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req models.RefreshRequest
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

	// TODO: Implementar refresh token

	c.JSON(http.StatusOK, gin.H{
		"access_token": "new_token",
		"expires_in":   900,
	})
}

// Logout maneja el cierre de sesión
func (h *AuthHandler) Logout(c *gin.Context) {
	// TODO: Implementar logout (invalidar token)
	c.Status(http.StatusNoContent)
}

// generateTokens genera access y refresh tokens
func (h *AuthHandler) generateTokens(userID, tenantID, email, role string) (string, string, error) {
	// Access token
	accessClaims := jwt.MapClaims{
		"user_id":    userID,
		"tenant_id":  tenantID,
		"email":      email,
		"role":       role,
		"exp":        time.Now().Add(15 * time.Minute).Unix(),
		"iat":        time.Now().Unix(),
		"type":       "access",
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString([]byte(h.jwtSecret))
	if err != nil {
		return "", "", err
	}

	// Refresh token
	refreshClaims := jwt.MapClaims{
		"user_id":   userID,
		"tenant_id": tenantID,
		"exp":       time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":       time.Now().Unix(),
		"type":      "refresh",
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString([]byte(h.jwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}

// hashPassword genera hash de contraseña
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// checkPassword verifica contraseña
func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
