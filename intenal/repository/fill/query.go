package fill

import (
	"dota2scheduler/intenal/models"
	"encoding/json"
	"net/http"
)

// TODO: make methods for querying database
func (f *FillRepo) CreateLeaderboard(leaderboard models.Leaderboard) error {
	for _, player := range leaderboard.Players {
		query := "INSERT INTO Players (Rank, Name, TeamID, TeamTag, Country, Sponsor) VALUES (?, ?, ?, ?, ?, ?)"
		if _, err := f.DB.Exec(query, player.Rank, player.Name, player.TeamID, player.TeamTag, player.Country, player.Sponsor); err != nil {
			return err
		}

	}

	return nil
}

func (f *FillRepo) ReadLeaderboard() (models.Leaderboard, error) {
	var Result models.Leaderboard

	query := "SELECT Rank, Name, TeamID, TeamTag, Country, Sponsor FROM Players"
	rows, err := f.DB.Query(query)
	if err != nil {
		return Result, err
	}

	for rows.Next() {
		var player models.Player
		if err := rows.Scan(&player.Rank, &player.Name, &player.TeamID, &player.TeamTag, &player.Country, &player.Sponsor); err != nil {
			return models.Leaderboard{}, err
		}
		Result.Players = append(Result.Players, player)
	}

	if err := rows.Err(); err != nil {
		return Result, err
	}

	return Result, err
}

func (f *FillRepo) DeleteLeaderboard() error {
	query := "DELETE FROM Players"

	_, err := f.DB.Exec(query)

	return err
}

func (f *FillRepo) RequestDotaLeaderboard(regions []string) ([]models.Leaderboard, error) {
	result := []models.Leaderboard{}

	for _, region := range regions {
		url := models.FormURL(region)

		resp, err := http.Get(url)
		if err != nil {
			return []models.Leaderboard{}, err
		}

		var leaderboard models.Leaderboard
		if err = json.NewDecoder(resp.Body).Decode(&leaderboard); err != nil {
			return []models.Leaderboard{}, err
		}
		//for _, player := range leaderboard.Players {
		//	log.Printf("Rank: %d, Name: %s\n", player.Rank, player.Name)
		//}

		result = append(result, leaderboard)

	}

	return result, nil
}
