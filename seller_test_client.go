package main

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"io/ioutil"
	"jesgoo_interface"
	"log"
	"net/http"
)

func main() {

	var req_body jesgoo_interface.SellerRequest
	var resp_body jesgoo_interface.SellerResponse
	var media_tmp jesgoo_interface.Media
	media_tmp.Id = new(string)
	*media_tmp.Id = "123"
	media_tmp.ChannelId = new(string)
	*media_tmp.ChannelId = "123"
	media_tmp.Type = new(jesgoo_interface.MediaType)
	*media_tmp.Type = jesgoo_interface.MediaType_APP
	req_body.Media = new(jesgoo_interface.Media)
	*req_body.Media = media_tmp
	var device_tmp jesgoo_interface.Device
	device_tmp.Type = new(jesgoo_interface.DeviceType)
	*device_tmp.Type = jesgoo_interface.DeviceType_DEV_PHONE
	device_tmp.OsType = new(jesgoo_interface.OSType)
	*device_tmp.OsType = jesgoo_interface.OSType_ANDROID
	device_tmp.OsVersion = new(jesgoo_interface.Version)
	device_tmp.OsVersion.Major = new(uint32)
	*device_tmp.OsVersion.Major = 1
	req_body.Device = new(jesgoo_interface.Device)
	*req_body.Device = device_tmp
	req_body.Network = new(jesgoo_interface.Network)
	req_body.Network.Ip = new(string)
	*req_body.Network.Ip = "61.171.44.61" 
	req_body.Adslots = make([]jesgoo_interface.AdSlot,0)
	var adslot jesgoo_interface.AdSlot 
	adslot.Id = new(string)
	*adslot.Id = "123"

	var client_tmp jesgoo_interface.Client
	client_tmp.Type = new(jesgoo_interface.ClientType)
	*client_tmp.Type = jesgoo_interface.ClientType_NATIVESDK
	client_tmp.Version = new(jesgoo_interface.Version)
	client_tmp.Version.Major = new(uint32)
	*client_tmp.Version.Major = 1
	req_body.Client = new(jesgoo_interface.Client)
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
	fmt.Printf("response header is: %d\n", res.StatusCode)
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = proto.Unmarshal(result, &resp_body)
	fmt.Println("response is: " + resp_body.String())
}
