package main

import (
	"context"
	"fmt"
	"log"

	"github.com/username/myAwesomeProject/service/csv"
	redisService "github.com/username/myAwesomeProject/service/redis"
)

// docker run --rm -it -p 6379:6379 redis:7.2.5-alpine
// ZRANGE leaderboard 0 -1 withscores
func main() {
	// Initialize Redis client
	rs, err := redisService.InitializeRedisClient(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Load user scores from a CSV file
	filePath := "scores.csv" // Replace with the path to your CSV file
	userScores, err := csv.LoadCSV(filePath)
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
