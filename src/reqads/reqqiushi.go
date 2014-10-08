package reqads

//package main

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"context"
	"errors"
	"io/ioutil"
	"log"
	"mobads_api"
	"net/http"
)

type ReqQiushiModule struct {
	client *http.Client
}

func (this *ReqQiushiModule) packreq(request *mobads_api.BidRequest, inner_data *context.Context) (err error) {
	request.RequestId = new(string)
	*request.RequestId = inner_data.Searchid

	//version parament
	request.ApiVersion = new(mobads_api.Version)
	version_tmp := request.ApiVersion
	version_tmp.Major = new(uint32)
	*version_tmp.Major = 1

	//app parament
	request.App = new(mobads_api.App)
	app_tmp := request.App
	app_tmp.StaticInfo = new(mobads_api.App_StaticInfo)
	app_static_info := app_tmp.StaticInfo
	app_static_info.BundleId = new(string)
	*app_static_info.BundleId = "com.jesgoo.app"
	app_tmp.Id = new(string)
	*app_tmp.Id = "10042c1f"

	//device parament
	request.Device = new(mobads_api.Device)
	device_tmp := request.Device
	device_tmp.Type = new(mobads_api.Device_Type)
	*device_tmp.Type = mobads_api.Device_PHONE
	device_tmp.Os = new(mobads_api.Device_Os)
	*device_tmp.Os = mobads_api.Device_ANDROID
	*version_tmp.Major = 1
	device_tmp.OsVersion = new(mobads_api.Version)
	device_tmp.OsVersion.Major = new(uint32)
	*device_tmp.OsVersion.Major = 1
	device_tmp.Udid = new(mobads_api.Device_UdId)
	device_udid := device_tmp.Udid
	device_udid.Imei = new(string)
	*device_udid.Imei = "013474004923670"
	device_tmp.Vendor = new(string)
	*device_tmp.Vendor = "levnovo"
	device_tmp.Model = new(string)
	*device_tmp.Model = "lenovo"

	//network
	request.Network = new(mobads_api.Network)
	network_tmp := request.Network
	network_tmp.Ipv4 = new(string)
	*network_tmp.Ipv4 = inner_data.Req.Network.Ip

	//adslot
	var adslot_tmp mobads_api.AdSlot
	adslot_tmp.Id = new(string)
	*adslot_tmp.Id = "L0000041"
	var size_tmp mobads_api.Size
	size_tmp.Width = new(uint32)
	size_tmp.Height = new(uint32)
	*size_tmp.Width = 480
	*size_tmp.Height = 96
	adslot_tmp.Size = new(mobads_api.Size)
	*adslot_tmp.Size = size_tmp
	request.Adslots = make([]*mobads_api.AdSlot, 0)
	request.Adslots = append(request.Adslots, &adslot_tmp)
	return
}

func (this *ReqQiushiModule) parse_resp(response *mobads_api.BidResponse, inner_data *context.Context) (err error) {
	log.Printf("req success ans: %s \n", response.String())
	return
}
func (this *ReqQiushiModule) Run(inner_data *context.Context) (err error) {
	//	client := &http.Client{}
	var request_body = mobads_api.BidRequest{}
	err = this.packreq(&request_body, inner_data)
	log.Println(request_body)
	var request_byte []byte
	request_byte = make([]byte, 0)
	request_byte, err = proto.Marshal(&request_body)
	if err != nil {
		log.Printf("proto marshal fail ! %s", err.Error())
		return
	}
	var request *http.Request
	request, err = http.NewRequest("POST", "http://61.135.186.214/api", bytes.NewBuffer(request_byte))
	if err != nil {
		log.Println("create http post request fail")
		return
	}
	request.Header.Set("Content-Type", "application/x-protobuf")
	request.Header.Set("Accept", "application/x-protobuf")
	request.Header.Set("User-Agent", "Jesgoo-API")
	request.Header.Set("Connection", "Keep-alive")
	var response *http.Response
	response, err = this.client.Do(request)
	if err != nil {
		log.Println("request qiushi server fail")
		return
	}
	if response.StatusCode != 200 {
		err = errors.New("qiushi respose code is " + string(response.StatusCode))
		log.Printf("qiushi response code is %d\n", response.StatusCode)
		return
	}
	var response_body = mobads_api.BidResponse{}
	var response_byte []byte
	response_byte, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("error occured %s\n", err.Error())
	}
	err = proto.Unmarshal(response_byte, &response_body)
	if err != nil {
		log.Printf("error occur %s\n", err.Error())
	}
	err = this.parse_resp(&response_body, inner_data)

	return
}

func (this *ReqQiushiModule) Init(inner_data *context.GlobalContext) (err error) {
	this.client = &http.Client{}
	return
}

/*
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var temp context.Context
	temp.Req.Network.Ip = "220.181.111.85"
	var module *ReqQiushiModule
	module = new(ReqQiushiModule)
	module.Init(&temp)
	module.Run(&temp)
}
*/
