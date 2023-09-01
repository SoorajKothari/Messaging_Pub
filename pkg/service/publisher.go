package service

import (
	"context"
	"encoding/json"
	. "github.com/SoorajKothari/Messaging_Pub/pkg/context"
	. "github.com/SoorajKothari/Messaging_Pub/pkg/model"
	"log"
)

func Publish(ctx *Context, msg *Message) error {
	jsonMessage, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	cmd, err := ctx.Client.Publish(context.Background(), "main", jsonMessage).Result()
	if err != nil {
		log.Println("Error publishing message to redis", err)
		return err
	}
	log.Println("Published Message", cmd)
	return nil
}
