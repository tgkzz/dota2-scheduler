package repository

import (
	"database/sql"
	"dota2scheduler/intenal/repository/fill"
)

type Repo struct {
	fill.Fill
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		Fill: fill.NewFillRepo(db),
	}
}
