package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/coles243/red/red"
	"github.com/redis/go-redis/v9"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var ReturnedUser Person

	// Create a sample user
	UserJson := Person{
		Name: "John Doe",
		Age:  30,
	}

	// Create a connection with Redis
	db := red.RedisDB{
		DB: redis.NewClient(
			&redis.Options{},
		),
	}

	// Ensure the Redis connection is closed when done
	defer db.DB.Close()

	// Serialize the user to JSON
	data, err := json.Marshal(UserJson)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Store the JSON data in Redis with a key and expiration time
	response, err := db.CreateSet("User1", data, time.Duration(10)*time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Print the response status
	fmt.Println(response)

	// Retrieve the JSON data from Redis
	Getdata, err := db.FetchValue("User1")
	if err != nil {
		fmt.Println(err.Error())
	}

	// Deserialize the JSON data back to a Go struct
	// Values returned are always interface object, Validate Below
	if v, ok := Getdata.(string); ok {
		err = json.Unmarshal([]byte(v), &ReturnedUser)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	// Print the deserialized Go object
	fmt.Printf("Hi, my name is %v and my age is %v\n", ReturnedUser.Name, ReturnedUser.Age)
}
