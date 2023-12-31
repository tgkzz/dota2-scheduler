package models

type Leaderboard struct {
	TimePosted            int      `json:"time_poster"`
	NextScheduledPostTime int      `json:"next_scheduled_post_time"`
	ServerTime            int      `json:"server_time"`
	Players               []Player `json:"leaderboard"`
}

type Player struct {
	Rank    int    `json:"rank"`
	Name    string `json:"name"`
	TeamID  int    `json:"team_id,omitempty"`
	TeamTag string `json:"team_tag,omitempty"`
	Country string `json:"country,omitempty"`
	Sponsor string `json:"sponsor,omitempty"`
}
