package main

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func main() {
	// connect to redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	slog.Info("Connected to redis")

	ctx := context.Background()
	// get value
	val, err := client.HGetAll(ctx, "bike").Result()
	if err != nil {
		panic(err)
	}
	slog.Info("Model: ", val)
}
