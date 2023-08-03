package routes

import (
	"github.com/labstack/echo/v4"
	"iwhite/handlers"
)

// SetupRoutes 设置应用程序的路由
func SetupRoutes(e *echo.Echo) {
	serverHandler := &handlers.ServerHandler{}
	// 路由组示例
	api := e.Group("/api")
	// 用户相关路由
	api.GET("/servers", serverHandler.GetServerHandler)
	api.POST("/servers", serverHandler.CreateServerHandler)
	api.DELETE("/servers", serverHandler.DeleteServerHandler)
	// 其他路由...
}
