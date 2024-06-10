package main

import (
	routes "go-gin-jwt-auth/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api version 1"})
	})

	router.GET("/api/v2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api version 2"})
	})

	router.Run(":" + port)
}
