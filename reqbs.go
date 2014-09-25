package main

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
	"ui2bs"
)

var transportFactory thrift.TTransportFactory
var protocolFactory *thrift.TBinaryProtocolFactory
var transport *thrift.TSocket
var useTransport thrift.TTransport
var client *ui2bs.BSServiceClient
var err error

func InitReqBs(host string, port string) {

	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	transport, err = thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal("create transport fail")
	}

	useTransport = transportFactory.GetTransport(transport)
	client = ui2bs.NewBSServiceClientFactory(useTransport, protocolFactory)
	err := transport.Open()
	if err != nil {
		log.Fatal("open transport fail")
	}

}

func pack_req(inner_data *context.Context, bs_req *ui2bs.BSRequest) {
	bs_req.Device.Os = ui2bs.OSType_IOS

}

func ReqBs(inner_data *context.Context) {

	bs_req := new(ui2bs.BSRequest)
	bs_req.Media = new(ui2bs.Media)
	bs_req.Device = new(ui2bs.Device)
	bs_req.Adslot = new(ui2bs.AdSlot)
	bs_req.Searchid = "123"
	bs_req.Media.Appsid = "123"
	bs_req.Media.ChannelId = "123"
	bs_req.Device.Os = ui2bs.OSType_IOS
	bs_resp := new(ui2bs.BSResponse)
	transport.Open()
	defer transport.Close()
	pack_req(inner_data, bs_req)
	bs_resp, err = client.Search(bs_req)
	if err != nil {
		log.Println("request bs fail")
		log.Println(err)
	}
	log.Println(bs_resp)
	return
}

func main () {
	InitReqBs("218.244.131.175", "8900")
	data := new(context.Context)
	ReqBs(data)
}
