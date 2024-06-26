package controllers

import (
	"log"
	"net/http"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/models"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BranchCreate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userClaim := c.MustGet("userClaims").(*utils.UserClaims)

	// Get user details
	var user models.User

	if err := db.Where("id = ?", userClaim.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":      "User not found",
			"error_code": http.StatusNotFound,
		})
		return
	}

	// if user.Role != "admin" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error":      "Unauthorized",
	// 		"error_code": http.StatusUnauthorized,
	// 	})
	// 	return
	// }

	var branchBody struct {
		BranchName    string `json:"branch_name"`
		BranchAddress string `json:"branch_address"`
		Lat       float64 `json:"lat"`
		Lng          float64 `json:"lng"`
	}

	// Bind request body to struct
	if err := c.ShouldBindJSON(&branchBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create new branch
	newBranch := models.Branch{
		BranchName:    branchBody.BranchName,
		BranchAddress: branchBody.BranchAddress,
		Lat:     branchBody.Lat,
		Lng:    branchBody.Lng,
	}

	// Save new branch to database
	if err := db.Create(&newBranch).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"branch": newBranch,
	})
}

func BranchGetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var branches []models.Branch

	if err := db.Find(&branches).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, branches)
}