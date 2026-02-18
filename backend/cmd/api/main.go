package main

import (
	"log"
	"os"

	"github.com/ariveratij40-lab/skia/backend/internal/config"
	"github.com/ariveratij40-lab/skia/backend/internal/handlers"
	"github.com/ariveratij40-lab/skia/backend/internal/middleware"
	"github.com/ariveratij40-lab/skia/backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Configuración
	cfg := config.Load()

	// Conexión a base de datos
	db, err := repository.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Repositorios
	userRepo := repository.NewUserRepository(db)
	nodeRepo := repository.NewNodeRepository(db)
	scanRepo := repository.NewScanRepository(db)
	rackRepo := repository.NewRackRepository(db)

	// Handlers
	authHandler := handlers.NewAuthHandler(userRepo, cfg.JWTSecret)
	userHandler := handlers.NewUserHandler(userRepo)
	nodeHandler := handlers.NewNodeHandler(nodeRepo)
	scanHandler := handlers.NewScanHandler(scanRepo, nodeRepo)
	rackHandler := handlers.NewRackHandler(trackRepo)

	// Router
	r := gin.Default()

	// Middleware global
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"version": "1.0.0",
		})
	})

	// API v1
	v1 := r.Group("/v1")
	{
		// Auth (público)
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.Refresh)
		}

		// Protegido
		protected := v1.Group("")
		protected.Use(middleware.JWTAuth(cfg.JWTSecret))
		protected.Use(middleware.RLSContext())
		{
			auth.POST("/logout", authHandler.Logout)

			// Users
			users := protected.Group("/users")
			{
				users.GET("", userHandler.List)
				users.GET("/:id", userHandler.Get)
				users.POST("", middleware.RequireRole("admin"), userHandler.Create)
				users.PUT("/:id", userHandler.Update)
				users.DELETE("/:id", middleware.RequireRole("admin"), userHandler.Delete)
			}

			// Racks
			racks := protected.Group("/racks")
			{
				racks.GET("", rackHandler.List)
				racks.GET("/:id", rackHandler.Get)
				racks.POST("", middleware.RequireRole("admin"), rackHandler.Create)
				racks.PUT("/:id", middleware.RequireRole("admin"), rackHandler.Update)
				racks.DELETE("/:id", middleware.RequireRole("admin"), rackHandler.Delete)
			}

			// Nodes
			nodes := protected.Group("/nodes")
			{
				nodes.GET("", nodeHandler.List)
				nodes.GET("/:id", nodeHandler.Get)
				nodes.GET("/by-rfid/:rfid", nodeHandler.GetByRFID)
				nodes.POST("", middleware.RequireRole("admin"), nodeHandler.Create)
				nodes.PUT("/:id", middleware.RequireRole("admin"), nodeHandler.Update)
				nodes.DELETE("/:id", middleware.RequireRole("admin"), nodeHandler.Delete)
			}

			// Scans
			scans := protected.Group("/scans")
			{
				scans.GET("", scanHandler.List)
				scans.POST("", scanHandler.Create)
			}

			// Reports
			reports := protected.Group("/reports")
			{
				reports.GET("/dashboard", scanHandler.Dashboard)
				reports.GET("/activity", scanHandler.Activity)
			}
		}
	}

	// Iniciar servidor
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
