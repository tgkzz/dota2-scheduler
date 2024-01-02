package service

import (
	"dota2scheduler/intenal/repository"
	"dota2scheduler/intenal/service/fill"
	"dota2scheduler/intenal/service/filter"
)

type Service struct {
	fill.Filler
	filter.Filterer
}

func NewService(repo repository.Repo) *Service {
	return &Service{
		fill.NewFillService(repo),
		filter.NewFilterService(repo),
	}
}
