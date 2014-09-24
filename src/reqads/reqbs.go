package reqads

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"bs"
	"log"
)

var transportFactory thrift.TTransportFactory
var protocolFactory *thrift.TBinaryProtocolFactory
var transport *thrift.TSocket
var useTransport thrift.TTransport
var client *bs.BSServiceClient
var err error
func InitReqBs(host string, port string) {

	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	transport, err = thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal("create transport fail")
	}
	
	useTransport = transportFactory.GetTransport(transport)
	client = bs.NewBSServiceClientFactory(useTransport, protocolFactory)
	err := transport.Open()
	if err != nil {
		log.Fatal("open transport fail")
	}

}

func pack_req(inner_data *context.Context, bs_req *bs.BSRequest) {
	 bs_req.Os = "android"

}

func ReqBs(inner_data *context.Context) {

	InitReqBs("218.244.131.175", "8900")

	bs_req := new(bs.BSRequest)
	bs_resp := new(bs.BSResponse)
//	transport.Open()
	pack_req(inner_data, bs_req)
	bs_resp,err = client.Search(bs_req)
	if err != nil {
		log.Println("request bs fail")
		log.Println(err.Error())
		log.Println(bs_resp)
	}
	log.Println(bs_resp)
	return 
}
