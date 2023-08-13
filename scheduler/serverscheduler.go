package scheduler

import (
	"gorm.io/gorm"
	"iwhite/handlers"
	"iwhite/models"
	"log"
	"time"
)

func StartScheduler(database *gorm.DB) {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				servers, err := (&models.Server{}).QueryAllServers(database)
				if err != nil {
					log.Println("Failed to query servers:", err)
					continue
				}
				handlers.UpdateMetricsForServers(database, servers)
			}
		}
	}()
}
