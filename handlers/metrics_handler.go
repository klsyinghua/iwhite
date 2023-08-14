package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
	"iwhite/models"
	"net/http"
	"time"
)

var serverStatus *prometheus.GaugeVec

func init() {
	serverStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "server_status",
			Help: "Server status (0: offline, 1: online)",
		},
		[]string{"uuid", "name", "hostname", "ipv4_address", "host_status", "create_date", "update_date", "terminate_date", "env", "owner", "features"},
	)
	prometheus.MustRegister(serverStatus)

}
func UpdateMetricsForServers(database *gorm.DB, servers []models.Server) {
	for _, server := range servers {
		//log.Printf("Processing server: %s", server.Name)
		statusValue := 0.0
		if server.HostStatus == "ACTIVE" {
			statusValue = 1.0
		}
		// 将时间转换为东8区时间
		loc, _ := time.LoadLocation("Asia/Shanghai") // 或者 "Asia/Chongqing"
		labels := prometheus.Labels{
			"uuid":           server.ID,
			"name":           server.Name,
			"hostname":       server.HostName,
			"ipv4_address":   server.IPv4Address,
			"host_status":    server.HostStatus,
			"create_date":    server.CreatedAt.In(loc).Format("2006-01-02 15:04:05"),
			"update_date":    server.UpdatedAt.In(loc).Format("2006-01-02 15:04:05"),
			"terminate_date": server.TerminateAt.In(loc).Format("2006-01-02 15:04:05"),
			"env":            server.Environment,
			"owner":          server.Owner,
			"features":       server.Features,
		}
		//log.Printf("Setting labels: %+v", labels)
		serverStatus.With(labels).Set(statusValue)
	}

	// 设置 Prometheus 的 serverCount 指标
	serverCount.Set(float64(len(servers)))
}

func (h *ServerHandler) GetServerMetricsHandler(c echo.Context) error {
	servers, err := (&models.Server{}).GetAllServers(h.db)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query servers")
	}

	UpdateMetricsForServers(h.db, servers)

	return c.JSON(http.StatusOK, echo.Map{"message": "Server metrics updated"})
}
