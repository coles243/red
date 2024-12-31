package red

import (
	"context"
	"errors"
	"fmt"
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
		return "", fmt.Errorf("error Establishing connection: %v", err)
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
		return "", fmt.Errorf("error Establishing connection: %v", err)
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
		return 1, fmt.Errorf("error Establishing connection: %v", err)
	}

	response, err := r.DB.Del(ctx, key).Result()
	if err != nil {
		redErr := fmt.Sprintln(err)
		return 1, errors.New("There was a problem deleting from redis:  " + redErr)
	}

	return response, nil
}

// Updating from Database
// Update an existing key's value in Redis
func (r *RedisDB) UpdateValue(key string, newValue interface{}, exp time.Duration) (string, error) {

	_, err := r.DB.Ping(ctx).Result()
	if err != nil {
		return "", fmt.Errorf("error Establishing connection: %v", err)
	}

	// Check if the key exists
	exists, err := r.DB.Exists(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("error checking if key exists: %v", err)
	}
	if exists == 0 {
		return "", errors.New("key does not exist")
	}

	// Update the key's value
	err = r.DB.Set(ctx, key, newValue, exp).Err()
	if err != nil {
		return "", fmt.Errorf("error updating key: %v", err)
	}

	return "Value updated successfully", nil
}
