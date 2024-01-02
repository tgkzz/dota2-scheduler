package handler

import "net/http"

func (h *Handler) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api", h.IndexHandler)
	mux.HandleFunc("/api/filter", h.FitlerHandler)

	return h.ServeHTTP(mux)
}
