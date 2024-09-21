package database

import (
	"fmt"
	"api-campaign/user"
	"api-campaign/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_USER"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_NAME"),
		utils.GetConfig("DB_PORT"),
		utils.GetConfig("DB_SSL_MODE"),
		utils.GetConfig("DB_TIMEZONE"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	err := DB.AutoMigrate(&user.User{})

	if err != nil {
		log.Printf("Error when migrating the database: %v", err)
	}
}