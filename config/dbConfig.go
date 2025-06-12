package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		getENV("DB_HOST"),
		getENV("DB_USER"),
		getENV("DB_PASSWORD"),
		getENV("DB_NAME"),
		getENV("DB_PORT"),
	)
	var db *gorm.DB
	var err error
	maxtries := 5
	for i := range maxtries {
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err == nil {
			log.Println("✅ Successfully connected to the database")
			break
		}
		log.Printf("❌ Failed to connect to DB (attempt %d/%d): %v", i+1, maxtries, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("💥 Could not connect to the database after %d attempts: %v", maxtries, err)
	}
	DB = db
	return db
}

func getENV(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Environment variable %s not set", key))
	}
	return val
}
