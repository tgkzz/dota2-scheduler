package filter

import "dota2scheduler/intenal/models"

func (f *FilterService) GetLeaderboardByFilter(filter models.Filter) (models.Leaderboard, error) {
	return f.repo.GetLeaderboardByFilter(filter)
}
