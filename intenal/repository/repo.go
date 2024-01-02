package repository

import (
	"database/sql"
	"dota2scheduler/intenal/repository/fill"
	"dota2scheduler/intenal/repository/filter"
)

type Repo struct {
	fill.Fill
	filter.Filter
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		Fill:   fill.NewFillRepo(db),
		Filter: filter.NewFilterRepo(db),
	}
}
