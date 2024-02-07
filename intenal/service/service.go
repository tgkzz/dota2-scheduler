package service

import (
	"dota2scheduler/intenal/repository"
	"dota2scheduler/intenal/service/fill"
	"dota2scheduler/intenal/service/filter"
	"dota2scheduler/intenal/service/server"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	fill.Filler
	filter.Filterer
	server.Serverer
}

func NewService(repo repository.Repo, client *redis.Client) *Service {
	return &Service{
		fill.NewFillService(repo, client),
		filter.NewFilterService(repo),
		server.NewServerService(repo),
	}
}
