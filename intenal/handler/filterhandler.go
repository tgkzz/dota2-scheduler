package handler

import (
	"dota2scheduler/intenal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) FitlerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var filter models.Filter
		if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
			h.errorLogger.Printf("error while decoding json of filter %s", err)
			ServerError(w, err, http.StatusInternalServerError)
			return
		}

		fmt.Println(filter)

		ldbrd, err := h.service.GetLeaderboardByFilter(filter)
		if err != nil {
			if err == models.ErrEmptyFilter {
				h.errorLogger.Printf("%s", err)
				ServerError(w, err, http.StatusBadRequest)
				return
			}
			h.errorLogger.Printf("error while filtering leaderboard", err)
			ServerError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ldbrd); err != nil {
			h.errorLogger.Printf("error while encoding result %s", err)
			ServerError(w, err, http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
