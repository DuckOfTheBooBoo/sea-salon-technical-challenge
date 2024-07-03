package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

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

	if err := gormDB.Create(&user).Error; err != nil {
		log.Println(err.Error())

		if err == gorm.ErrDuplicatedKey {
			log.Print("Admin user already exists")
		}
	}

	log.Print("Admin user created successfully")
}

func createServices(gormDB *gorm.DB) {
	var services []models.Service = []models.Service{
		{
			ServiceName: "Haircut",
			ServiceCode: "haircut",
			Duration:    30,
		},
		{
			ServiceName: "Manicures & Pedicures",
			ServiceCode: "manicures-pedicures",
			Duration:    40,
		},
		{
			ServiceName: "Facial Treatment",
			ServiceCode: "facial-treatment",
			Duration:    20,
		},
		{
			ServiceName: "Nail Art",
			ServiceCode: "nail-art",
			Duration:    40,
		},
		{
			ServiceName: "Scalp Treatments",
			ServiceCode: "scalp-treatments",
			Duration:    35,
		},
	}

	if err := gormDB.Create(&services).Error; err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Services created successfully")
}

func intTimeParser(s string) time.Time {
	t, err := utils.ParseTime(s)

	if err != nil {
		log.Fatalf("Error parsing time: %v", err)
	}

	return t
}

func createBranches(gormDB *gorm.DB) {
	var services []models.Service

	if err := gormDB.Find(&services).Error; err != nil {
		log.Fatal(err.Error())
	}

	var branches []models.Branch = []models.Branch{
		{
			BranchName:    "SEA Salon - Pondok Indah Mall II",
			BranchAddress: "Pondok Indah Mall 2, Level 1 Jl. Metro Pondok Indah Blok III-B Pondok Indah, Jakarta Selatan DKI Jakarta 12310 Indonesia",
			Lat:           -6.264692675855703,
			Lng:           106.78453850564463,
			OpenAt:        intTimeParser("09:00"),
			ClosedAt:      intTimeParser("20:00"),
		},
		{
			BranchName:    "SEA Salon - Cilandak Town Square",
			BranchAddress: "Jl. Cilandak No. 1, Cilandak, Jakarta Selatan 12430, Indonesia",
			Lat:           -6.291703706041596,
			Lng:           106.79988239264421,
			OpenAt:        intTimeParser("10:00"),
			ClosedAt:      intTimeParser("21:00"),
		},
		{
			BranchName:    "SEA Salon - Plaza Blok M",
			BranchAddress: "Plaza Bl. M, Blok M Plaza, Lantai 2, Jl. Bulungan No.76, RT.6/RW.6, Kramat Pela, Kebayoran Baru, South Jakarta City, Jakarta 12130",
			Lat:           -6.244134189991264,
			Lng:           106.7975735709911,
			OpenAt:        intTimeParser("10:00"),
			ClosedAt:      intTimeParser("22:00"),
		},
	}

	// Run the loop 3 times
	for i := 0; i < 3; i++ {
		numToPick := rand.Intn(2) + 2 // Generates either 2 or 3

		// Slice to store the picked services
		var pickedServices []models.Service

		// Pick the services
		for len(pickedServices) < numToPick {
			index := rand.Intn(len(services)) // Generate a random index
			pickedServices = append(pickedServices, services[index])
		}

		if err := gormDB.Create(&branches[i]).Error; err != nil {
			log.Fatal(err.Error())
		}

		log.Print("Branch created successfully")

		if err := gormDB.Model(&branches[i]).Association("Services").Append(&pickedServices); err != nil {
			log.Fatal(err.Error())
		}

		log.Print("Branch services created successfully")
	}
}

func createReviews(gormDB *gorm.DB) {
	reviews := []models.Review{
		{
			FullName: "John Doe",
			Rating:   5,
			Comment:  "Great haircut, very satisfied!",
		},
		{
			FullName: "Jane Smith",
			Rating:   4,
			Comment:  "Facial treatment was very relaxing.",
		},
		{
			FullName: "Emily Davis",
			Rating:   3,
			Comment:  "Love my new nails!",
		},
		{
			FullName: "Michael Johnson",
			Rating:   5,
			Comment:  "Manicure & pedicure were excellent.",
		},
		{
			FullName: "Sarah Brown",
			Rating:   4,
			Comment:  "Scalp treatment was very soothing.",
		},
		{
			FullName: "John Doe",
			Rating:   5,
			Comment:  "Excellent service and friendly staff.",
		},
		{
			FullName: "Jane Smith",
			Rating:   4,
			Comment:  "Skin feels so smooth and rejuvenated.",
		},
		{
			FullName: "Emily Davis",
			Rating:   3,
			Comment:  "Beautiful nail art, very creative.",
		},
		{
			FullName: "Michael Johnson",
			Rating:   5,
			Comment:  "My nails look fantastic.",
		},
		{
			FullName: "Sarah Brown",
			Rating:   4,
			Comment:  "Felt very relaxed after the treatment.",
		},
	}

	if err := gormDB.Create(&reviews).Error; err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Reviews created successfully")
}

func setMigrated() {
	data, err := os.ReadFile("status.json")

	if err != nil {
		panic(err)
	}

	var statusMap map[string]bool
	if err := json.Unmarshal(data, &statusMap); err != nil {
		panic(err)
	}

	_, ok := statusMap["is_migrated"]
	if !ok {
		fmt.Println("Key 'is_migrated' not found. Creating new key.")
		statusMap["is_migrated"] = false
	} else {
		statusMap["is_migrated"] = true
	}

	jsonDataBytes, err := json.Marshal(statusMap)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("status.json", jsonDataBytes, 0644); err != nil {
		panic(err)
	}
}

func isMigratedCheck() bool {
	data, err := os.ReadFile("status.json")

	if err != nil {
		panic(err)
	}

	var statusMap map[string]bool
	if err := json.Unmarshal(data, &statusMap); err != nil {
		panic(err)
	}

	_, ok := statusMap["is_migrated"]
	if !ok {
		fmt.Println("Key 'is_migrated' not found. Creating new key.")
		statusMap["is_migrated"] = false
		return false
	}

	return statusMap["is_migrated"]
}

func migrate() {

	db, err := configs.ConnectToDB()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	gormDB := db.GetDB()

	if err := gormDB.AutoMigrate(models.Branch{}, models.Service{}, &models.User{}, &models.Reservation{}, &models.Review{}, &models.Token{}); err != nil {
		log.Fatal(err.Error())
		return
	}

	createAdminUser(gormDB)
	createServices(gormDB)
	createBranches(gormDB)
	createReviews(gormDB)

	setMigrated()
}
