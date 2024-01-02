package filter

import "dota2scheduler/intenal/models"

func (f *FilterRepo) GetLeaderboardByFilter(filter models.Filter) (models.Leaderboard, error) {
	var leaderboard models.Leaderboard

	query := "SELECT Rank, Name, TeamId, TeamTag, Country, Sponsor FROM Players WHERE 1=1"
	params := []interface{}{}

	if filter.FromRank != 0 {
		query += " AND Rank >= ?"
		params = append(params, filter.FromRank)
	}
	if filter.ToRank != 0 {
		query += " AND Rank <= ?"
		params = append(params, filter.ToRank)
	}
	if filter.Name != "" {
		query += " AND Name = ?"
		params = append(params, filter.Name)
	}
	if filter.TeamTag != "" {
		query += " AND TeamTag = ?"
		params = append(params, filter.TeamTag)
	}
	if filter.Sponsor != "" {
		query += " AND Sponsor = ?"
		params = append(params, filter.Sponsor)
	}
	if filter.Country != "" {
		query += " AND Country LIKE ?"
		params = append(params, filter.Country)
	}

	rows, err := f.DB.Query(query, params...)
	if err != nil {
		return leaderboard, err
	}
	defer rows.Close()

	for rows.Next() {
		var player models.Player
		err := rows.Scan(&player.Rank, &player.Name, &player.TeamID, &player.TeamTag, &player.Country, &player.Sponsor)
		if err != nil {
			return leaderboard, err
		}
		leaderboard.Players = append(leaderboard.Players, player)
	}

	return leaderboard, nil
}
