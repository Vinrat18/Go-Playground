package main

import (
	"log"
	"time"
	"web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client IP address
		clientIP := c.ClientIP()
		// Get the current time
		now := time.Now()
		// Log the request
		log.Printf("[%s] %s %s %s", now.Format(time.RFC3339), c.Request.Method, c.Request.URL.Path, clientIP)
		// Proceed to the next handler
		c.Next()
	}
}

func main() {
	router := gin.New()

	router.Use(Logger())

	routes.TsRoutes(router.Group("/ts"))
	routes.AssetRoutes(router.Group("/asset"))

	router.Run(":8081")
}
