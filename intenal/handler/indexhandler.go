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
			ServerError(w, err, http.StatusInternalServerError)
		}

		if err = json.NewEncoder(w).Encode(ldbrd); err != nil {
			ServerError(w, err, http.StatusInternalServerError)
		}

	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}
