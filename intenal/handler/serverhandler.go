package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) ShowServerInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		server, err := h.service.GetServerInfo()
		if err != nil {
			ServerError(w, err, http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(server); err != nil {
			ServerError(w, err, http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
