package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	. "pub.go/pkg/model"
	"pub.go/pkg/service"
	"pub.go/pkg/store"
)

var upGrader = websocket.Upgrader{}

func HandleConnection(writer http.ResponseWriter, request *http.Request) {
	connection, err := upGrader.Upgrade(writer, request, nil)

	if err != nil {
		log.Println("Error upgrading connection to websocket ", err)
		return
	}
	id, err := store.Add(connection)
	if err != nil {
		log.Println("Error storing connection", err)
		return
	}
	for {
		var message Message
		err := connection.ReadJSON(&message)
		if err != nil {
			log.Println("Error reading message from websocket", err)
			store.DeleteById(id)
		}
		message.SessionId = id
		err = service.Publish(globalContext, &message)
		if err != nil {
			log.Println("Error processing message", err)
			return
		}
	}
}