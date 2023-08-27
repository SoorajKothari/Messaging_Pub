package store

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
)

var store = make(map[string]*websocket.Conn)

func Add(connection *websocket.Conn) (string, error) {
	connectionId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := connectionId.String()
	store[id] = connection
	log.Println("Connection Store: ", len(store))
	return id, nil
}

func GetConnectionById(connectionId string) *websocket.Conn {
	return store[connectionId]
}

func DeleteById(connectionId string) {
	delete(store, connectionId)
}
