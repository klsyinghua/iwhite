package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Server struct {
	ID             int        `json:"id"`
	Hostname       string     `json:"hostname"`
	IPAddress      string     `json:"ip_address"`
	UUID           string     `json:"uuid"`
	Category       string     `json:"category"`
	Owner          string     `json:"owner"`
	Status         string     `json:"status"`
	ExpirationDate string     `json:"expiration_date"`
	OfflineDate    *time.Time `json:"offline_date"` // 使用指针指向时间类型
}

// QueryServers 方法用于查询服务器数据
func QueryServers(db *gorm.DB, search string) ([]Server, error) {
	var servers []Server
	query := "hostname LIKE ? OR ip_address LIKE ?"
	if err := db.Where(query, "%"+search+"%", "%"+search+"%").Find(&servers).Error; err != nil {
		log.Printf("Error querying servers: %v", err)
		return nil, err
	}

	return servers, nil
}

// CreateServer 方法用于创建新的服务器记录
func (s *Server) CreateServer(db *gorm.DB) error {
	query := "INSERT INTO servers (hostname, ip_address, owner, status, expiration_date, offline_date) VALUES (?, ?, ?, ?, ?, ?)"
	result := db.Exec(query, s.Hostname, s.IPAddress, s.Owner, s.Status, s.ExpirationDate, s.OfflineDate)
	return result.Error
}

// ExistsHostname 检查数据库中是否存在相同的主机名
func (s *Server) ExistsHostname(db *gorm.DB) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM servers WHERE hostname = ?"
	result := db.Raw(query, s.Hostname).Scan(&count)
	return count > 0, result.Error
}

// ExistsIPAddress checks if a server with the given IP address already exists in the database.
func (s *Server) ExistsIPAddress(db *gorm.DB) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM servers WHERE ip_address = ?"
	result := db.Raw(query, s.IPAddress).Scan(&count)
	return count > 0, result.Error
}

// DeleteServer 方法用于删除服务器记录
func (s *Server) DeleteServer(db *gorm.DB, search string) error {
	query := "DELETE FROM servers WHERE hostname LIKE ? OR ip_address LIKE ?"
	result := db.Exec(query, "%"+search+"%", "%"+search+"%")
	return result.Error
}
