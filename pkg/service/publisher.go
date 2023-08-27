package service

import (
	"context"
	"encoding/json"
	"log"
	. "pub.go/pkg/context"
	. "pub.go/pkg/model"
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
