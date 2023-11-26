package aetherAPI

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Data struct {
	ComputerName string `json:"computerName"`
	ScreenNumber string `json:"screenNumber"`
	DateTime     string `json:"dateTime"`
	Base64Value  string `json:"base64Value"`
}

func storeToRedis(screenshot string) {
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

	// Get the current time
	currentTime := time.Now()
	// Format the current time as a string in a specific layout
	timeString := currentTime.Format("2006-01-02 15:04:05") // YYYY-MM-DD HH:MM:SS

	// Create a Data struct instance
	data := Data{
		ComputerName: "aether",
		ScreenNumber: "0",
		DateTime:     timeString,
		Base64Value:  encodeBase64(screenshot),
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data to JSON:", err)
		return
	}

	// Store the JSON data in a Redis List
	err = rdb.LPush(context.Background(), "aether", string(jsonData)).Err()
	if err != nil {
		fmt.Println("Error pushing JSON data to Redis list:", err)
		return
	}
	fmt.Println("Stored latest JSON data in Redis list 'latestData'")

	// // Trimming the list to keep only the latest 10 items
	// err = rdb.LTrim(context.Background(), "latestData", 0, 9).Err()
	// if err != nil {
	// 	fmt.Println("Error trimming Redis list:", err)
	// 	return
	// }
}

func encodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
