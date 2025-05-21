package main

import (
	"auth/handlers"
	"auth/middleware"
	"auth/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	// Public routes
	r.POST("/register", handlers.Register(db))
	r.POST("/login", handlers.Login(db))

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/user", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "User route", "username": c.GetString("username")})
		})

		// Admin-only route
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleMiddleware("admin"))
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Admin dashboard"})
			})
		}
	}

	r.Run(":8080")
}
