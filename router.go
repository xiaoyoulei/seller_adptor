package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"pack"
	"parser"
	"reqads"
)

type IModule interface {
	Run(inner_data *context.Context) (err error)
	Init(inner_data *context.Context) (err error)
}

func InitThread() (inner_data *context.Context) {
	inner_data = new(context.Context)
	if inner_data == nil {
		log.Fatal("first inner_data is nil")
	}
	inner_data.ReqBody = make([]byte, 0)
	inner_data.RespBody = make([]byte, 0)
	return
}

func InitServer() {
	//	reqads.InitReqBs("127.0.0.1", "8084")
	reqads.InitReqBs("218.244.131.175", "8900")
	log.Println("init server succ")
}

func CallbackJesgoo(resp http.ResponseWriter, req *http.Request) {
	var inner_data *context.Context
	inner_data = InitThread()

	var err error
	if req.Body == nil {
		log.Fatal("req.body is nil")
		return
	}
	if inner_data.ReqBody == nil {
		log.Fatal("reqbody is nil")
	}
	inner_data.ReqBody, err = ioutil.ReadAll(req.Body)
	var module IModule
	module = parser.ParseJesgooRequestModule{}
	err = module.Run(inner_data)
	if err != nil {
		log.Println("parse" + err.Error())
	}
	module = reqads.ReqBSModule{}
	err = module.Run(inner_data)
	if err != nil {
		log.Println("router" + err.Error())
		return
	}
	module = pack.PackJesgooResponseModule{}
	err = module.Run(inner_data)
	if err != nil {
		resp.Write([]byte("error"))
		log.Println("pack ads fail")
		return
	}
	resp.Write(inner_data.RespBody)

	return
}
