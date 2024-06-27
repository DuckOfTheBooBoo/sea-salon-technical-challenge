package controllers

import (
	"log"
	"net/http"
	"time"

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
		BranchName    string   `json:"branch_name"`
		BranchAddress string   `json:"branch_address"`
		Lat           float64  `json:"lat"`
		Lng           float64  `json:"lng"`
		CloseTime     string   `json:"close_time"`
		OpenTime      string   `json:"open_time"`
		Services      []string `json:"services"`
	}

	// Bind request body to struct
	if err := c.ShouldBindJSON(&branchBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	layout := "15:04"

	// Convert open time to time.Time
	_, err := time.Parse(layout, branchBody.OpenTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong time format. Should be hh:mm",
			"time":  branchBody.OpenTime,
		})

		return
	}

	// Convert close time to time.Time
	_, err = time.Parse(layout, branchBody.CloseTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong time format. Should be hh:mm",
			"time":  branchBody.CloseTime,
		})

		return
	}

	var services []models.Service

	// Get services from service code
	for _, code := range branchBody.Services {
		var service models.Service
		result := db.Where("service_code = ?", code).First(&service)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": code + " service not found",
				})
				return
			}
		}
		services = append(services, service)
	}

	// Create new branch
	newBranch := models.Branch{
		BranchName:    branchBody.BranchName,
		BranchAddress: branchBody.BranchAddress,
		Lat:           branchBody.Lat,
		Lng:           branchBody.Lng,
		OpenAt:        branchBody.OpenTime,
		ClosedAt:      branchBody.CloseTime,
	}

	// Save new branch to database
	if err := db.Create(&newBranch).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Add services to branch
	if err := db.Model(&newBranch).Association("Services").Append(&services); err != nil {
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

	if err := db.Preload("Services").Find(&branches).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, branches)
}

func BranchDelete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	branchId := c.Param("id")

	var branch models.Branch
	if err := db.Preload("Services").Find(&branch).Where("id = ?", branchId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := db.Model(&branch).Association("Services").Clear(); err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := db.Delete(&branch).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	c.Status(http.StatusOK)
}
