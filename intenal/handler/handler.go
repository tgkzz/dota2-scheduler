package handler

import (
	"dota2scheduler/intenal/service"
	"log"
)

type Handler struct {
	service     *service.Service
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewHandler(service *service.Service, infoLogger *log.Logger, errorLogger *log.Logger) *Handler {
	return &Handler{
		service:     service,
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}
}
