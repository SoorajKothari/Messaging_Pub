package main

import (
	. "github.com/SoorajKothari/Messaging_Pub/pkg/context"
	"github.com/SoorajKothari/Messaging_Pub/pkg/server"
	"log"
)

func main() {
	log.Println("Starting App")
	context, err := GetContext()
	go server.Reply()
	if err != nil {
		log.Println("Fail to start app", err)
		return
	}
	log.Println("App Started : ", context.Name)
	server.Start(context)
}
