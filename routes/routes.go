package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"iwhite/handlers"
)

// SetupRoutes 设置应用程序的路由
func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	serverHandler := handlers.NewServerHandler(db) // 路由组示例
	api := e.Group("/api")
	// 用户相关路由
	api.GET("/servers", serverHandler.GetServerHandler)
	api.GET("/servers", serverHandler.GetServerHandler)
	api.GET("/servers/:identifier", serverHandler.GetServerByHostnameOrIP)
	api.POST("/servers", serverHandler.CreateServerHandler)
	api.DELETE("/servers", serverHandler.DeleteServerHandler)
	api.PUT("/servers", serverHandler.UpdateServerHandler)
	// 其他路由...
}
