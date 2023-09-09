package context

import (
	. "github.com/SoorajKothari/Messaging_Pub/pkg/messaging"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

type Context struct {
	Client   *redis.Client
	redisUrl string
	Name     string
}

func GetContext() (*Context, error) {
	context := &Context{
		Name: "Publisher",
	}
	context.UpdateEnvVariables()
	client, err := InitializeRedis(context.redisUrl)
	if err != nil {
		log.Println("Fail to Initialize redis client", err)
		return nil, err
	}
	context.Client = client
	log.Println("Client Initialized Successfully")
	return context, nil
}

func (context *Context) UpdateEnvVariables() {
	if val, found := os.LookupEnv("REDIS_HOST"); found {
		context.redisUrl = val
		log.Println("Environment Variable Resolved")
	}
}
