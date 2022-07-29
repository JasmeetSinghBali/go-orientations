package main

import (
	routes "go_rest_jwt_gin_mongodb/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port = os.Getenv("PORT")

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
}
