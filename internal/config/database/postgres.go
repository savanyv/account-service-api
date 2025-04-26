package database

import (
	"fmt"

	"github.com/savanyv/account-service-api/internal/config"
	"github.com/savanyv/account-service-api/internal/models"
	"github.com/savanyv/account-service-api/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", // Tambahkan sslmode=disable
        config.DBHost,
        config.DBUser,
        config.DBPassword,
        config.DBName,
        config.DBPort,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        utils.LogCritical("Database", "Failed to connect to database: %v", err)
        return nil, err
    }

    if err := db.AutoMigrate(&models.Customer{}); err != nil {
        utils.LogError("Database", "Failed to migrate Customer table: %v", err)
        return nil, err
    }

    if err := db.AutoMigrate(&models.Transaction{}); err != nil {
        utils.LogError("Database", "Failed to migrate Transaction table: %v", err)
        return nil, err
    }

    DB = db
    return db, nil
}
