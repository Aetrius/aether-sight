package aetherRedisConsumer

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func startConsumer() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.2.4:6379", // Redis server address
	})

	// Ping the Redis server to check the connection
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// Fetch all keys
	keys, err := rdb.Keys(context.Background(), "*").Result()
	if err != nil {
		fmt.Println("Error fetching keys:", err)
		return
	}

	// Iterate through all keys
	for _, key := range keys {
		// Fetch the length of the list
		length, err := rdb.LLen(context.Background(), key).Result()
		if err != nil {
			fmt.Println("Error fetching length of key:", err)
			continue
		}

		// Delete items if length > 10
		if length > 10 {
			itemsToDelete := length - 10
			err := rdb.LTrim(context.Background(), key, itemsToDelete, -1).Err()
			if err != nil {
				fmt.Println("Error trimming key:", key, err)
			} else {
				fmt.Printf("Deleted %d items from key '%s'\n", itemsToDelete, key)
			}
		} else {
			fmt.Printf("Key '%s' has fewer than 10 items, nothing to delete\n", key)
		}
	}
}
