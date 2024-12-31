package red

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisDB struct {
	DB *redis.Client
}

// Create a Key,Value
func (r *RedisDB) CreateSet(key string, value interface{}, exp time.Duration) (string, error) {
	_, err := r.DB.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}

	err = r.DB.Set(ctx, key, value, exp).Err()
	if err != nil {
		redErr := fmt.Sprintln(err.Error())
		return "", errors.New("Unable to establish New Record: " + redErr)
	}
	return "Record created successfully", nil
}

// retrieve set based on key input
func (r *RedisDB) FetchValue(key string) (interface{}, error) {

	var data interface{}
	_, err := r.DB.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}

	response, err := r.DB.Get(ctx, key).Result()
	if err != nil {
		redErr := fmt.Sprintln(err.Error())
		data = response
		return data, errors.New("There was a problem retrieving from redis: " + redErr)
	}
	data = response
	return data, nil

}

// Delete from database
func (r *RedisDB) Delete(key string) (int64, error) {

	_, err := r.DB.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}

	response, err := r.DB.Del(ctx, key).Result()
	if err != nil {
		redErr := fmt.Sprintln(err)
		return 1, errors.New("There was a problem deleting from redis:  " + redErr)
	}

	return response, nil
}
