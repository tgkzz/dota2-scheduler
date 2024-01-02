package models

type Filter struct {
	FromRank int    `json:"from_rank,omitempty"`
	ToRank   int    `json:"to_Rank,omitempty"`
	Name     string `json:"name,omitempty"`
	TeamTag  string `json:"team_tag,omitempty"`
	Country  string `json:"country,omitempty"`
	Sponsor  string `json:"sponsor,omitempty"`
}
