package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"iwhite/config"
	"iwhite/db"
	"iwhite/handlers"
	"iwhite/models"
	"iwhite/routes"
	"log"
	"time"
)

func main() {
	// Initialize configuration
	configPath := "config.yaml"
	if err := config.LoadConfig(configPath); err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}
	e := echo.New()
	db.InitializeDatabase(e.Logger)
	if db.GetDB() == nil {
		log.Fatal("Failed to initialize database connection")
	} else {
		log.Println("Database connection successful")
	}
	database := db.GetDB()

	// 添加日志中间件
	e.Use(middleware.Logger())

	// 设置路由
	routes.SetupRoutes(e, database)
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				servers, err := (&models.Server{}).QueryAllServers(database)
				if err != nil {
					log.Println("Failed to query servers:", err)
					continue
				}
				handlers.UpdateMetricsForServers(database, servers)
			}
		}
	}()
	log.Fatal(e.Start(":8080"))
}
