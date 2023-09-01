package server

import (
	"context"
	"encoding/json"
	. "github.com/SoorajKothari/Messaging_Pub/pkg/model"
	"github.com/SoorajKothari/Messaging_Pub/pkg/service"
	"github.com/SoorajKothari/Messaging_Pub/pkg/store"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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

func Reply() {
	pubsub := globalContext.Client.Subscribe(context.Background(), "reply")
	channel := pubsub.Channel()
	for msg := range channel {
		var message Message
		err := json.Unmarshal([]byte(msg.Payload), &message)
		if err != nil {
			log.Println("Invalid message")
		}
		log.Println("Sending reply.")
		conn := store.GetConnectionById(message.SessionId)
		err = conn.WriteJSON(message.Content)
		if err != nil {
			log.Println("Error sending reply.")
			return
		}
	}
}
