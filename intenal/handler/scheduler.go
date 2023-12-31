package handler

import (
	"log"
	"time"
)

func (h *Handler) Scheduler() {
	for {
		regions := []string{"europe"}

		leaderboard, err := h.service.MakeRequest(regions)
		if err != nil {
			log.Fatalf("Error while making request to dotaapi: %s", err)
		}

		if err := h.service.UpdateLeaderboard(leaderboard[0]); err != nil {
			log.Fatalf("Error while updating leaderboard: %s", err)
		}

		time.Sleep(5 * time.Minute)
	}
}
