package service

import (
	"dota2scheduler/intenal/repository"
	"dota2scheduler/intenal/service/fill"
)

type Service struct {
	fill.Filler
}

func NewService(repo repository.Repo) *Service {
	return &Service{
		fill.NewFillService(repo),
	}
}
