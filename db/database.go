package db

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) Initialize(connStr string, logger echo.Logger) {
	var err error
	d.DB, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to the models: %v", err)
	}
	return
}

func (d *Database) GetConnection() *gorm.DB {
	return d.DB
}
