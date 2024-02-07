package fill

import (
	"context"
	"dota2scheduler/intenal/models"
	"fmt"
)

// UpdateLeaderboard TODO: update method in order to load another regions
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

func (f *FillService) UpdateRedis(leaderboard models.Leaderboard) error {
	ctx := context.Background()

	if err := f.redis.Del(ctx, "leaderboard").Err(); err != nil {
		fmt.Println("asd")
		return err
	}

	playerMap := map[string]interface{}{}
	for _, player := range leaderboard.Players {
		playerMap[fmt.Sprintf("%d", player.Rank)] = player.Name
	}

	if err := f.redis.HSet(ctx, "leaderboard", playerMap).Err(); err != nil {
		fmt.Println("qwe")
		return err
	}

	return nil
}

func (f *FillService) GetRedisData() (map[string]string, error) {
	result := map[string]string{}

	ctx := context.Background()

	for i := 1; i < 5000; i++ {
		rank := fmt.Sprintf("%d", i)
		name, err := f.redis.Get(ctx, rank).Result()
		if err != nil {
			return nil, err
		}

		result[rank] = name
	}

	return result, nil
}
