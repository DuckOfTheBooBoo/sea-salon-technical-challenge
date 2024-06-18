package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/utils"
)

func init() {
	utils.LoadEnv()
}

func main() {
	r := gin.Default()

	// Define root endpoint
	api := r.Group("/api")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run(":3000")
}