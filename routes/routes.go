package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
	"iwhite/handlers"
)

// SetupRoutes 设置应用程序的路由
func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	serverHandler := handlers.NewServerHandler(db) // 路由组示例
	api := e.Group("/api")
	metrics := e.Group("/metrics")

	// 用户相关路由
	api.GET("/servers", serverHandler.GetServersHandler)

	api.GET("/servers/:id", serverHandler.GetServerHandler)
	api.POST("/servers", serverHandler.CreateServerHandler)
	api.DELETE("/servers/:id", serverHandler.DeleteServerHandler)
	api.PUT("/servers/:id", serverHandler.UpdateServerHandler)
	// 其他路由...
	//api.GET("/search", serverHandler.SearchHandler)
	//api.POST("/query", serverHandler.QueryHandler)
	//api.POST("/annotations", serverHandler.AnnotationsHandler)
	//api.GET("/", serverHandler.HelloHandler)
	//
	metrics.GET("/server", echo.WrapHandler(promhttp.Handler()))
	metrics.POST("/sever/updatestatus", serverHandler.GetServerMetricsHandler)
}
