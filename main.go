package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"utils"
)

//import _ "net/http/pprof"

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("start server")

	var global_context context.GlobalContext
	err := InitServer(&global_context, "conf/ui.conf")
	if err != nil {
		log.Fatalf("Init Server fail [%s]\n", err.Error())
		return
	}
	http.HandleFunc("/v1/protobuf", CallbackJesgoo)
	http.HandleFunc("/v1/json", CallbackJesgooJson)
	listen_str := fmt.Sprintf(":%d", global_context.Server.ListenPort)
	log.Printf("start server . port [%s]", listen_str)
	err = http.ListenAndServe(listen_str, nil)
	if err != nil {
		utils.FatalLog.Write("start server fail . err[%s]", err.Error())
	}
	log.Printf("start server success . port [%s]", listen_str)
	return
}
