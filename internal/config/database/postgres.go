package database

import (
	"fmt"
	"log"

	"github.com/savanyv/account-service-api/internal/config"
	"github.com/savanyv/account-service-api/internal/models"
	"github.com/savanyv/account-service-api/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		utils.LogCritical("DATABASE", "Failed to get db instance: %v", err)
		log.Fatal(err)
	}

	if err := db.AutoMigrate(
		&models.Customer{},
		&models.Transaction{},
	); err != nil {
		utils.LogCritical("DATABASE", "Failed to migrate database: %v", err)
		log.Fatal(err)
	}

	DB = db

	utils.LogInfo("DATABASE", "Database connected successfully")
	return db
}
