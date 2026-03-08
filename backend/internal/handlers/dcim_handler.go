package handlers

import (
	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/ariveratij40-lab/skia/backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DCIMHandler struct {
	dcimService *services.DCIMService
}

func NewDCIMHandler(db *gorm.DB) *DCIMHandler {
	return &DCIMHandler{
		dcimService: services.NewDCIMService(db),
	}
}

// Site handlers
func (h *DCIMHandler) GetSites(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)

	sites, err := h.dcimService.GetSites(tenantID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch sites",
		})
	}

	return c.JSON(fiber.Map{
		"sites": sites,
		"count": len(sites),
	})
}

func (h *DCIMHandler) CreateSite(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)

	var site models.Site
	if err := c.BodyParser(&site); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.CreateSite(tenantID, &site); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create site",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(site)
}

func (h *DCIMHandler) GetSite(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)
	siteID := c.Params("id")

	site, err := h.dcimService.GetSite(tenantID, siteID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "site not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch site",
		})
	}

	return c.JSON(site)
}

func (h *DCIMHandler) UpdateSite(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)
	siteID := c.Params("id")

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.UpdateSite(tenantID, siteID, updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update site",
		})
	}

	return c.JSON(fiber.Map{
		"message": "site updated successfully",
	})
}

func (h *DCIMHandler) DeleteSite(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)
	siteID := c.Params("id")

	if err := h.dcimService.DeleteSite(tenantID, siteID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete site",
		})
	}

	return c.JSON(fiber.Map{
		"message": "site deleted successfully",
	})
}

// Building handlers
func (h *DCIMHandler) GetBuildings(c *fiber.Ctx) error {
	siteID := c.Query("site_id")

	buildings, err := h.dcimService.GetBuildings(siteID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch buildings",
		})
	}

	return c.JSON(fiber.Map{
		"buildings": buildings,
		"count":     len(buildings),
	})
}

func (h *DCIMHandler) CreateBuilding(c *fiber.Ctx) error {
	var building models.Building
	if err := c.BodyParser(&building); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.CreateBuilding(&building); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create building",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(building)
}

func (h *DCIMHandler) GetBuilding(c *fiber.Ctx) error {
	buildingID := c.Params("id")

	building, err := h.dcimService.GetBuilding(buildingID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "building not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch building",
		})
	}

	return c.JSON(building)
}

func (h *DCIMHandler) UpdateBuilding(c *fiber.Ctx) error {
	buildingID := c.Params("id")

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.UpdateBuilding(buildingID, updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update building",
		})
	}

	return c.JSON(fiber.Map{
		"message": "building updated successfully",
	})
}

func (h *DCIMHandler) DeleteBuilding(c *fiber.Ctx) error {
	buildingID := c.Params("id")

	if err := h.dcimService.DeleteBuilding(buildingID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete building",
		})
	}

	return c.JSON(fiber.Map{
		"message": "building deleted successfully",
	})
}

// Floor handlers
func (h *DCIMHandler) GetFloors(c *fiber.Ctx) error {
	buildingID := c.Query("building_id")

	floors, err := h.dcimService.GetFloors(buildingID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch floors",
		})
	}

	return c.JSON(fiber.Map{
		"floors": floors,
		"count":  len(floors),
	})
}

func (h *DCIMHandler) CreateFloor(c *fiber.Ctx) error {
	var floor models.Floor
	if err := c.BodyParser(&floor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.CreateFloor(&floor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create floor",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(floor)
}

func (h *DCIMHandler) GetFloor(c *fiber.Ctx) error {
	floorID := c.Params("id")

	floor, err := h.dcimService.GetFloor(floorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "floor not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch floor",
		})
	}

	return c.JSON(floor)
}

func (h *DCIMHandler) UpdateFloor(c *fiber.Ctx) error {
	floorID := c.Params("id")

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.UpdateFloor(floorID, updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update floor",
		})
	}

	return c.JSON(fiber.Map{
		"message": "floor updated successfully",
	})
}

func (h *DCIMHandler) DeleteFloor(c *fiber.Ctx) error {
	floorID := c.Params("id")

	if err := h.dcimService.DeleteFloor(floorID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete floor",
		})
	}

	return c.JSON(fiber.Map{
		"message": "floor deleted successfully",
	})
}

