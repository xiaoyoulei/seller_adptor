package main

import (
	"code.google.com/p/gcfg"
	"context"
	"dict"
	"dmp"
	"io/ioutil"
	"log"
	"net/http"
	"pack"
	"parser"
	"prepack"
	"rank"
	"reqads"
	"searchlog"
	"utils"
)

type IModule interface {
	Run(inner_data *context.Context) (err error)
	Init(inner_data *context.GlobalContext) (err error)
}

var jesgoo_modules []IModule
var jesgoo_json_modules []IModule

func InitServer(global_context *context.GlobalContext, conf_path string) (err error) {
	log.Println("init server succ")
	err = gcfg.ReadFileInto(global_context, conf_path)
	if err != nil {
		log.Fatalf("init global conf fail . err[%s]", err.Error())
	}

	/************ init log **************/
	switch global_context.Log.LogLevel {
	case 1:
		utils.GlobalLogLevel = utils.NoticeLevel
	case 2:
		utils.GlobalLogLevel = utils.FatalLevel
	case 3:
		utils.GlobalLogLevel = utils.WarningLevel
	case 4:
		utils.GlobalLogLevel = utils.DebugLevel
	default:
		utils.GlobalLogLevel = utils.DebugLevel
	}
	//	utils.GlobalLogLevel = utils.WarningLevel
	utils.DebugLog = &utils.LogControl{}
	utils.FatalLog = &utils.LogControl{}
	utils.WarningLog = &utils.LogControl{}
	utils.NoticeLog = &utils.LogControl{}
	err = utils.DebugLog.Init(60, "ui.dg", "./log/", utils.DebugLevel)
	if err != nil {
		return
	}
	err = utils.FatalLog.Init(60, "ui.fatal", "./log/", utils.FatalLevel)
	if err != nil {
		return
	}
	err = utils.WarningLog.Init(60, "ui.warn", "./log/", utils.WarningLevel)
	if err != nil {
		return
	}
	err = utils.NoticeLog.Init(60, "ui.log", "./log/", utils.NoticeLevel)
	if err != nil {
		return
	}
	/************ init log end **********/

	/************ init dict  **********/
	dict.IPDict = &dict.IPDictModule{}
	err = dict.IPDict.Init(global_context)
	if err != nil {
		utils.FatalLog.Write("Init ipdict fail. err[%s]", err.Error())
		return
	}
	/************ init dict end**********/

	/************ init module ************/
	jesgoo_modules = append(jesgoo_modules, &parser.ParseJesgooRequestModule{})
	jesgoo_modules = append(jesgoo_modules, &dmp.DMPModule{})
	jesgoo_modules = append(jesgoo_modules, &reqads.ReqQiushiModule{})
	jesgoo_modules = append(jesgoo_modules, &rank.RankModule{})
	jesgoo_modules = append(jesgoo_modules, &prepack.PrePackModule{})
	jesgoo_modules = append(jesgoo_modules, &pack.PackJesgooResponseModule{})
	jesgoo_modules = append(jesgoo_modules, &searchlog.SearchLogModule{})

	for i := 0; i < len(jesgoo_modules); i++ {
		jesgoo_modules[i].Init(global_context)
	}
	jesgoo_json_modules = append(jesgoo_json_modules, &parser.ParseJesgooJsonRequestModule{})
	jesgoo_json_modules = append(jesgoo_json_modules, &dmp.DMPModule{})
	jesgoo_json_modules = append(jesgoo_json_modules, &reqads.ReqQiushiModule{})
	jesgoo_json_modules = append(jesgoo_json_modules, &rank.RankModule{})
	jesgoo_json_modules = append(jesgoo_json_modules, &prepack.PrePackModule{})
	jesgoo_json_modules = append(jesgoo_json_modules, &pack.PackJesgooResponseJsonModule{})
	jesgoo_json_modules = append(jesgoo_json_modules, &searchlog.SearchLogModule{})
	for i := 0; i < len(jesgoo_json_modules); i++ {
		jesgoo_json_modules[i].Init(global_context)
	}
	/************ init module end************/
	return
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

	if req.Header["Remoteaddr"] != nil {
		inner_data.Req.Network.Ip = req.Header["Remoteaddr"][0]
	}
	var err error
	if req.Body == nil {
		utils.WarningLog.Write("req.Body is nil")
		return
	}
	inner_data.ReqBody, err = ioutil.ReadAll(req.Body)

	for i := 0; i < len(jesgoo_json_modules); i++ {
		err = jesgoo_json_modules[i].Run(inner_data)
		if err != nil {
			utils.FatalLog.Write("run module %d fail ! err[%s]", i, err.Error())
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	resp.Write(inner_data.RespBody)

}

func CallbackJesgoo(resp http.ResponseWriter, req *http.Request) {
	var inner_data *context.Context
	inner_data = InitThread()

	if req.Header["Remoteaddr"] != nil {
		inner_data.Req.Network.Ip = req.Header["Remoteaddr"][0]
	}
	var err error
	if req.Body == nil {
		log.Fatal("req.body is nil")
		return
	}
	if inner_data.ReqBody == nil {
		log.Fatal("reqbody is nil")
	}
	inner_data.ReqBody, err = ioutil.ReadAll(req.Body)
	for i := 0; i < len(jesgoo_modules); i++ {
		err = jesgoo_modules[i].Run(inner_data)
		if err != nil {
			utils.FatalLog.Write("run module %d fail ! err[%s]", i, err.Error())
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	resp.Write(inner_data.RespBody)

	return
}
