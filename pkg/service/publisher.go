package service

import (
	"context"
	"log"
	. "pub.go/pkg/context"
	. "pub.go/pkg/model"
)

func Publish(ctx *Context, msg Message) error {
	cmd, err := ctx.Client.Publish(context.Background(), "main", msg.Content).Result()
	if err != nil {
		log.Println("Error publishing message to redis", err)
		return err
	}
	log.Println("Published Message", cmd)
	return nil
}
