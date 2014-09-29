package main

import (
	"log"
	"net/http"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("start server")

	InitServer()
	http.HandleFunc("/", CallbackJesgoo)
	http.ListenAndServe(":8081", nil)
}
