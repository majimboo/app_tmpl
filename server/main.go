package main

import (
	"log"
	"os"

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

	// Auto-migrate the User model
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	app.Use(logger.New())
	app.Use(recover.New())

	// Auth routes
	app.Post("/api/auth/login", handlers.Login)
	app.Post("/api/auth/register", handlers.Register)

	// Protected routes
	api := app.Group("/api", middleware.JWTMiddleware())
	api.Get("/users", handlers.GetUsers)
	api.Get("/users/:id", handlers.GetUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)

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
