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
	result := db.Create(s)
	return result.Error
}

// ExistsHostname 检查数据库中是否存在相同的主机名
func (s *Server) ExistsHostname(db *gorm.DB) (bool, error) {
	var count int64

	result := db.Model(s).Where("hostname = ?", s.Hostname).Count(&count)
	return count > 0, result.Error
}

// ExistsIPAddress checks if a server with the given IP address already exists in the database.
func (s *Server) ExistsIPAddress(db *gorm.DB) (bool, error) {
	var count int64
	result := db.Model(s).Where("ip_address = ?", s.IPAddress).Count(&count)
	return count > 0, result.Error
}

// DeleteServer 方法用于删除服务器记录
func (s *Server) DeleteServer(db *gorm.DB, identifier string) error {
	query := "hostname LIKE ? OR ip_address LIKE ?"
	result := db.Where(query, "%"+identifier+"%", "%"+identifier+"%").Delete(&Server{})
	return result.Error
}

// QueryAllServers 获取所有服务器的信息列表
func (s *Server) QueryAllServers(db *gorm.DB) ([]Server, error) {
	var servers []Server
	result := db.Find(&servers).Limit(10)
	if result.Error != nil {
		log.Printf("Error querying all servers: %v", result.Error)
		return nil, result.Error
	}
	return servers, nil
}

func (s *Server) QueryServerByHostnameOrIP(db *gorm.DB, identifier string) error {
	query := "hostname LIKE ? OR ip_address LIKE ?"
	err := db.Where(query, "%"+identifier+"%", "%"+identifier+"%").First(s).Error
	return err
}
