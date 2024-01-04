package models

type Server struct {
	TimePosted            int `json:"time_poster"`
	NextScheduledPostTime int `json:"next_scheduled_post_time"`
	ServerTime            int `json:"server_time"`
}
