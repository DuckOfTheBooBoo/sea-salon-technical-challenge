package main

import (
	"log"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/configs"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/models"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/utils"
)

func init() {
	utils.LoadEnv()
}

func main() {

	db, err := configs.ConnectToDB()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	gormDB := db.GetDB()

	if err := gormDB.AutoMigrate(&models.Reservation{}); err != nil {
		log.Fatal(err.Error())
		return
	}


}
