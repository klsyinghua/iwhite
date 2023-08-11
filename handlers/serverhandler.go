package handlers

import (
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
	"iwhite/models"
	"net/http"

	"github.com/labstack/echo/v4"
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
func (h *ServerHandler) GetServerHandler(c echo.Context) error {
	servers, err := (&models.Server{}).QueryAllServers(h.db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query servers")
	}
	return c.JSON(http.StatusOK, servers)
}
func (h *ServerHandler) GetServerByHostnameOrIP(c echo.Context) error {
	identifier := c.Param("identifier")

	server := &models.Server{}
	err := server.QueryServerByHostnameOrIP(h.db, identifier)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query server")
	}

	return c.JSON(http.StatusOK, server)
}

func (h *ServerHandler) CreateServerHandler(c echo.Context) error {
	server := new(models.Server)

	if err := c.Bind(server); err != nil {
		return err
	}

	existsHostname, err := server.ExistsHostname(h.db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert server")
	}

	existsIPAddress, err := server.ExistsIPAddress(h.db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert server")
	}

	if existsHostname && existsIPAddress {
		return c.String(http.StatusConflict, "Server "+server.Hostname+" and IP address already exist")
	} else if existsHostname {
		return c.String(http.StatusConflict, "Server hostname already exists")
	} else if existsIPAddress {
		return c.String(http.StatusConflict, "Server IP address already exists")
	}

	err = server.CreateServer(h.db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert server")
	}

	return c.String(http.StatusCreated, "Server information inserted successfully")
}
func (h *ServerHandler) DeleteServerHandler(c echo.Context) error {
	search := c.QueryParam("search")

	server := &models.Server{}
	err := server.DeleteServer(h.db, search)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete server")
	}

	return c.String(http.StatusOK, "Server information deleted successfully")
}

func (h *ServerHandler) UpdateServerHandler(c echo.Context) error {
	serverID := c.Param("id") // Assuming you can extract the server ID from the URL

	// Get the existing server data from the database using the server ID
	existingServer := &models.Server{}
	if err := h.db.First(existingServer, serverID).Error; err != nil {
		return c.String(http.StatusNotFound, "Server not found")
	}

	// Instantiate a Server struct with updated data
	updatedServer := new(models.Server)
	if err := c.Bind(updatedServer); err != nil {
		return err
	}

	// Update the fields of the existing server with fields from updatedServer
	existingServer.Hostname = updatedServer.Hostname
	existingServer.IPAddress = updatedServer.IPAddress
	// ... Update other fields as needed ...

	// Save the updated server data back to the database
	if err := h.db.Save(existingServer).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update server")
	}

	return c.String(http.StatusOK, "Server information updated successfully")
}

var serverStatus *prometheus.GaugeVec

func init() {
	serverStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "server_status",
			Help: "Server status (0: offline, 1: online)",
		},
		[]string{"hostname", "ip_address", "uuid", "category", "owner"},
	)
	prometheus.MustRegister(serverStatus)

}
func UpdateMetricsForServers(database *gorm.DB, servers []models.Server) {
	for _, server := range servers {
		statusValue := 0.0
		if server.Status == "Running" {
			statusValue = 1.0
		}
		labels := prometheus.Labels{
			"hostname":   server.Hostname,
			"ip_address": server.IPAddress,
			"uuid":       server.UUID,
			"category":   server.Category,
			"owner":      server.Owner,
		}
		serverStatus.With(labels).Set(statusValue)
	}

	// 设置 Prometheus 的 serverCount 指标
	serverCount.Set(float64(len(servers)))
}

func (h *ServerHandler) GetServerMetricsHandler(c echo.Context) error {
	servers, err := (&models.Server{}).QueryAllServers(h.db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query servers")
	}

	UpdateMetricsForServers(h.db, servers)

	return c.JSON(http.StatusOK, echo.Map{"message": "Server metrics updated"})
}
