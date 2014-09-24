package main

import (
	"./gen-go/bs"
	"fmt"
	"net"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
//	"time"
)

func main() {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("218.244.131.175", "8900"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := bs.NewBSServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:8084", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	bs_req := new(bs.BSRequest)
	bs_req.Os = "android"
	bs_resp := new(bs.BSResponse)
	bs_resp,err = client.Search(bs_req)
//	for i:=0; i < 10; i++ {
//		err := client.Search2()
//		if err != nil {
//			fmt.Println("serach fail")
//			transport, err = thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "8084"))
//			transport.Open()
//			useTransport = transportFactory.GetTransport(transport)
//			client = bs.NewBSServiceClientFactory(useTransport, protocolFactory)
//		}
//		time.Sleep(1*time.Second)
//	}
//
	fmt.Println(bs_resp)
	fmt.Println("I am exit ")

}
