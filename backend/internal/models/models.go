package models

import (
	"time"

	"github.com/google/uuid"
)

// Tenant representa una organización
type Tenant struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Slug      string    `json:"slug" db:"slug"`
	DBSchema  string    `json:"db_schema" db:"db_schema"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	Settings  JSONB     `json:"settings" db:"settings"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// User representa un usuario del sistema
type User struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	TenantID     uuid.UUID  `json:"tenant_id" db:"tenant_id"`
	Email        string     `json:"email" db:"email"`
	PasswordHash string     `json:"-" db:"password_hash"`
	Role         string     `json:"role" db:"role"`
	FirstName    string     `json:"first_name" db:"first_name"`
	LastName     string     `json:"last_name" db:"last_name"`
	Phone        string     `json:"phone" db:"phone"`
	IsActive     bool       `json:"is_active" db:"is_active"`
	LastLogin    *time.Time `json:"last_login,omitempty" db:"last_login"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

// UserResponse es la versión segura de User para respuestas JSON
type UserResponse struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Phone     string     `json:"phone"`
	IsActive  bool       `json:"is_active"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

// ToResponse convierte User a UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Role:      u.Role,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		IsActive:  u.IsActive,
		LastLogin: u.LastLogin,
		CreatedAt: u.CreatedAt,
	}
}

// Rack representa un rack de cableado
type Rack struct {
	ID          uuid.UUID `json:"id" db:"id"`
	TenantID    uuid.UUID `json:"tenant_id" db:"tenant_id"`
	Name        string    `json:"name" db:"name"`
	Location    string    `json:"location" db:"location"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// PatchPanel representa un panel de parcheo
type PatchPanel struct {
	ID          uuid.UUID `json:"id" db:"id"`
	TenantID    uuid.UUID `json:"tenant_id" db:"tenant_id"`
	RackID      uuid.UUID `json:"rack_id" db:"rack_id"`
	Name        string    `json:"name" db:"name"`
	PortCount   int       `json:"port_count" db:"port_count"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Node representa un nodo/puerto con RFID
type Node struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	TenantID     uuid.UUID  `json:"tenant_id" db:"tenant_id"`
	PatchPanelID uuid.UUID  `json:"patch_panel_id" db:"patch_panel_id"`
	RFIDUID      string     `json:"rfid_uid" db:"rfid_uid"`
	PortNumber   int        `json:"port_number" db:"port_number"`
	Status       string     `json:"status" db:"status"`
	Description  string     `json:"description" db:"description"`
	LastScanAt   *time.Time `json:"last_scan_at,omitempty" db:"last_scan_at"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

// NodeDetail incluye información relacionada
type NodeDetail struct {
	Node       Node       `json:"node"`
	PatchPanel PatchPanel `json:"patch_panel"`
	Rack       Rack       `json:"rack"`
}

// Cable representa un cable de red
type Cable struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	TenantID    uuid.UUID  `json:"tenant_id" db:"tenant_id"`
	NodeAID     uuid.UUID  `json:"node_a_id" db:"node_a_id"`
	NodeBID     *uuid.UUID `json:"node_b_id,omitempty" db:"node_b_id"`
	CableType   string     `json:"cable_type" db:"cable_type"`
	Length      *float64   `json:"length,omitempty" db:"length"`
	Color       string     `json:"color" db:"color"`
	Label       string     `json:"label" db:"label"`
	Status      string     `json:"status" db:"status"`
	Destination string     `json:"destination" db:"destination"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

// Scan representa un escaneo RFID
type Scan struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	TenantID  uuid.UUID  `json:"tenant_id" db:"tenant_id"`
	NodeID    uuid.UUID  `json:"node_id" db:"node_id"`
	UserID    uuid.UUID  `json:"user_id" db:"user_id"`
	ScanType  string     `json:"scan_type" db:"scan_type"`
	Latitude  *float64   `json:"latitude,omitempty" db:"latitude"`
	Longitude *float64   `json:"longitude,omitempty" db:"longitude"`
	Notes     string     `json:"notes" db:"notes"`
	ScannedAt time.Time  `json:"scanned_at" db:"scanned_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
}

// ScanResponse incluye información relacionada
type ScanResponse struct {
	Scan      Scan         `json:"scan"`
	Node      Node         `json:"node"`
	User      UserResponse `json:"user"`
	PatchPanel PatchPanel  `json:"patch_panel,omitempty"`
	Rack      Rack         `json:"rack,omitempty"`
}

// ScanHistory representa el historial de cambios de un nodo
type ScanHistory struct {
	ID             uuid.UUID `json:"id" db:"id"`
	TenantID       uuid.UUID `json:"tenant_id" db:"tenant_id"`
	NodeID         uuid.UUID `json:"node_id" db:"node_id"`
	ScanID         *uuid.UUID `json:"scan_id,omitempty" db:"scan_id"`
	PreviousStatus string    `json:"previous_status" db:"previous_status"`
	NewStatus      string    `json:"new_status" db:"new_status"`
	ChangedBy      uuid.UUID `json:"changed_by" db:"changed_by"`
	ChangedAt      time.Time `json:"changed_at" db:"changed_at"`
	Reason         string    `json:"reason" db:"reason"`
}

// JSONB tipo para campos JSON
type JSONB map[string]interface{}

// LoginRequest representa la petición de login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse representa la respuesta de login
type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresIn    int          `json:"expires_in"`
	User         UserResponse `json:"user"`
}

// RefreshRequest representa la petición de refresh token
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// CreateUserRequest representa la petición de creación de usuario
type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	Role      string `json:"role" binding:"required,oneof=admin technician viewer"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

// UpdateUserRequest representa la petición de actualización de usuario
type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	IsActive  *bool  `json:"is_active,omitempty"`
}

// CreateNodeRequest representa la petición de creación de nodo
type CreateNodeRequest struct {
	PatchPanelID uuid.UUID `json:"patch_panel_id" binding:"required"`
	RFIDUID      string    `json:"rfid_uid" binding:"required"`
	PortNumber   int       `json:"port_number" binding:"required,min=1"`
	Description  string    `json:"description"`
}

// CreateScanRequest representa la petición de creación de escaneo
type CreateScanRequest struct {
	RFIDCode  string   `json:"rfid_code" binding:"required"`
	ScanType  string   `json:"scan_type" binding:"required,oneof=check_in check_out audit"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Notes     string   `json:"notes"`
}

// DashboardResponse representa las métricas del dashboard
type DashboardResponse struct {
	Summary struct {
		TotalRacks         int `json:"total_racks"`
		TotalPanels        int `json:"total_panels"`
		TotalNodes         int `json:"total_nodes"`
		ActiveNodes        int `json:"active_nodes"`
		InactiveNodes      int `json:"inactive_nodes"`
		MaintenanceNodes   int `json:"maintenance_nodes"`
	} `json:"summary"`
	RecentScans []RecentScan `json:"recent_scans"`
	ScansByDay  []ScansByDay `json:"scans_by_day"`
}

// RecentScan representa un escaneo reciente
type RecentScan struct {
	ID        uuid.UUID `json:"id"`
	NodeRFID  string    `json:"node_rfid"`
	UserName  string    `json:"user_name"`
	ScannedAt time.Time `json:"scanned_at"`
}

// ScansByDay representa escaneos agrupados por día
type ScansByDay struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

// Pagination representa la información de paginación
type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// ListResponse es una respuesta paginada genérica
type ListResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// ErrorResponse representa un error de la API
type ErrorResponse struct {
	Error struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Details interface{} `json:"details,omitempty"`
	} `json:"error"`
}
