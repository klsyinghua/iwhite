package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Server struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	Name        string    `gorm:"type:varchar(255)"`
	VCPUs       int       `gorm:"type:int"`
	RAM         int       `gorm:"type:int"`
	Disk        int       `gorm:"type:int"`
	IPv4Address string    `gorm:"type:varchar(255)"`
	IPv6Address string    `gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"type:timestamp"` //
	UpdatedAt   time.Time `gorm:"type:timestamp"` //
	TerminateAt time.Time `gorm:"type:timestamp"` // 自动销毁时间
	PowerState  int       `gorm:"type:int"`
	HostName    string    `gorm:"type:varchar(255)"`
	HostStatus  string    `gorm:"type:varchar(255)"` //
	Owner       string    `gorm:"type:varchar(255)"` // 添加拥有者字段
	Environment string    `gorm:"type:varchar(255)"` // 所属环境字段
	Features    string    `gorm:"type:varchar(255)"` //
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

	result := db.Model(s).Where("hostname = ?", s.HostName).Count(&count)
	return count > 0, result.Error
}

// ExistsIPAddress checks if a server with the given IP address already exists in the database.
func (s *Server) ExistsIPAddress(db *gorm.DB) (bool, error) {
	var count int64
	result := db.Model(s).Where("ip_address = ?", s.IPv4Address).Count(&count)
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
	result := db.Find(&servers)
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
