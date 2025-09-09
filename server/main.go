package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"powerapp/server/database"
	"powerapp/server/handlers"
	"powerapp/server/middleware"
	"powerapp/server/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	database.ConnectDatabase()

	// Auto-migrate the models
	if err := database.DB.AutoMigrate(&models.User{}, &models.AuthToken{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Create default admin user if it doesn't exist
	createDefaultAdmin()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "CDK Engine",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	app.Use(logger.New())
	app.Use(recover.New())

	// Serve static files
	app.Static("/uploads", "./uploads")

	// Auth routes
	app.Post("/api/auth/login", handlers.Login)
	app.Post("/api/auth/register", handlers.Register)

	// Protected routes
	api := app.Group("/api", middleware.AuthMiddleware())
	api.Post("/auth/logout", handlers.Logout)
	api.Get("/auth/me", handlers.GetCurrentUser)
	api.Get("/users", handlers.GetUsers)
	api.Get("/users/:id", handlers.GetUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)
	
	// Upload routes
	api.Post("/upload/avatar", handlers.UploadAvatar)
	api.Delete("/upload/avatar", handlers.DeleteAvatar)

	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! This is your Go Fiber template.")
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func createDefaultAdmin() {
	var user models.User
	if err := database.DB.Where("username = ?", "admin").First(&user).Error; err != nil {
		// User doesn't exist, create it
		adminUser := models.User{
			Username: "admin",
			Password: "admin",
		}

		if err := adminUser.HashPassword(); err != nil {
			log.Printf("Failed to hash admin password: %v", err)
			return
		}

		if err := database.DB.Create(&adminUser).Error; err != nil {
			log.Printf("Failed to create admin user: %v", err)
			return
		}

		log.Println("Default admin user created (username: admin, password: admin)")
	} else {
		log.Println("Admin user already exists")
	}
}
