package services

import (
	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"gorm.io/gorm"
)

type DCIMService struct {
	db *gorm.DB
}

func NewDCIMService(db *gorm.DB) *DCIMService {
	return &DCIMService{db: db}
}

// Site operations
func (s *DCIMService) CreateSite(tenantID string, site *models.Site) error {
	site.TenantID = tenantID
	return s.db.Create(site).Error
}

func (s *DCIMService) GetSites(tenantID string) ([]models.Site, error) {
	var sites []models.Site
	err := s.db.Where("tenant_id = ?", tenantID).Find(&sites).Error
	return sites, err
}

func (s *DCIMService) GetSite(tenantID, siteID string) (*models.Site, error) {
	var site models.Site
	err := s.db.Where("id = ? AND tenant_id = ?", siteID, tenantID).First(&site).Error
	return &site, err
}

func (s *DCIMService) UpdateSite(tenantID, siteID string, updates map[string]interface{}) error {
	return s.db.Model(&models.Site{}).
		Where("id = ? AND tenant_id = ?", siteID, tenantID).
		Updates(updates).Error
}

func (s *DCIMService) DeleteSite(tenantID, siteID string) error {
	return s.db.Where("id = ? AND tenant_id = ?", siteID, tenantID).Delete(&models.Site{}).Error
}

// Building operations
func (s *DCIMService) CreateBuilding(building *models.Building) error {
	return s.db.Create(building).Error
}

func (s *DCIMService) GetBuildings(siteID string) ([]models.Building, error) {
	var buildings []models.Building
	err := s.db.Where("site_id = ?", siteID).Find(&buildings).Error
	return buildings, err
}

func (s *DCIMService) GetBuilding(buildingID string) (*models.Building, error) {
	var building models.Building
	err := s.db.First(&building, "id = ?", buildingID).Error
	return &building, err
}

func (s *DCIMService) UpdateBuilding(buildingID string, updates map[string]interface{}) error {
	return s.db.Model(&models.Building{}).Where("id = ?", buildingID).Updates(updates).Error
}

func (s *DCIMService) DeleteBuilding(buildingID string) error {
	return s.db.Delete(&models.Building{}, "id = ?", buildingID).Error
}

// Floor operations
func (s *DCIMService) CreateFloor(floor *models.Floor) error {
	return s.db.Create(floor).Error
}

func (s *DCIMService) GetFloors(buildingID string) ([]models.Floor, error) {
	var floors []models.Floor
	err := s.db.Where("building_id = ?", buildingID).Find(&floors).Error
	return floors, err
}

func (s *DCIMService) GetFloor(floorID string) (*models.Floor, error) {
	var floor models.Floor
	err := s.db.First(&floor, "id = ?", floorID).Error
	return &floor, err
}

func (s *DCIMService) UpdateFloor(floorID string, updates map[string]interface{}) error {
	return s.db.Model(&models.Floor{}).Where("id = ?", floorID).Updates(updates).Error
}

func (s *DCIMService) DeleteFloor(floorID string) error {
	return s.db.Delete(&models.Floor{}, "id = ?", floorID).Error
}

// Room operations
func (s *DCIMService) CreateRoom(room *models.Room) error {
	return s.db.Create(room).Error
}

func (s *DCIMService) GetRooms(floorID string) ([]models.Room, error) {
	var rooms []models.Room
	err := s.db.Where("floor_id = ?", floorID).Find(&rooms).Error
	return rooms, err
}

func (s *DCIMService) GetRoom(roomID string) (*models.Room, error) {
	var room models.Room
	err := s.db.First(&room, "id = ?", roomID).Error
	return &room, err
}

func (s *DCIMService) UpdateRoom(roomID string, updates map[string]interface{}) error {
	return s.db.Model(&models.Room{}).Where("id = ?", roomID).Updates(updates).Error
}

func (s *DCIMService) DeleteRoom(roomID string) error {
	return s.db.Delete(&models.Room{}, "id = ?", roomID).Error
}

// Rack operations
func (s *DCIMService) CreateRack(rack *models.Rack) error {
	return s.db.Create(rack).Error
}

func (s *DCIMService) GetRacks(roomID string) ([]models.Rack, error) {
	var racks []models.Rack
	err := s.db.Where("room_id = ?", roomID).Find(&racks).Error
	return racks, err
}

func (s *DCIMService) GetRack(rackID string) (*models.Rack, error) {
	var rack models.Rack
	err := s.db.First(&rack, "id = ?", rackID).Error
	return &rack, err
}

func (s *DCIMService) UpdateRack(rackID string, updates map[string]interface{}) error {
	return s.db.Model(&models.Rack{}).Where("id = ?", rackID).Updates(updates).Error
}

func (s *DCIMService) DeleteRack(rackID string) error {
	return s.db.Delete(&models.Rack{}, "id = ?", rackID).Error
}

// Device operations
func (s *DCIMService) CreateDevice(device *models.Device) error {
	return s.db.Create(device).Error
}

func (s *DCIMService) GetDevices(rackID string) ([]models.Device, error) {
	var devices []models.Device
	err := s.db.Where("rack_id = ?", rackID).Find(&devices).Error
	return devices, err
}

func (s *DCIMService) GetDevice(deviceID string) (*models.Device, error) {
	var device models.Device
	err := s.db.First(&device, "id = ?", deviceID).Error
	return &device, err
}

func (s *DCIMService) UpdateDevice(deviceID string, updates map[string]interface{}) error {
	return s.db.Model(&models.Device{}).Where("id = ?", deviceID).Updates(updates).Error
}

func (s *DCIMService) DeleteDevice(deviceID string) error {
	return s.db.Delete(&models.Device{}, "id = ?", deviceID).Error
}

// Dashboard statistics
type DashboardStats struct {
	TotalSites      int64   `json:"total_sites"`
	TotalBuildings  int64   `json:"total_buildings"`
	TotalRooms      int64   `json:"total_rooms"`
	TotalRacks      int64   `json:"total_racks"`
	TotalDevices    int64   `json:"total_devices"`
	OnlineDevices   int64   `json:"online_devices"`
	OfflineDevices  int64   `json:"offline_devices"`
	TotalPower      float64 `json:"total_power"`
}

func (s *DCIMService) GetDashboardStats(tenantID string) (*DashboardStats, error) {
	stats := &DashboardStats{}

	s.db.Model(&models.Site{}).Where("tenant_id = ?", tenantID).Count(&stats.TotalSites)
	s.db.Model(&models.Building{}).Count(&stats.TotalBuildings)
	s.db.Model(&models.Room{}).Count(&stats.TotalRooms)
	s.db.Model(&models.Rack{}).Count(&stats.TotalRacks)
	s.db.Model(&models.Device{}).Count(&stats.TotalDevices)
	s.db.Model(&models.Device{}).Where("status = ?", "online").Count(&stats.OnlineDevices)
	s.db.Model(&models.Device{}).Where("status = ?", "offline").Count(&stats.OfflineDevices)

	var totalPower float64
	s.db.Model(&models.Device{}).Select("COALESCE(SUM(power), 0)").Row().Scan(&totalPower)
	stats.TotalPower = totalPower

	return stats, nil
}
