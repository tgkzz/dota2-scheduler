package service

import (
	"dota2scheduler/intenal/repository"
	"dota2scheduler/intenal/service/fill"
	"dota2scheduler/intenal/service/filter"
	"dota2scheduler/intenal/service/server"
)

type Service struct {
	fill.Filler
	filter.Filterer
	server.Serverer
}

func NewService(repo repository.Repo) *Service {
	return &Service{
		fill.NewFillService(repo),
		filter.NewFilterService(repo),
		server.NewServerService(repo),
	}
}
