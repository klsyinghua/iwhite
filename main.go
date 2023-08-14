package main

import (
	"fmt"
	"iwhite/config"
	"iwhite/scheduler"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"iwhite/db"
	"iwhite/routes"
)

func main() {
	// Initialize configuration
	var appConfig config.AppConfig
	var dbInstance db.Database
	var configPath string

	if len(os.Args) > 1 {
		configPath = os.Args[1] // 获取命令行参数作为配置文件路径
	}
	if err := appConfig.InitConfig(configPath); err != nil {
		fmt.Printf("Failed to load config file: %v\n", err)
		return
	}
	connectionString := appConfig.GetMysqlConnectionString(appConfig.GetEnvConfig())
	e := echo.New()
	dbInstance.Initialize(connectionString, e.Logger)
	database := dbInstance.GetConnection()
	authUsername := "username"
	authPassword := "password"
	authMiddleware := middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == authUsername && password == authPassword {
			return true, nil
		}
		return false, nil
	})
	// 添加日志中间件
	e.Use(middleware.Logger())
	e.Use(authMiddleware)

	// 设置路由
	routes.SetupRoutes(e, database)
	// 启动定时更新指标的调度器
	schedulerInstance := scheduler.NewScheduler(database)
	schedulerInstance.Start()
	log.Fatal(e.Start(":8080"))
}
