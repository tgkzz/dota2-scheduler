package models

type Leaderboard struct {
	Players []Player `json:"leaderboard"`
}

type Player struct {
	Rank    int    `json:"rank"`
	Name    string `json:"name"`
	TeamID  int    `json:"team_id,omitempty"`
	TeamTag string `json:"team_tag,omitempty"`
	Country string `json:"country,omitempty"`
	Sponsor string `json:"sponsor,omitempty"`
}
