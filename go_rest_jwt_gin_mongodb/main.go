package main

import (
	"log"
	"os"

	routes "github.com/Jasmeet-1998/go-orientations/go_rest_jwt_gin_mongodb/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// gin context to access response & request
	router.GET("/api_v0.0.1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Welcome to api_v0.0.1!"})
	})

	router.GET("/api_v0.0.2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Welcome to api_v0.0.2!"})
	})

	router.Run(":" + port)
}
