package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
	"iwhite/models"
	"net/http"
)

var (
	serverCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "servers_total",
		Help: "Total number of servers",
	})
)

type ServerHandler struct {
	db *gorm.DB
}

func NewServerHandler(db *gorm.DB) *ServerHandler {
	return &ServerHandler{
		db: db,
	}
}
func (h *ServerHandler) GetServersHandler(c echo.Context) error {
	server := &models.Server{}
	servers, err := server.GetAllServers(h.db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query servers")
	}
	return c.JSON(http.StatusOK, servers)
}
func (h *ServerHandler) GetServerHandler(c echo.Context) error {
	id := c.Param("id")
	server := &models.Server{}
	err := server.GetServer(h.db, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query server")
	}
	return c.JSON(http.StatusOK, server)
}

func (h *ServerHandler) CreateServerHandler(c echo.Context) error {
	var newServer models.Server
	if err := c.Bind(&newServer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON data"})
	}

	// Check for required fields
	if newServer.HostName == "" || newServer.IPv4Address == "" || newServer.Owner == "" || newServer.Environment == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required fields"})
	}
	if err := newServer.CreateServer(h.db); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, "Server created successfully")
}

func (h *ServerHandler) DeleteServerHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "无效的服务器ID")
	}
	server := &models.Server{}
	err := server.DeleteServer(h.db, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete server")
	}

	return c.String(http.StatusOK, "Server information deleted successfully")
}

func (h *ServerHandler) UpdateServerHandler(c echo.Context) error {
	// Parse request body to get the updated server details
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "无效的服务器ID")
	}
	var updatedServer models.Server
	if err := c.Bind(&updatedServer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON data"})
	}

	// Check for required fields
	if updatedServer.HostName == "" || updatedServer.IPv4Address == "" || updatedServer.Owner == "" || updatedServer.Environment == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required fields"})
	}
	// Create a Server instance and install ID vaule of the server to be updated
	existingServer := models.Server{ID: id}
	//log.Printf("existingServer: %v", existingServer)
	//updatedServer.ID = id
	// Update the server record
	err := existingServer.UpdateServer(h.db, updatedServer)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Server not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update server"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Server updated successfully"})
}
