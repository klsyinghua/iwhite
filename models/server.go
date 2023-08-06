package models

import (
	"database/sql"
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

// queryServers 是一个单独的数据库查询函数
func queryServers(db *sql.DB, query string, args ...interface{}) ([]Server, error) {
	log.Printf("Executing query: %s, args: %v", query, args)

	rows, err := db.Query(query, args...)

	if err != nil {
		log.Printf("Error executing query: %v", err)

		return nil, err
	}
	defer rows.Close()

	var servers []Server
	for rows.Next() {
		var server Server
		offlineDate := sql.NullTime{}
		err := rows.Scan(&server.ID, &server.Hostname, &server.IPAddress, &server.UUID, &server.Category, &server.Owner, &server.Status, &server.ExpirationDate, &offlineDate)
		if err != nil {
			log.Printf("Error scanning row: %v", err)

			return nil, err
		}
		if offlineDate.Valid {
			server.OfflineDate = &offlineDate.Time
		}
		servers = append(servers, server)
	}
	log.Printf("Query executed successfully, retrieved %d rows", len(servers))
	return servers, nil
}

// GetServer 方法用于基于IP地址或主机名进行模糊搜索
func (s *Server) GetServer(db *sql.DB, search string) ([]Server, error) {
	query := "SELECT DISTINCT id, hostname, ip_address, owner, status, expiration_date, offline_date FROM servers WHERE hostname LIKE ? OR ip_address LIKE ?"
	return queryServers(db, query, "%"+search+"%", "%"+search+"%")
}

// CreateServer 方法用于创建新的服务器记录
func (s *Server) CreateServer(db *sql.DB, server *Server) error {
	query := "INSERT INTO servers (hostname, ip_address, owner, status, expiration_date, offline_date) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, server.Hostname, server.IPAddress, server.Owner, server.Status, server.ExpirationDate, server.OfflineDate)
	return err
}

// DeleteServer 方法用于删除服务器记录
func (s *Server) DeleteServer(db *sql.DB, search string) error {
	query := "DELETE FROM servers WHERE hostname LIKE ? OR ip_address LIKE ?"
	_, err := db.Exec(query, "%"+search+"%", "%"+search+"%")
	return err
}

// InsertServer 将当前 Server 对象的信息插入数据库
func (s *Server) InsertServer(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO servers (hostname, ip_address, owner, status, expiration_date, offline_date) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.Hostname, s.IPAddress, s.Owner, s.Status, s.ExpirationDate, s.OfflineDate)
	if err != nil {
		return err
	}

	return nil
}

// ExistsHostname 检查数据库中是否存在相同的主机名
func (s *Server) ExistsHostname(db *sql.DB) (bool, error) {
	query := "SELECT COUNT(*) FROM servers WHERE hostname = ?"
	var count int
	err := db.QueryRow(query, s.Hostname).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// ExistsIPAddress checks if a server with the given IP address already exists in the database.
func (s *Server) ExistsIPAddress(db *sql.DB) (bool, error) {
	// Implement the logic to check if a server with the given IP address exists in the database here.
	query := "SELECT COUNT(*) FROM servers WHERE ip_address = ?"
	var count int
	err := db.QueryRow(query, s.IPAddress).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// DeleteServer deletes the server record from the database based on the given search criteria (hostname or IP address).
func DeleteServer(db *sql.DB, search string) error {
	// Implement the logic to delete the server record from the database based on the search criteria here.
	// ...
	return nil
}

// QueryServers retrieves a list of servers from the database based on the given search criteria (hostname or IP address).
func QueryServers(db *sql.DB, search string) ([]Server, error) {
	// Implement the logic to query the servers from the database based on the search criteria here.
	// ...
	query := "SELECT DISTINCT id, hostname, ip_address,uuid,category , owner, status, expiration_date, offline_date FROM servers WHERE hostname LIKE ? OR ip_address LIKE ? "
	log.Printf("Executing query: %s, args: [%s %s]", query, "%"+search+"%", "%"+search+"%")

	servers, err := queryServers(db, query, "%"+search+"%", "%"+search+"%")
	if err != nil {
		log.Printf("Error executing queryServers: %v", err) // 输出错误信息
		return nil, err
	}

	return servers, nil
}
