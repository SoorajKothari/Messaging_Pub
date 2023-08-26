package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "pub.go/pkg/context"
)

var globalContext *Context

func Start(context *Context) {
	globalContext = context
	router := mux.NewRouter()
	router.HandleFunc("/v1/send", ProcessMessage).Methods("POST")
	err := http.ListenAndServe(":9000", router)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
}
