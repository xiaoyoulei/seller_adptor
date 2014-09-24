package main

import (
	"context"
	"log"
	"net/http"
	"pack"
	"parser"
	"reqads"
)

func InitServer() {
	//	reqads.InitReqBs("127.0.0.1", "8084")
	reqads.InitReqBs("218.244.131.175", "8900")
	log.Println("init server succ")
}
func CallbackTanx(resp http.ResponseWriter, req *http.Request) {
	var inner_data *context.Context
	inner_data = new(context.Context)
	parser.ParseTanxRequest(req, inner_data)
	reqads.ReqBs(inner_data)
	pack.PackTanxResponse(inner_data)
	resp.Write([]byte("hello"))
	return
}
