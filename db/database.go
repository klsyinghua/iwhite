package db

import (
	"gorm.io/driver/mysql"
	"iwhite/config"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDatabase(logger echo.Logger) {
	configPath := "config.yaml" // Replace with the path to your configuration file
	if err := config.LoadConfig(configPath); err != nil {
		logger.Fatalf("Failed to load config file: %v", err)
	}
	connStr := config.GetMySQLConfig()
	var err error
	db, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to the models: %v", err)
	}
	return
}

// GetDB returns the models connection.
func GetDB() *gorm.DB {
	return db
}
