package server

import (
	"dota2scheduler/intenal/models"
	"dota2scheduler/intenal/repository/server"
)

type ServerService struct {
	repo server.Server
}

type Serverer interface {
	UpdateServerInfo(server models.Server) error
	GetServerInfo() (models.Server, error)
	RequestServerInfo() (models.Server, error)
}

func NewServerService(repo server.Server) *ServerService {
	return &ServerService{repo: repo}
}
