package server

import "dota2scheduler/intenal/models"

func (s *ServerService) GetServerInfo() (models.Server, error) {
	return s.repo.ReadServer()
}

func (s *ServerService) RequestServerInfo() (models.Server, error) {
	return s.repo.GetServerInfo()
}

func (s *ServerService) UpdateServerInfo(server models.Server) error {
	if err := s.repo.DeleteServer(); err != nil {
		return err
	}

	if err := s.repo.CreateServer(server); err != nil {
		return err
	}

	return nil
}
