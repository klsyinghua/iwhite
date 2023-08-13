package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"iwhite/config"
	"iwhite/db"
	"iwhite/models"
	"iwhite/routes"
	"iwhite/scheduler"
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
	err := database.AutoMigrate(&models.Server{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// 添加日志中间件
	e.Use(middleware.Logger())

	// 设置路由
	routes.SetupRoutes(e, database)
	// 启动定时更新指标的调度器
	scheduler.StartScheduler(database)
	log.Fatal(e.Start(":8080"))
}
