package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/models"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReservationCreate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userClaim := c.MustGet("userClaims").(*utils.UserClaims)

	// Get user details
	var user models.User

	if err := db.Where("id = ?", userClaim.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
			"error_code": http.StatusNotFound,
		})
		return
	}

	var reservationBody struct {
		Service     string `json:"service"`
		Date        string `json:"date"`
		Time		string `json:"time"`
		BranchID    uint   `json:"branch_id"`
	}

	// Bind request body to struct
	if err := c.ShouldBindJSON(&reservationBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var branch models.Branch

	if err := db.Where("id = ?", reservationBody.BranchID).First(&branch).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Branch not found",
			"error_code": http.StatusNotFound,
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
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,	
		Service:     reservationBody.Service,
		Date:        date,
		UserID:      user.ID,
		BranchID:    reservationBody.BranchID,
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

func ReservationGetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userClaim := c.MustGet("userClaims").(*utils.UserClaims)

	var reservations []models.Reservation

	if err := db.Where("user_id = ?", userClaim.ID).Find(&reservations).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	c.JSON(http.StatusOK, reservations)
}

func ReservationDelete(c * gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userClaim := c.MustGet("userClaims").(*utils.UserClaims)
	reservationId := c.Param("id")

	var reservation models.Reservation
	if err := db.Where("id = ?", reservationId).First(&reservation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Reservation not found",
			})
			return
		}
	}

	if reservation.UserID != userClaim.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	if err := db.Delete(&reservation).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusOK)
}