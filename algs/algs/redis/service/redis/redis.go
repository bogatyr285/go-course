package redis

import (
	"context"
	"fmt"

	"github.com/username/myAwesomeProject/models"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	ZSETKey string
	rdb     *redis.Client
}

func InitializeRedisClient(ctx context.Context) (RedisService, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return RedisService{}, fmt.Errorf("Could not connect to Redis: %v", err)
	}
	return RedisService{rdb: rdb}, nil
}

// AddScoresToLeaderboard adds user scores to the Redis sorted set
func (r RedisService) AddScoresToLeaderboard(ctx context.Context, userScores []models.UserScore) error {
	var redisZs []redis.Z
	for _, userScore := range userScores {
		redisZs = append(redisZs, redis.Z{
			Score:  userScore.Score,
			Member: userScore.UserID,
		})
	}

	err := r.rdb.ZAdd(ctx, r.ZSETKey, redisZs...).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetTopNUsers returns the top N users from the leaderboard
func (r RedisService) GetTopNUsers(ctx context.Context, n int) ([]models.UserScore, error) {
	zs, err := r.rdb.ZRevRangeWithScores(ctx, r.ZSETKey, 0, int64(n-1)).Result()
	if err != nil {
		return nil, err
	}
	var userScores []models.UserScore
	for _, z := range zs {
		userScores = append(userScores, models.UserScore{
			UserID: z.Member.(string),
			Score:  z.Score,
		})
	}
	return userScores, nil
}
