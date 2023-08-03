package main

import (
	"iwhite/config"
	"iwhite/db"
	"iwhite/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize configuration
	configPath := "config.yaml"
	if err := config.LoadConfig(configPath); err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}
	e := echo.New()
	db.InitDB(e.Logger)
	// 添加日志中间件
	e.Use(middleware.Logger())
	// 将数据库连接传递给处理函数
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db.GetDB())
			return next(c)
		}
	})

	// 设置路由
	routes.SetupRoutes(e)
	// 启动服务器并监听端口
	log.Fatal(e.Start(":8080"))
}
