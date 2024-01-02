package filter

import (
	"dota2scheduler/intenal/models"
	"dota2scheduler/intenal/repository/filter"
)

type FilterService struct {
	repo filter.Filter
}

type Filterer interface {
	GetLeaderboardByFilter(filter models.Filter) (models.Leaderboard, error)
}

func NewFilterService(repo filter.Filter) *FilterService {
	return &FilterService{
		repo: repo,
	}
}
