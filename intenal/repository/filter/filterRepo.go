package filter

import (
	"database/sql"
	"dota2scheduler/intenal/models"
)

type FilterRepo struct {
	DB *sql.DB
}

type Filter interface {
	GetLeaderboardByFilter(fitler models.Filter) (models.Leaderboard, error)
}

func NewFilterRepo(db *sql.DB) *FilterRepo {
	return &FilterRepo{DB: db}
}
