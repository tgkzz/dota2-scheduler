package fill

import (
	"dota2scheduler/intenal/models"
	"dota2scheduler/intenal/repository/fill"
	"github.com/redis/go-redis/v9"
)

type FillService struct {
	repo  fill.Fill
	redis *redis.Client
}

type Filler interface {
	MakeRequest(regions []string) ([]models.Leaderboard, error)
	UpdateLeaderboard(leaderboard models.Leaderboard) error
	ShowLeaderboard() (models.Leaderboard, error)
	UpdateRedis(leaderboard models.Leaderboard) error
	GetRedisData() (map[string]string, error)
}

func NewFillService(repo fill.Fill, client *redis.Client) *FillService {
	return &FillService{
		repo:  fill.Fill(repo),
		redis: client,
	}
}
