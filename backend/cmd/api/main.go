package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ariveratij40-lab/skia/backend/internal/config"
	"github.com/ariveratij40-lab/skia/backend/internal/handlers"
	"github.com/ariveratij40-lab/skia/backend/internal/middleware"
	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Auto migrate models
	if err := db.AutoMigrate(
		&models.User{},
		&models.Tenant{},
		&models.Role{},
		&models.Permission{},
		&models.AuditLog{},
		&models.Site{},
		&models.Building{},
		&models.Floor{},
		&models.Room{},
		&models.Rack{},
		&models.Device{},
		&models.Subscription{},
	); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "SKIA API v1.0.0",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS_ORIGINS"),
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"version": "1.0.0",
			"phase":   "0-1",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// Auth routes (public)
	authHandler := handlers.NewAuthHandler(db)
	api.Post("/auth/register", authHandler.Register)
	api.Post("/auth/login", authHandler.Login)
	api.Post("/auth/refresh", authHandler.RefreshToken)

	// Protected routes
	api.Use(middleware.AuthMiddleware())

	// User routes
	userHandler := handlers.NewUserHandler(db)
	api.Get("/users", userHandler.GetUsers)
	api.Get("/users/:id", userHandler.GetUser)
	api.Put("/users/:id", userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.DeleteUser)

	// DCIM routes
	dcimHandler := handlers.NewDCIMHandler(db)

	// Sites
	api.Get("/sites", dcimHandler.GetSites)
	api.Post("/sites", dcimHandler.CreateSite)
	api.Get("/sites/:id", dcimHandler.GetSite)
	api.Put("/sites/:id", dcimHandler.UpdateSite)
	api.Delete("/sites/:id", dcimHandler.DeleteSite)

	// Buildings
	api.Get("/buildings", dcimHandler.GetBuildings)
	api.Post("/buildings", dcimHandler.CreateBuilding)
	api.Get("/buildings/:id", dcimHandler.GetBuilding)
	api.Put("/buildings/:id", dcimHandler.UpdateBuilding)
	api.Delete("/buildings/:id", dcimHandler.DeleteBuilding)

	// Floors
	api.Get("/floors", dcimHandler.GetFloors)
	api.Post("/floors", dcimHandler.CreateFloor)
	api.Get("/floors/:id", dcimHandler.GetFloor)
	api.Put("/floors/:id", dcimHandler.UpdateFloor)
	api.Delete("/floors/:id", dcimHandler.DeleteFloor)

	// Rooms
	api.Get("/rooms", dcimHandler.GetRooms)
	api.Post("/rooms", dcimHandler.CreateRoom)
	api.Get("/rooms/:id", dcimHandler.GetRoom)
	api.Put("/rooms/:id", dcimHandler.UpdateRoom)
	api.Delete("/rooms/:id", dcimHandler.DeleteRoom)

	// Racks
	api.Get("/racks", dcimHandler.GetRacks)
	api.Post("/racks", dcimHandler.CreateRack)
	api.Get("/racks/:id", dcimHandler.GetRack)
	api.Put("/racks/:id", dcimHandler.UpdateRack)
	api.Delete("/racks/:id", dcimHandler.DeleteRack)

	// Devices
	api.Get("/devices", dcimHandler.GetDevices)
	api.Post("/devices", dcimHandler.CreateDevice)
	api.Get("/devices/:id", dcimHandler.GetDevice)
	api.Put("/devices/:id", dcimHandler.UpdateDevice)
	api.Delete("/devices/:id", dcimHandler.DeleteDevice)

	// Dashboard
	api.Get("/dashboard", dcimHandler.GetDashboard)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 SKIA API Server starting on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
