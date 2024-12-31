package main

import (
	"fmt"
	"time"

	"github.com/coles243/red/red"
	"github.com/redis/go-redis/v9"
)

// Define sample user details
var Name string = "John Doe"
var Email string = "JohnDoe@gmail.com"

func main() {
	// Create a connection with Redis
	db := red.RedisDB{
		DB: redis.NewClient(
			&redis.Options{},
		),
	}

	// Ensure the Redis connection is closed when done
	defer db.DB.Close()

	// Store the user details in Redis with a key and expiration time
	response, err := db.CreateSet(Name, Email, time.Duration(10)*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	// Print the response status
	fmt.Println(response)
}
