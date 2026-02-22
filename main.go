package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Containerized microservice with Kubernetes support
func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "neuro-flow-worker-iot-nl2",
			"status": "running",
		})
	})
	
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	
	r.Run(":8080")
}
