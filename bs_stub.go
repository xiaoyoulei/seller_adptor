package main

import(
	"./gen-go/bs/"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
		)

const (
			NetworkAddr = "127.0.0.1:8084"
	  )

type BsServiceImpl struct {
}

func (this* BsServiceImpl) Search (req *bs.BSRequest) (resp *bs.BSResponse, err error) {
	fmt.Println("start do search")
	return 
}

func (this* BsServiceImpl) Search2 () ( err error) {
	fmt.Println("start do search2")
	return 
}

func main () {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("new socket fail")
	}
	handler := &BsServiceImpl{}
	processor := bs.NewBSServiceProcessor(handler)
	
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("server start")
	server.Serve()

}

