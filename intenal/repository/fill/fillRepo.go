package fill

import (
	"database/sql"
	"dota2scheduler/intenal/models"
)

type FillRepo struct {
	DB *sql.DB
}

type Fill interface {
	CreateLeaderboard(leaderboard models.Leaderboard) error
	ReadLeaderboard() (models.Leaderboard, error)
	DeleteLeaderboard() error
	RequestDotaLeaderboard(regions []string) ([]models.Leaderboard, error)
}

func NewFillRepo(db *sql.DB) *FillRepo {
	return &FillRepo{DB: db}
}
