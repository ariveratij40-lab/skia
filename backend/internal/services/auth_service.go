package services

import (
	"errors"
	"os"
	"time"

	"github.com/ariveratij40-lab/skia/backend/internal/middleware"
	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

type RegisterRequest struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	TenantName string `json:"tenant_name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	User         *UserResponse `json:"user"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	TenantID string `json:"tenant_id"`
}

func (s *AuthService) Register(req *RegisterRequest) (*AuthResponse, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create tenant
	tenant := &models.Tenant{
		Name:   req.TenantName,
		Email:  req.Email,
		Plan:   "free",
		Status: "active",
	}

	if err := s.db.Create(tenant).Error; err != nil {
		return nil, err
	}

	// Create default role
	role := &models.Role{
		TenantID:    tenant.ID,
		Name:        "admin",
		Description: "Administrator",
	}

	if err := s.db.Create(role).Error; err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		TenantID: tenant.ID,
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedPassword),
		RoleID:   role.ID,
		Status:   "active",
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	// Generate tokens
	token, refreshToken, err := s.generateTokens(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User: &UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Name:     user.Name,
			TenantID: user.TenantID,
		},
	}, nil
}

func (s *AuthService) Login(req *LoginRequest) (*AuthResponse, error) {
	var user models.User

	if err := s.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate tokens
	token, refreshToken, err := s.generateTokens(&user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User: &UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Name:     user.Name,
			TenantID: user.TenantID,
		},
	}, nil
}

func (s *AuthService) generateTokens(user *models.User) (string, string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", "", errors.New("JWT_SECRET not set")
	}

	// Access token (15 minutes)
	accessClaims := &middleware.Claims{
		UserID:   user.ID,
		TenantID: user.TenantID,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	// Refresh token (7 days)
	refreshClaims := &middleware.Claims{
		UserID:   user.ID,
		TenantID: user.TenantID,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	claims := &middleware.Claims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	// Generate new access token
	accessClaims := &middleware.Claims{
		UserID:   claims.UserID,
		TenantID: claims.TenantID,
		Email:    claims.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}
