package main

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		Protocol: 2,
	})
	log.Println("Connected to Redis")

	ctx:= context.Background()

	err:= client.Set(ctx,"foo","bar-bar",time.Duration(2 * time.Minute)).Err()
	if err != nil {
		panic(err)
	}
	log.Println("Set key foo with value bar")
}