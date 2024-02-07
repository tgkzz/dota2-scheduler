package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) RedisHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		result, err := h.service.GetRedisData()
		if err != nil {
			h.errorLogger.Printf("error while getting redis data %s", err)
			ServerError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			h.errorLogger.Printf("error while encoding result %s", err)
			ServerError(w, err, http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}
