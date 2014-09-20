package pack

import(
	  "context"
		)

func PackTanxResponse(inner_data *context.Context) {
	inner_data.Resp.Content = "hello_world"
	return 
}
