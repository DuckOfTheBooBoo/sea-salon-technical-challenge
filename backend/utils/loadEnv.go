package utils

import (
	"log"
	"path/filepath"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	dir, err := os.Getwd()
    if err != nil {
        return err
    }

    envPath := filepath.Join(dir, ".env")

    err = godotenv.Load(envPath)
    if err != nil {
        log.Printf("Error loading .env file: %v\n", err)
    }
    return err
}