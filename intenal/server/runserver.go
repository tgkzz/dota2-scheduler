package server

import (
	"dota2scheduler/config"
	"log"
	"net/http"
)

func Runserver(config config.Config, mux http.Handler, infoLogger *log.Logger) error {
	server := &http.Server{
		Addr:    config.Host + config.Port,
		Handler: mux,
	}

	infoLogger.Printf("server is listening on http://%s%s", config.Host, config.Port)

	err := server.ListenAndServe()

	return err
}
