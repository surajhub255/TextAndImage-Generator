package main

import (
	"back/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/generate", services.GenerateImage)
	router.POST("/generate-text", services.HandleGenerateText)

	port := 8080
	fmt.Printf("Server running on :%d...\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}
