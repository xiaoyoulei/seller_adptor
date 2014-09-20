package context

type inner_req_t struct {
	Content string
}

type inner_resp_t struct {
	Content string

}

type Context struct {
Req  inner_req_t
Resp inner_resp_t
}
