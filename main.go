package main

import (
	. "github.com/SoorajKothari/Messaging_Pub/pkg/context"
	"github.com/SoorajKothari/Messaging_Pub/pkg/server"
	"log"
)

func main() {
	log.Println("Starting App")
	context, err := GetContext()
	if err != nil {
		log.Println("Fail to start app", err)
		return
	}
	log.Println("App Started", context.Version)
	server.Start(context)
}
