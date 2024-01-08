package handler

import (
	"log"
	"time"
)

// TODO: change the scheduled time, to time which includes in the server
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

		log.Print("new req\n")

		time.Sleep(5 * time.Minute)
	}
}