// Room handlers
func (h *DCIMHandler) GetRooms(c *fiber.Ctx) error {
	floorID := c.Query("floor_id")

	rooms, err := h.dcimService.GetRooms(floorID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch rooms",
		})
	}

	return c.JSON(fiber.Map{
		"rooms": rooms,
		"count": len(rooms),
	})
}

func (h *DCIMHandler) CreateRoom(c *fiber.Ctx) error {
	var room models.Room
	if err := c.BodyParser(&room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.CreateRoom(&room); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create room",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(room)
}

func (h *DCIMHandler) GetRoom(c *fiber.Ctx) error {
	roomID := c.Params("id")

	room, err := h.dcimService.GetRoom(roomID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "room not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch room",
		})
	}

	return c.JSON(room)
}

func (h *DCIMHandler) UpdateRoom(c *fiber.Ctx) error {
	roomID := c.Params("id")

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.UpdateRoom(roomID, updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update room",
		})
	}

	return c.JSON(fiber.Map{
		"message": "room updated successfully",
	})
}

func (h *DCIMHandler) DeleteRoom(c *fiber.Ctx) error {
	roomID := c.Params("id")

	if err := h.dcimService.DeleteRoom(roomID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete room",
		})
	}

	return c.JSON(fiber.Map{
		"message": "room deleted successfully",
	})
}

// Rack handlers
func (h *DCIMHandler) GetRacks(c *fiber.Ctx) error {
	roomID := c.Query("room_id")

	racks, err := h.dcimService.GetRacks(roomID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch racks",
		})
	}

	return c.JSON(fiber.Map{
		"racks": racks,
		"count": len(racks),
	})
}

func (h *DCIMHandler) CreateRack(c *fiber.Ctx) error {
	var rack models.Rack
	if err := c.BodyParser(&rack); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.CreateRack(&rack); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create rack",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(rack)
}

func (h *DCIMHandler) GetRack(c *fiber.Ctx) error {
	rackID := c.Params("id")

	rack, err := h.dcimService.GetRack(rackID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "rack not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch rack",
		})
	}

	return c.JSON(rack)
}

func (h *DCIMHandler) UpdateRack(c *fiber.Ctx) error {
	rackID := c.Params("id")

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.UpdateRack(rackID, updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update rack",
		})
	}

	return c.JSON(fiber.Map{
		"message": "rack updated successfully",
	})
}

func (h *DCIMHandler) DeleteRack(c *fiber.Ctx) error {
	rackID := c.Params("id")

	if err := h.dcimService.DeleteRack(rackID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete rack",
		})
	}

	return c.JSON(fiber.Map{
		"message": "rack deleted successfully",
	})
}

// Device handlers
func (h *DCIMHandler) GetDevices(c *fiber.Ctx) error {
	rackID := c.Query("rack_id")

	devices, err := h.dcimService.GetDevices(rackID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch devices",
		})
	}

	return c.JSON(fiber.Map{
		"devices": devices,
		"count":   len(devices),
	})
}

func (h *DCIMHandler) CreateDevice(c *fiber.Ctx) error {
	var device models.Device
	if err := c.BodyParser(&device); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.CreateDevice(&device); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create device",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(device)
}

func (h *DCIMHandler) GetDevice(c *fiber.Ctx) error {
	deviceID := c.Params("id")

	device, err := h.dcimService.GetDevice(deviceID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "device not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch device",
		})
	}

	return c.JSON(device)
}

func (h *DCIMHandler) UpdateDevice(c *fiber.Ctx) error {
	deviceID := c.Params("id")

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.dcimService.UpdateDevice(deviceID, updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update device",
		})
	}

	return c.JSON(fiber.Map{
		"message": "device updated successfully",
	})
}

func (h *DCIMHandler) DeleteDevice(c *fiber.Ctx) error {
	deviceID := c.Params("id")

	if err := h.dcimService.DeleteDevice(deviceID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete device",
		})
	}

	return c.JSON(fiber.Map{
		"message": "device deleted successfully",
	})
}

// Dashboard handler
func (h *DCIMHandler) GetDashboard(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)

	stats, err := h.dcimService.GetDashboardStats(tenantID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch dashboard stats",
		})
	}

	return c.JSON(stats)
}
