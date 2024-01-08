package server

import (
	"dota2scheduler/intenal/models"
	"encoding/json"
	"net/http"
)

func (s *ServerRepo) CreateServer(server models.Server) error {
	query := "INSERT INTO Server (TimePosted, NextPostTime, ServerTime) VALUES (?, ?, ?)"

	_, err := s.DB.Exec(query, server.TimePosted, server.NextScheduledPostTime, server.ServerTime)

	return err
}

func (s *ServerRepo) GetServerInfo() (models.Server, error) {
	url := models.FormURL("europe")

	resp, err := http.Get(url)
	if err != nil {
		return models.Server{}, err
	}

	var result models.Server
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.Server{}, err
	}

	return result, err
}

func (s *ServerRepo) ReadServer() (models.Server, error) {
	query := "SELECT TimePosted, ServerTime, NextPostTime FROM Server"

	var server models.Server

	err := s.DB.QueryRow(query).Scan(&server.TimePosted, &server.ServerTime, &server.NextScheduledPostTime)

	return server, err
}

func (s *ServerRepo) DeleteServer() error {
	query := "DELETE FROM Server"

	_, err := s.DB.Exec(query)

	return err
}
