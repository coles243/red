package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/coles243/red/red"
	"github.com/redis/go-redis/v9"
)

func main() {

	test := red.RedisDB{
		DB: redis.NewClient(&redis.Options{}),
	}

	tester := map[string]int{
		"red":  3,
		"blue": 4,
	}

	for key, value := range tester {
		data, err := test.CreateSet(key, value, time.Duration(time.Duration(10).Seconds()))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)
	}

	getdata, err := test.SetReteriver("blue")

	if err != nil {
		fmt.Println(err)
	}

	//
	v, ok := getdata.(string)
	if ok {
		data, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(data + 1)
	}

	//Redis Easy Driver
}
