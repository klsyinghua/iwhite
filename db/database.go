package db

import (
	"database/sql"
	"iwhite/config"

	"github.com/labstack/echo/v4"

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
	// _, err = db.Exec("SELECT * from servers")
	// if err != nil {
	// 	logger.Fatalf("Failed to execute test query: %v", err)
	// }
	// _, err = db.Exec("USE iwhite")
	// if err != nil {
	// 	logger.Fatalf("Failed to select database: %v", err)
	// }
}

// GetDB returns the models connection.
func GetDB() *sql.DB {
	return db
}
