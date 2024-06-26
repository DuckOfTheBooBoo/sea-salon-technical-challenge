package main

import (
	"log"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/configs"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/models"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/utils"
	"gorm.io/gorm"
)

func init() {
	utils.LoadEnv()
}

func createAdminUser(gormDB *gorm.DB) {
	var user models.User
	if err := gormDB.Where("email = ?", "thomas.n@compfest.id").First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			hashedPass, err := utils.HashPassword("Admin123")

			if err != nil {
				log.Fatal(err.Error())
				return
			}

			user = models.User{
				FullName: "Thomas N",
				Email:    "thomas.n@compfest.id",
				Password: hashedPass,
				Role:     "Admin",
			}

			if err := gormDB.Create(&user); err != nil {
				log.Fatal(err)
			}

			log.Print("Admin user created successfully")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Print("Admin user already exists")
	}
}

func main() {

	db, err := configs.ConnectToDB()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	gormDB := db.GetDB()

	if err := gormDB.AutoMigrate(&models.User{}, &models.Reservation{}, &models.Review{}, &models.Token{}, &models.Branch{}); err != nil {
		log.Fatal(err.Error())
		return
	}

	createAdminUser(gormDB)
}
