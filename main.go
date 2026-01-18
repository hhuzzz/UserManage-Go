package main

import (
	"hello/auth"
	"hello/config"
	"hello/controllers"
	"hello/database"
	"hello/repositories"
	"hello/routes"
	"hello/services"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	err = database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize layers
	userRepo := repositories.NewUserRepository(database.GetDB())
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Initialize JWT manager
	jwtManager := auth.NewJWTManager(cfg.JWTSecret, time.Duration(cfg.JWTExpiration)*time.Second)

	// Initialize auth service and controller
	authService := auth.NewAuthService(userRepo, jwtManager)
	authController := controllers.NewAuthController(authService)

	// Setup Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Static files
	r.Static("/static", "./static")

	// Setup routes
	routes.SetupRoutes(r, userController, authController, jwtManager)

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("Server starting on http://localhost%s", addr)
	log.Fatal(r.Run(addr))
}
