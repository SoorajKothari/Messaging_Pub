package context

import (
	. "github.com/SoorajKothari/Messaging_Pub/pkg/messaging"
	"github.com/go-redis/redis/v8"
	"log"
)

type Context struct {
	Client  *redis.Client
	Version string
}

func GetContext() (*Context, error) {
	context := &Context{
		Version: "1",
	}
	client, err := InitializeRedis("localhost:6379")
	if err != nil {
		log.Println("Fail to Initialize redis client", err)
		return nil, err
	}
	context.Client = client
	log.Println("Client Initialized Successfully")
	return context, nil
}
