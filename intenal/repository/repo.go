package repository

import (
	"database/sql"
	"dota2scheduler/intenal/repository/fill"
	"dota2scheduler/intenal/repository/filter"
	"dota2scheduler/intenal/repository/server"
)

type Repo struct {
	fill.Fill
	filter.Filter
	server.Server
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		Server: server.NewServerRepo(db),
		Fill:   fill.NewFillRepo(db),
		Filter: filter.NewFilterRepo(db),
	}
}
