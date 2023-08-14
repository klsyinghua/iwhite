package scheduler

import (
	"gorm.io/gorm"
	"iwhite/handlers"
	"iwhite/models"
	"log"
	"time"
)

// Todo
// - config file conve
// Scheduler 结构体Scheduler time
type Scheduler struct {
	database *gorm.DB
}

// NewScheduler 创建 Scheduler 实例
func NewScheduler(database *gorm.DB) *Scheduler {
	return &Scheduler{
		database: database,
	}
}

// Start 启动调度器
func (s *Scheduler) Start() {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.queryAndHandleServers()
			}
		}
	}()
}

func (s *Scheduler) queryAndHandleServers() {
	servers, err := (&models.Server{}).GetAllServers(s.database)
	if err != nil {
		log.Println("Failed to query servers:", err)
		return
	}
	handlers.UpdateMetricsForServers(s.database, servers)
}
