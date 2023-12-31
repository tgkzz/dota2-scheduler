package fill

import (
	"dota2scheduler/intenal/models"
	"dota2scheduler/intenal/repository/fill"
)

type FillService struct {
	repo fill.Fill
}

type Filler interface {
	MakeRequest(regions []string) ([]models.Leaderboard, error)
	UpdateLeaderboard(leaderboard models.Leaderboard) error
	ShowLeaderboard() (models.Leaderboard, error)
}

func NewFillService(repo fill.Fill) *FillService {
	return &FillService{
		repo: fill.Fill(repo),
	}
}
