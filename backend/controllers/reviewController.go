package controllers

import (
	"log"
	"net/http"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ReviewGet(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var reviews []models.Review

	if err := db.Find(&reviews); err.Error != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error)
		return
	}

	c.JSON(http.StatusOK, reviews)
}


func ReviewCreate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	validate := validator.New()

	var reviewBody struct {
		FullName string `json:"full_name" validate:"required"`
		Rating   uint   `json:"rating" validate:"required,gt=0,lte=5"`
		Comement string `json:"comment" validate:"required"`
	}

	// Bind request body to struct
	if err := c.ShouldBindJSON(&reviewBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validate.Struct(reviewBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create review
	review := models.Review{
		FullName: reviewBody.FullName,
		Rating:   reviewBody.Rating,
		Comment:  reviewBody.Comement,
	}

	if err := db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"review": review,
	})
}
