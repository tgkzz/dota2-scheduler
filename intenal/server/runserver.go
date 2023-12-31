package server

import (
	"dota2scheduler/config"
	"log"
	"net/http"
)

func Runserver(config config.Config, mux http.Handler) error {
	server := &http.Server{
		Addr:    config.Host + config.Port,
		Handler: mux,
	}

	log.Printf("server is listening on http://%s%s", config.Host, config.Port)

	err := server.ListenAndServe()

	return err
}
