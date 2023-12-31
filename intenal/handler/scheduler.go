package handler

import (
	"net/http"
	"time"
)

func (h *Handler) Scheduler() {
	for {
		regions := []string{"europe"}

		leaderboard, err := h.service.MakeRequest(regions)
		if err != nil {
			ServerError(err, http.StatusInternalServerError)
		}

		if err := h.service.UpdateLeaderboard(leaderboard[0]); err != nil {
			ServerError(err, http.StatusInternalServerError)
		}

		time.Sleep(5 * time.Minute)
	}
}
