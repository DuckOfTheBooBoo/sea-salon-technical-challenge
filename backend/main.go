package main

import (
	"log"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/configs"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnv()
}

func main() {
	r := gin.Default()

	// Set up database connection
	_, err := configs.ConnectToDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	// Define root endpoint
	api := r.Group("/api")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run(":3000")
}
