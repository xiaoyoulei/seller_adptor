package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"pack"
	"parser"
	"rank"
	"reqads"
)

type IModule interface {
	Run(inner_data *context.Context) (err error)
	Init(inner_data *context.GlobalContext) (err error)
}

var jesgoo_models []IModule
var jesgoo_json_modules []IModule

func InitServer() {
	log.Println("init server succ")
	var global_context context.GlobalContext
	jesgoo_models = append(jesgoo_models, &parser.ParseJesgooRequestModule{})
	jesgoo_models = append(jesgoo_models, &reqads.ReqBSModule{})
	jesgoo_models = append(jesgoo_models, &rank.RankModule{})
	jesgoo_models = append(jesgoo_models, &pack.PackJesgooResponseModule{})

	for i := 0; i < len(jesgoo_models); i++ {
		jesgoo_models[i].Init(&global_context)
	}
	jesgoo_json_modules = append(jesgoo_json_modules, &parser.ParseJesgooJsonRequestModule{})
	jesgoo_json_modules = append(jesgoo_json_modules, &reqads.ReqBSModule{})
	jesgoo_json_modules = append(jesgoo_json_modules, &rank.RankModule{})
	jesgoo_json_modules = append(jesgoo_json_modules, &pack.PackJesgooResponseJsonModule{})
	for i := 0; i < len(jesgoo_json_modules); i++ {
		jesgoo_json_modules[i].Init(&global_context)
	}
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

func CallbackJesgooJson(resp http.ResponseWriter, req *http.Request) {
	var inner_data *context.Context
	inner_data = InitThread()

	var err error
	if req.Body == nil {
		log.Println("req.Body is nil")
		return
	}
	inner_data.ReqBody, err = ioutil.ReadAll(req.Body)

	for i := 0; i < len(jesgoo_json_modules); i++ {
		err = jesgoo_json_modules[i].Run(inner_data)
		if err != nil {
			log.Printf("run module %d fail !", i)
			resp.Write([]byte("error"))
			return
		}
	}
	resp.Write(inner_data.RespBody)

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
	for i := 0; i < len(jesgoo_models); i++ {
		err = jesgoo_models[i].Run(inner_data)
		if err != nil {
			log.Printf("run module %d fail !", i)
			resp.Write([]byte("error"))
			return
		}
	}
	resp.Write(inner_data.RespBody)

	return
}
