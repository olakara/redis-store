package main

import (
	"context"
	"log"

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

	value, err:= client.Get(ctx,"foo").Result()
	if err != nil {
		panic(err)
	}
	log.Println("Value of foo is: ", value)
}