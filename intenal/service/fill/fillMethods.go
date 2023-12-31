package fill

import (
	"dota2scheduler/intenal/models"
)

func (f *FillService) UpdateLeaderboard(leaderboard models.Leaderboard) error {
	if err := f.repo.DeleteLeaderboard(); err != nil {
		return err
	}

	if err := f.repo.CreateLeaderboard(leaderboard); err != nil {
		return err
	}

	return nil
}

func (f *FillService) ShowLeaderboard() (models.Leaderboard, error) {
	return f.repo.ReadLeaderboard()
}

func (f *FillService) MakeRequest(regions []string) ([]models.Leaderboard, error) {
	return f.repo.RequestDotaLeaderboard(regions)
}
