package main

import(
	"log"	
	"net/http"
		)

func main () {

	log.Println("start server")
	
	InitServer()
	http.HandleFunc("/", CallbackTanx)
	http.ListenAndServe(":8081", nil)
}
