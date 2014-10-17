package main

import (
	"log"
	"net/http"
	"utils"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("start server")

	err := InitServer()
	if err != nil {
		log.Fatalf("Init Server fail [%s]\n", err.Error())
	}
	http.HandleFunc("/v1/protobuf", CallbackJesgoo)
	http.HandleFunc("/v1/json", CallbackJesgooJson)
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		utils.FatalLog.Write("start server fail . err[%s]", err.Error())
	}
}
