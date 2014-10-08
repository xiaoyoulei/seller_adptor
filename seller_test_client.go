package main

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"io/ioutil"
	"jesgoo_protocol"
	"log"
	"net/http"
)

func main() {

	var req_body jesgoo_protocol.SellerRequest
	var resp_body jesgoo_protocol.SellerResponse
	var media_tmp jesgoo_protocol.Media
	media_tmp.Id = new(string)
	*media_tmp.Id = "123"
	media_tmp.ChannelId = new(string)
	*media_tmp.ChannelId = "123"
	media_tmp.Type = new(jesgoo_protocol.MediaType)
	*media_tmp.Type = jesgoo_protocol.MediaType_APP
	req_body.Media = new(jesgoo_protocol.Media)
	*req_body.Media = media_tmp
	var device_tmp jesgoo_protocol.Device
	device_tmp.Type = new(jesgoo_protocol.DeviceType)
	*device_tmp.Type = jesgoo_protocol.DeviceType_DEV_PHONE
	device_tmp.OsType = new(jesgoo_protocol.OSType)
	*device_tmp.OsType = jesgoo_protocol.OSType_ANDROID
	device_tmp.OsVersion = new(jesgoo_protocol.Version)
	device_tmp.OsVersion.Major = new(uint32)
	*device_tmp.OsVersion.Major = 1
	req_body.Device = new(jesgoo_protocol.Device)
	*req_body.Device = device_tmp
	req_body.Network = new(jesgoo_protocol.Network)

	var client_tmp jesgoo_protocol.Client
	client_tmp.Type = new(jesgoo_protocol.ClientType)
	*client_tmp.Type = jesgoo_protocol.ClientType_NATIVESDK
	client_tmp.Version = new(jesgoo_protocol.Version)
	client_tmp.Version.Major = new(uint32)
	*client_tmp.Version.Major = 1
	req_body.Client = new(jesgoo_protocol.Client)
	*req_body.Client = client_tmp

	fmt.Println(req_body)
	b, err := proto.Marshal(&req_body)
	if err != nil {
		fmt.Println("format err:", err)
	}

	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post("http://localhost:8081/v1/protobuf", "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = proto.Unmarshal(result, &resp_body)
	fmt.Println("response is: " + resp_body.String())
}
