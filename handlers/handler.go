package handlers

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"iwhite/models"
	"net/http"
)

type ServerHandler struct {
	//db *sql.DB
}

// NewServerHandler 创建一个新的 ServerHandler 实例
func NewServerHandler(db *sql.DB) *ServerHandler {
	return &ServerHandler{}

}
func (h *ServerHandler) GetServerHandler(c echo.Context) error {
	search := c.QueryParam("search")
	fmt.Println("Database connection:", c.Get("db").(*sql.DB)) // 调试输出，检查是否获取到了数据库连接
	servers, err := models.QueryServers(c.Get("db").(*sql.DB), search)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query servers")
	}

	return c.JSON(http.StatusOK, servers)
}

func (h *ServerHandler) CreateServerHandler(c echo.Context) error {
	server := new(models.Server)

	if err := c.Bind(server); err != nil {
		return err
	}

	existsHostname, err := server.ExistsHostname(c.Get("db").(*sql.DB))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert server")
	}

	existsIPAddress, err := server.ExistsIPAddress(c.Get("db").(*sql.DB))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert server")
	}

	if existsHostname && existsIPAddress {
		return c.String(http.StatusConflict, "Server hostname and IP address already exist")
	} else if existsHostname {
		return c.String(http.StatusConflict, "Server hostname already exists")
	} else if existsIPAddress {
		return c.String(http.StatusConflict, "Server IP address already exists")
	}

	err = server.InsertServer(c.Get("db").(*sql.DB))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert server")
	}

	return c.String(http.StatusCreated, "Server information inserted successfully")
}

func (h *ServerHandler) DeleteServerHandler(c echo.Context) error {
	search := c.QueryParam("search")
	err := models.DeleteServer(c.Get("db").(*sql.DB), search)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete server")
	}

	return c.String(http.StatusOK, "Server information deleted successfully")
}
