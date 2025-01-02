package main

import (
	"fmt"
	"time"

	"github.com/coles243/red/red"
	"github.com/redis/go-redis/v9"
)

func main() {
	test := red.RedisDB{
		DB: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}

	// Ensure the Redis connection is closed when done
	defer test.DB.Close()

	test.CreateSet("John", "Doe", time.Duration(time.Duration(10).Seconds()))
	// Check current state of value
	fmt.Print(test.FetchValue("John"))

	update, err := test.UpdateValue("John", "Matters", time.Duration(time.Duration(10).Seconds()))
	if err != nil {
		fmt.Println(err)
	}
	//validation
	fmt.Println(update)
	//Check New Value Update
	fmt.Println(test.FetchValue("John"))

}
