package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"

	"iwhite/config"
	"iwhite/db"
	"iwhite/routes"
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
	// 启动服务器并监听端口
	log.Fatal(e.Start(":8080"))
}
