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
		Protocol: 2,
	})
	slog.Info("Connected to Redis")

	ctx := context.Background()

	hashFields := []string{
		"model", "Deimos",
		"brand", "Ergonom",
		"type", "Enduro bikes",
		"price", "4972",
	}

	hash, err := client.HSet(ctx, "bike", hashFields).Result()
	if err != nil {
		panic(err)
	}

	println(hash)
	slog.Info("Hash set", "hash", hash)

}
