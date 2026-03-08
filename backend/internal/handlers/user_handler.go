package handlers

import (
	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)

	var users []models.User
	if err := h.db.Where("tenant_id = ?", tenantID).Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch users",
		})
	}

	return c.JSON(fiber.Map{
		"users": users,
		"count": len(users),
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)
	userID := c.Params("id")

	var user models.User
	if err := h.db.Where("id = ? AND tenant_id = ?", userID, tenantID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user",
		})
	}

	return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)
	userID := c.Params("id")

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.db.Model(&models.User{}).
		Where("id = ? AND tenant_id = ?", userID, tenantID).
		Updates(updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "user updated successfully",
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	tenantID := c.Locals("tenant_id").(string)
	userID := c.Params("id")

	if err := h.db.Where("id = ? AND tenant_id = ?", userID, tenantID).Delete(&models.User{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "user deleted successfully",
	})
}
