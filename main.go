package main

import (
	"log"
	. "pub.go/pkg/context"
	"pub.go/pkg/server"
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
