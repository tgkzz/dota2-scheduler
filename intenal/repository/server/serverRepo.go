package server

import (
	"database/sql"
	"dota2scheduler/intenal/models"
)

type ServerRepo struct {
	DB *sql.DB
}

type Server interface {
	CreateServer(server models.Server) error
	DeleteServer() error
}

func NewServerRepo(db *sql.DB) *ServerRepo {
	return &ServerRepo{DB: db}
}
