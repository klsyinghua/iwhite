package models

import (
	"errors"
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

// GetAllServers 获取所有服务器的信息列表
func (s *Server) GetAllServers(db *gorm.DB) ([]Server, error) {
	var servers []Server
	result := db.Find(&servers)
	return servers, result.Error
}

// GetServer QueryServer
func (s *Server) GetServer(db *gorm.DB, id string) error {
	var server Server
	result := db.Where("id = ?", id).First(&server)
	return result.Error
}

// GetServerByHostnameOrIP QueryServers with ip_add or hostname or  host 方法用于查询服务器数据
func (s *Server) GetServerByHostnameOrIP(db *gorm.DB, search string) ([]Server, error) {
	var servers []Server
	result := db.Where(Server{HostName: search}).Or(Server{IPv4Address: search}).Or(Server{Name: search}).Find(&servers)
	return servers, result.Error
}

// CreateServer 方法用于创建新的服务器记录
func (s *Server) CreateServer(db *gorm.DB) error {
	// Check if ID already exists
	if exists, err := s.ExistsID(db); err != nil {
		return err
	} else if exists {
		return errors.New("ID already exists")
	}

	// Check if HostName already exists
	if exists, err := s.ExistsHostName(db); err != nil {
		return err
	} else if exists {
		return errors.New("HostName already exists")
	}

	// Check if IPv4Address already exists
	if exists, err := s.ExistsIPv4Address(db); err != nil {
		return err
	} else if exists {
		return errors.New("IPv4Address already exists")
	}

	// Create the server record
	result := db.Create(s)
	return result.Error
}

// DeleteServer 方法用于删除服务器记录
func (s *Server) DeleteServer(db *gorm.DB, id string) error {
	var existingServer Server
	result := db.Where("id = ?", id).First(&existingServer)
	//result := db.First(&existingServer, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Server ID %s not found", id)
			return errors.New("Server not found")
		}
		log.Println("Database error:", result.Error)
		return result.Error
	}
	// Delete the server record
	result = db.Delete(&existingServer)
	return result.Error
}

func (s *Server) UpdateServer(db *gorm.DB, updatedServer Server) error {
	//  重新初始化 Server 结构体 接收数据，增加额外的开小
	//  var existingServer Server
	// 	result := db.Where("id = ?", s.ID).First(&existingServer)
	log.Printf("updatedServer: %+v", updatedServer)

	result := db.Where("id = ?", s.ID).First(&s)
	if result.Error != nil {
		return result.Error
	}
	// Update the server record
	updatedServer.ID = s.ID
	result = db.Model(&s).Updates(updatedServer)
	return result.Error
}

func (s *Server) ExistsID(db *gorm.DB) (bool, error) {
	var count int64
	result := db.Model(s).Where("id = ?", s.ID).Count(&count)
	return count > 0, result.Error
}

func (s *Server) ExistsHostName(db *gorm.DB) (bool, error) {
	var count int64
	result := db.Model(s).Where("host_name = ?", s.HostName).Count(&count)
	return count > 0, result.Error
}

func (s *Server) ExistsIPv4Address(db *gorm.DB) (bool, error) {
	var count int64
	result := db.Model(s).Where("ipv4_address = ?", s.IPv4Address).Count(&count)
	return count > 0, result.Error
}
