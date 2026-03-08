package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FASE 0: Multi-Tenancy & Billing

// Tenant represents a customer organization
type Tenant struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Plan      string    `json:"plan"` // free, starter, professional, enterprise
	Status    string    `json:"status"` // active, suspended, cancelled
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     []User    `gorm:"foreignKey:TenantID" json:"-"`
}

// User represents a user in the system
type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	TenantID  string    `json:"tenant_id"`
	Email     string    `gorm:"uniqueIndex:idx_email_tenant,composite:tenant_id" json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	RoleID    string    `json:"role_id"`
	Status    string    `json:"status"` // active, inactive, suspended
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tenant    Tenant    `gorm:"foreignKey:TenantID" json:"-"`
	Role      Role      `gorm:"foreignKey:RoleID" json:"-"`
}

// Role represents user roles (admin, user, viewer, etc.)
type Role struct {
	ID          string       `gorm:"primaryKey" json:"id"`
	TenantID    string       `json:"tenant_id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `gorm:"many2many:role_permissions" json:"-"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// Permission represents system permissions
type Permission struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// AuditLog tracks all changes
type AuditLog struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	TenantID  string    `json:"tenant_id"`
	UserID    string    `json:"user_id"`
	Action    string    `json:"action"` // CREATE, UPDATE, DELETE
	Entity    string    `json:"entity"` // User, Device, etc.
	EntityID  string    `json:"entity_id"`
	Changes   string    `json:"changes"` // JSON string of changes
	IPAddress string    `json:"ip_address"`
	CreatedAt time.Time `json:"created_at"`
}

// Subscription tracks billing information
type Subscription struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	TenantID      string    `json:"tenant_id"`
	StripeID      string    `json:"stripe_id"`
	Plan          string    `json:"plan"`
	Status        string    `json:"status"` // active, past_due, cancelled
	CurrentPeriod time.Time `json:"current_period"`
	NextBilling   time.Time `json:"next_billing"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// FASE 1: DCIM Inventory Management

// Site represents a physical location (data center, office, etc.)
type Site struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	TenantID  string    `json:"tenant_id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Status    string    `json:"status"` // active, inactive, maintenance
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Buildings []Building `gorm:"foreignKey:SiteID" json:"-"`
}

// Building represents a building within a site
type Building struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	SiteID    string    `json:"site_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Floors    []Floor   `gorm:"foreignKey:BuildingID" json:"-"`
}

// Floor represents a floor in a building
type Floor struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	BuildingID string    `json:"building_id"`
	Name       string    `json:"name"`
	Level      int       `json:"level"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Rooms      []Room    `gorm:"foreignKey:FloorID" json:"-"`
}

// Room represents a room (data center room, server room, etc.)
type Room struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	FloorID   string    `json:"floor_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Area      float64   `json:"area"` // in square meters
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Racks     []Rack    `gorm:"foreignKey:RoomID" json:"-"`
}

// Rack represents a server rack
type Rack struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	RoomID    string    `json:"room_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Height    int       `json:"height"` // in U units
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Devices   []Device  `gorm:"foreignKey:RackID" json:"-"`
}

// Device represents a physical device (server, switch, etc.)
type Device struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	RackID       string    `json:"rack_id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"` // server, switch, router, storage, etc.
	Model        string    `json:"model"`
	SerialNumber string    `json:"serial_number"`
	Position     int       `json:"position"` // U position in rack
	Height       int       `json:"height"` // in U units
	Power        float64   `json:"power"` // in watts
	Status       string    `json:"status"` // online, offline, maintenance
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// BeforeCreate hook to generate UUIDs
func (t *Tenant) BeforeCreate(tx *gorm.DB) error {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

func (a *AuditLog) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return nil
}

func (s *Subscription) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

func (s *Site) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

func (b *Building) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return nil
}

func (f *Floor) BeforeCreate(tx *gorm.DB) error {
	if f.ID == "" {
		f.ID = uuid.New().String()
	}
	return nil
}

func (r *Room) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

func (r *Rack) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

func (d *Device) BeforeCreate(tx *gorm.DB) error {
	if d.ID == "" {
		d.ID = uuid.New().String()
	}
	return nil
}
