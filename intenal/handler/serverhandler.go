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
			h.errorLogger.Printf("error while getting server info %s", err)
			ServerError(w, err, http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(server); err != nil {
			h.errorLogger.Printf("error while encoding the result %s", err)
			ServerError(w, err, http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
