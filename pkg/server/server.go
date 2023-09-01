package server

import (
	. "github.com/SoorajKothari/Messaging_Pub/pkg/context"
	"log"
	"net/http"
)

var globalContext *Context

func Start(context *Context) {
	globalContext = context
	http.HandleFunc("/v1/send", HandleConnection)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
}
