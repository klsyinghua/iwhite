package db

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"iwhite/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(logger echo.Logger) {
	configPath := "config.yaml" // Replace with the path to your configuration file
	if err := config.LoadConfig(configPath); err != nil {
		logger.Fatalf("Failed to load config file: %v", err)
	}

	connStr := config.GetMySQLConfig()
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		logger.Fatalf("Failed to connect to the models: %v", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatalf("Failed to ping the models: %v", err)
	}
}

// GetDB returns the models connection.
func GetDB() *sql.DB {
	return db
}
