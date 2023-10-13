package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.simple.redis/person"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	mockData := person.Mock()
	for k, v := range mockData {
		bytes, _ := json.Marshal(v)
		err := rdb.HSet(ctx, "redis:person:v1", k, bytes).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
}
