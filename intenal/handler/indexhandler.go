package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ldbrd, err := h.service.ShowLeaderboard()
		if err != nil {
			h.errorLogger.Printf("error while showing current leaderboard %s", err)
			ServerError(w, err, http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(ldbrd); err != nil {
			h.errorLogger.Printf("error while showing current leaderboard %s", err)
			ServerError(w, err, http.StatusInternalServerError)
		}

	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}
