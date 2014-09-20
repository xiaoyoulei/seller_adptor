package main

import(
	 "net/http"
	 "context"
	 "parser"
	 "pack"
		)

func CallbackTanx (resp http.ResponseWriter, req *http.Request) {
	var inner_data *context.Context
	inner_data = new(context.Context)
	parser.ParseTanxRequest(req, inner_data)
	pack.PackTanxResponse(inner_data)
	resp.Write([]byte(inner_data.Resp.Content))
	return
}
