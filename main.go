package main

import (
	"log"
	"net/http"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("start server")

	InitServer()
	http.HandleFunc("/v1/protobuf", CallbackJesgoo)
	http.HandleFunc("/v1/json", CallbackJesgooJson)
	http.ListenAndServe(":8081", nil)
}
