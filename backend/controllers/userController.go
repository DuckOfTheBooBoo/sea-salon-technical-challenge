package controllers

import (
	"log"
	"net/http"
	"fmt"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/models"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// User Sign Up
func UserCreate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	validate := validator.New()

	var userBody struct {
		FirstName string `json:"first_name" validate:"required"` 
		LastName string `json:"last_name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		PhoneNumber string `json:"phone_number" validate:"required,e164"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.ShouldBindJSON(&userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validate.Struct(userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Hash password
	hashedPass, err := utils.HashPassword(userBody.Password)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	// Create user
	user := models.User{
		FullName:  fmt.Sprintf("%s %s", userBody.FirstName, userBody.LastName),
		Email:     userBody.Email,
		PhoneNumber: userBody.PhoneNumber,
		Password:  hashedPass,
		Role:      "Customer",
	}

	if err := db.Create(&user).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}
