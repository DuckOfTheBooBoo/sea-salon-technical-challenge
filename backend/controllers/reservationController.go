package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReservationCreate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var reservationBody struct {
		FullName    string `json:"full_name"`
		PhoneNumber string `json:"phone_number"`
		Service     string `json:"service"`
		Date        string `json:"date"`
		Time		string `json:"time"`
	}

	// Bind request body to struct
	if err := c.ShouldBindJSON(&reservationBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Parse date
	date, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", reservationBody.Date, reservationBody.Time))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create reservation
	reservation := models.Reservation{
		FullName:    reservationBody.FullName,
		PhoneNumber: reservationBody.PhoneNumber,
		Service:     reservationBody.Service,
		Date:        date,
	}

	// Save reservation to database
	if err := db.Create(&reservation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"reservation": reservation,
	})
}
