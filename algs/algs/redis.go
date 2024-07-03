package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type UserScore struct {
	UserID string
	Score  float64
}

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

// LoadCSV loads user scores from a CSV file
func LoadCSV(filePath string) ([]UserScore, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var userScores []UserScore

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		score, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}

		userScores = append(userScores, UserScore{
			UserID: record[0],
			Score:  score,
		})
	}

	return userScores, nil
}

// AddScoresToLeaderboard adds user scores to the Redis sorted set
func (r RedisService) AddScoresToLeaderboard(ctx context.Context, userScores []UserScore) error {
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
func (r RedisService) GetTopNUsers(ctx context.Context, n int) ([]UserScore, error) {
	zs, err := r.rdb.ZRevRangeWithScores(ctx, r.ZSETKey, 0, int64(n-1)).Result()
	if err != nil {
		return nil, err
	}
	var userScores []UserScore
	for _, z := range zs {
		userScores = append(userScores, UserScore{
			UserID: z.Member.(string),
			Score:  z.Score,
		})
	}
	return userScores, nil
}

// docker run --rm -it -p 6379:6379 redis:7.2.5-alpine
// ZRANGE leaderboard 0 -1 withscores
func main() {
	// Initialize Redis client
	rs, err := InitializeRedisClient(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Load user scores from a CSV file
	filePath := "scores.csv" // Replace with the path to your CSV file
	userScores, err := LoadCSV(filePath)
	if err != nil {
		log.Fatalf("Error loading CSV: %v", err)
	}

	// Add scores to leaderboard
	err = rs.AddScoresToLeaderboard(context.Background(), userScores)
	if err != nil {
		log.Fatalf("Error adding scores to leaderboard: %v", err)
	}

	// Get top N users
	topN := 10
	topUsers, err := rs.GetTopNUsers(context.Background(), topN)
	if err != nil {
		log.Fatalf("Error getting top users: %v", err)
	}

	// Output top users
	fmt.Printf("Top %d Users:\n", topN)
	for i, user := range topUsers {
		fmt.Printf("%d. UserID: %s, Score: %.2f\n", i+1, user.UserID, user.Score)
	}
}
