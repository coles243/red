package red_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/coles243/red/red"
	"github.com/redis/go-redis/v9"
)

func TestFetch(t *testing.T) {
	var expect int

	var actual int = 10

	db := red.RedisDB{
		DB: redis.NewClient(
			&redis.Options{},
		),
	}

	_, err := db.CreateSet("Numeric", 10, time.Duration(10)*time.Second)
	if err != nil {
		t.Error(err.Error())

	}

	data, err := db.FetchValue("Numeric")
	if err != nil {
		t.Error(err.Error())

	}

	if v, ok := data.(string); ok {

		e, err := strconv.Atoi(v)

		if err != nil {
			t.Error(err.Error())

		}

		expect = e
	}

	if expect != actual {
		t.Errorf("Expected: %v  but actual: %v\n", expect, actual)
	}

}
