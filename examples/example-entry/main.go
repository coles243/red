package main

import (
	"fmt"
	"time"

	"github.com/coles243/red/red"
	"github.com/redis/go-redis/v9"
)

func main() {
	var Name string = "John Doe"
	var Email string = "JohnDoe@gmail.com"

	db := red.RedisDB{
		DB: redis.NewClient(
			&redis.Options{},
		),
	}

	defer db.DB.Close()

	response, err := db.CreateSet(Name, Email, time.Duration(time.Duration(10).Seconds()))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)

}
