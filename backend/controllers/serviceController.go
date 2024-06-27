package controllers

import (
	"log"
	"net/http"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ServiceCreate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var serviceBody struct {
		ServiceName string `json:"service_name"`
		ServiceCode string `json:"service_code"`
		Duration    uint   `json:"duration"`
	}

	// Bind request body to struct
	if err := c.ShouldBindJSON(&serviceBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create new service
	newService := models.Service{
		ServiceName: serviceBody.ServiceName,
		ServiceCode: serviceBody.ServiceCode,
		Duration:    serviceBody.Duration,
	}

	// Save service to database
	if err := db.Create(&newService).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusCreated, newService)
}

func ServiceGetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var services []models.Service

	if err := db.Find(&services).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, services)
}
