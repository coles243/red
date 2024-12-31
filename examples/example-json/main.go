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
	// var UnserializedUser Person
	UserJson := Person{
		Name: "John Doe",
		Age:  30,
	}

	// create connection with redisDb
	db := red.RedisDB{
		DB: redis.NewClient(
			&redis.Options{},
		),
	}

	//serialize to json
	data, err := json.Marshal(UserJson)
	if err != nil {
		fmt.Println(err.Error())
	}

	response, err := db.CreateSet("User1", data, time.Duration(time.Duration(10).Seconds()))

	if err != nil {
		fmt.Println(err.Error())
	}

	// return string status
	fmt.Println(response)

	// deserialize to a go type
	Getdata, err := db.SetReteriver("User1")
	if err != nil {
		fmt.Println(err.Error())
	}

	if v, ok := Getdata.(string); ok {
		err = json.Unmarshal([]byte(v), &ReturnedUser)
		if err != nil {
			fmt.Println(err.Error())

		}
	}
	// Return Go Object
	fmt.Printf("Hi my name is %v and age is %v\n", ReturnedUser.Name, ReturnedUser.Age)

}
