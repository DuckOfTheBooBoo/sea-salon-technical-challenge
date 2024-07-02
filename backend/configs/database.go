package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	GetDB() *gorm.DB
}

type mysqlDB struct {
	db *gorm.DB
}

func (db *mysqlDB) GetDB() *gorm.DB {
	return db.db
}

func ConnectToDB() (Database, error) {
	
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	retries := os.Getenv("MAX_RETRIES")
	retryDel := os.Getenv("RETRY_DELAY")

	log.Println("DB_NAME:", os.Getenv("DB_NAME"))
	log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("DB_USER:", os.Getenv("DB_USER"))
	log.Println("DB_PASS:", os.Getenv("DB_PASS"))
	log.Println("MAX_RETRIES:", os.Getenv("MAX_RETRIES"))
	
	maxRetries, err := strconv.Atoi(retries)
	if err != nil {
		return nil, err
	}
	
	retryDelay, err := strconv.Atoi(retryDel)
	if err != nil {
		return nil, err
	}
	
	retryTime := time.Duration(retryDelay) * time.Second

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	
	for attempts := 0; attempts < maxRetries; attempts++ {
		gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to database")
			return &mysqlDB{db: gormDB}, nil
		}

		log.Printf("Failed to connect to database: %v. Retrying in %v...", err, retryDelay)
		time.Sleep(retryTime)
	}

	// Return the last error if all retries fail
	return nil, fmt.Errorf("could not connect to database after %d attempts: %v", maxRetries, err)
}
