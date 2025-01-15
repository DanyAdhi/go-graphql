package db

import (
	"fmt"
	"log"

	"github.com/danyadhi/go-graphql/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dbConnStr := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s`,
		config.AppConfig.DB_HOST,
		config.AppConfig.DB_USER,
		config.AppConfig.DB_PASSWORD,
		config.AppConfig.DB_NAME,
		config.AppConfig.DB_PORT,
		config.AppConfig.DB_SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}
