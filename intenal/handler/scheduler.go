package handler

import (
	"log"
	"time"
)

// Scheduler TODO: realize the work of the scheduler using cron lib
func (h *Handler) Scheduler() {
	for {
		regions := []string{"europe"}

		leaderboard, err := h.service.MakeRequest(regions)
		if err != nil {
			log.Fatalf("Error while making request to dotaapi: %s", err)
		}

		server, err := h.service.RequestServerInfo()
		if err != nil {
			log.Fatalf("Error while requesting server info %s", err)
		}

		if err := h.service.UpdateServerInfo(server); err != nil {
			log.Fatalf("Error while updating server infor %s", err)
		}

		if err := h.service.UpdateLeaderboard(leaderboard[0]); err != nil {
			log.Fatalf("Error while updating leaderboard: %s", err)
		}

		sleepTime := server.NextScheduledPostTime - server.ServerTime

		log.Printf("next scheduled post time is %d\n", server.NextScheduledPostTime)

		log.Printf("server time is %d\n", server.ServerTime)

		log.Printf("next scheduled post time in seconds is %d\n", sleepTime)

		sleepTimeInMinute := sleepTime / 60

		log.Printf("Request has been handled! Next scheduled request will be in %d minute\n", sleepTimeInMinute)

		log.Println("------------------------------------")

		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
