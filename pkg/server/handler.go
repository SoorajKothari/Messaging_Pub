package server

import (
	"encoding/json"
	"log"
	"net/http"
	"pub.go/pkg/model"
	"pub.go/pkg/service"
)

func ProcessMessage(writer http.ResponseWriter, request *http.Request) {
	newMessage := model.Message{}
	err := json.NewDecoder(request.Body).Decode(&newMessage)
	if err != nil {
		log.Println("Error while decoding request body", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = service.Publish(globalContext, newMessage)
	if err != nil {
		log.Println("Error procesing message", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
