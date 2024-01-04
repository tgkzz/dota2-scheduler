package server

import "dota2scheduler/intenal/models"

func (s *ServerRepo) CreateServer(server models.Server) error {
	query := "INSERT INTO Server (TimePosted, NextPostTime, ServerTime) VALUES (?, ?, ?)"

	_, err := s.DB.Exec(query, server.TimePosted, server.NextScheduledPostTime, server.ServerTime)

	return err
}

func (s *ServerRepo) ReadServer() (models.Server, error) {
	query := "SELECT TimePosted, ServerTime, NextPostTime FROM Server"

	var server models.Server

	_, err := s.DB.Exec(query, &server.TimePosted, &server.ServerTime, &server.NextScheduledPostTime)

	return server, err

}

func (s *ServerRepo) DeleteServer() error {
	query := "DELETE FROM Server"

	_, err := s.DB.Exec(query)

	return err
}
