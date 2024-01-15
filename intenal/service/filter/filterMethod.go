package filter

import "dota2scheduler/intenal/models"

func (f *FilterService) GetLeaderboardByFilter(filter models.Filter) (models.Leaderboard, error) {
	var emptyFilter models.Filter

	if filter == emptyFilter {
		return models.Leaderboard{}, models.ErrEmptyFilter
	}

	return f.repo.GetLeaderboardByFilter(filter)
}
