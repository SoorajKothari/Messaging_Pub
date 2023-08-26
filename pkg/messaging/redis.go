package messaging

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func InitializeRedis(address string) (*redis.Client, error) {
	log.Println("Initializing Redis Client.")
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})
	err := client.Ping(context.Background()).Err()
	return client, err
}
