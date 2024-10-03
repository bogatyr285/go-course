package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func resetDB(rdb *redis.Client) {
	rdb.FlushDB(ctx)
}

func printAllKeys(rdb *redis.Client) {
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current keys in the Redis DB:")
	for _, key := range keys {
		val, _ := rdb.Get(ctx, key).Result()
		ttl, _ := rdb.TTL(ctx, key).Result()
		fmt.Printf("%s: (length: %d, ttl: %v)\n", key, len(val), ttl)
	}
	fmt.Println()
}

func setupDB(rdb *redis.Client, n int, withTTL bool) {
	largeValue := strings.Repeat("X", 150000)
	for i := 0; i < n; i++ {
		if withTTL {
			err := rdb.Set(ctx, fmt.Sprintf("init%d", i), largeValue, 10*time.Minute).Err()
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			rdb.Set(ctx, fmt.Sprintf("init%d", i), largeValue, 0)
		}
	}
}

func triggerEviction(rdb *redis.Client, withTTL bool) {
	largeValue := strings.Repeat("Y", 150000) //  150KB per value
	for i := 20; i < 40; i++ {
		if withTTL {
			rdb.Set(ctx, fmt.Sprintf("evictTTL%d", i), largeValue, 10*time.Minute)
		} else {
			rdb.Set(ctx, fmt.Sprintf("evict%d", i), largeValue, 0)
		}
	}
}

func demonstratePolicy(policy, memory string, withTTL bool) {
	fmt.Printf("Demonstrating policy: %s with maxmemory of %s\n", policy, memory)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := rdb.ConfigSet(ctx, "maxmemory", memory).Err()
	if err != nil {
		log.Fatalf("Could not set maxmemory: %v", err)
	}
	err = rdb.ConfigSet(ctx, "maxmemory-policy", policy).Err()
	if err != nil {
		log.Fatalf("Could not set maxmemory-policy: %v", err)
	}

	resetDB(rdb)
	setupDB(rdb, 20, withTTL)

	fmt.Println("Initial state of the DB:")
	printAllKeys(rdb)

	triggerEviction(rdb, withTTL)

	fmt.Printf("State of the DB after triggering eviction with %s policy:\n", policy)
	printAllKeys(rdb)

	resetDB(rdb)
	rdb.Close()
}

func main() {
	memory := "2mb"

	evictionPolicies := []struct {
		policy  string
		withTTL bool
	}{
		{"noeviction", false},
		{"allkeys-lru", false},
		{"volatile-lru", true},
		{"allkeys-random", false},
		{"volatile-random", true},
	}

	for _, ep := range evictionPolicies {
		demonstratePolicy(ep.policy, memory, ep.withTTL)
		time.Sleep(2 * time.Second)
	}
}
