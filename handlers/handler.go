package handlers

import (
	"database/sql"
	"iwhite/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
	log.Printf("Querying servers with search term: %s", search)
	// 获取数据库连接
	dbConn := c.Get("db").(*sql.DB)
	if dbConn == nil {
		log.Printf("Database connection not found") // 输出日志，表示未找到数据库连接
		return c.String(http.StatusInternalServerError, "Failed to query servers")
	}
	// 调用 models.QueryServers 函数查询服务器数据，并输出查询语句和查询参数的日志
	query := "SELECT DISTINCT id, hostname, ip_address,uuid,category , owner, status, expiration_date, offline_date FROM servers WHERE hostname LIKE ? OR ip_address LIKE ? "
	log.Printf("Executing query: %s, args: [%s %s]", query, "%"+search+"%", "%"+search+"%")

	servers, err := models.QueryServers(dbConn, search)
	if err != nil {
		log.Printf("Error querying servers: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to query servers")
	}

	return c.JSON(http.StatusOK, servers)
	// servers, err := models.QueryServers(c.Get("db").(*sql.DB), search)
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, "Failed to query servers")
	// }

	// return c.JSON(http.StatusOK, servers)
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
