package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/amliyanage/go-jwt-tasks/config"
	"github.com/amliyanage/go-jwt-tasks/controllers"
	"github.com/amliyanage/go-jwt-tasks/middleware"
	"github.com/amliyanage/go-jwt-tasks/repo"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	repo.Init()

	// Setup router
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "server is running",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Public routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controllers.Register(cfg))
			auth.POST("/login", controllers.Login(cfg))
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(cfg))
		{
			// User profile
			protected.GET("/profile", controllers.GetProfile())

			// Task routes
			tasks := protected.Group("/tasks")
			{
				tasks.POST("", controllers.CreateTask())
				tasks.GET("", controllers.GetTasks())
				tasks.GET("/:id", controllers.GetTask())
				tasks.PUT("/:id", controllers.UpdateTask())
				tasks.DELETE("/:id", controllers.DeleteTask())
			}
		}
	}

	// Start server
	log.Printf("Server starting on port %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
