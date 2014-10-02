package parser

import (
	"context"
	"encoding/json"
	"log"
)

type Media struct {
	Id        string
	ChannelId string
	Type      int
}
type Device struct {
	Type    int
	Os_type int
	Brand   string
	Model   string
}
type Network struct {
	Ip               string
	Type             int
	CellularOperator int
	CellularId       int
}
type Client struct {
	Type int
}

type Adslot struct {
	Id   string
	Type int
}
type SellerRequest struct {
	Media   Media
	Device  Device
	Network Network
	Client  Client
	Adslots []Adslot
	Debug   bool
}
type ParseJesgooJsonRequestModule struct {
}

func (this ParseJesgooJsonRequestModule) Init(inner_data *context.Context) (err error) {
	return

}

func (this ParseJesgooJsonRequestModule) Run(inner_data *context.Context) (err error) {

	var temp_req SellerRequest
	err = json.Unmarshal(inner_data.ReqBody, &temp_req)
	if err != nil {
		log.Println("error occur " + err.Error())
		return
	}
	inner_media := &inner_data.Req.Media
	inner_media.Appsid = temp_req.Media.Id
	inner_media.ChannelId = temp_req.Media.ChannelId
	inner_device := &inner_data.Req.Device
	switch temp_req.Device.Os_type {
	case 1:
		inner_device.OSType = context.OSType_ANDROID
	case 2:
		inner_device.OSType = context.OSType_IOS
	case 3:
		inner_device.OSType = context.OSType_WP
	default:
		inner_device.OSType = context.OSType_UNKNOWN
	}
	log.Printf("json os type %d\n", temp_req.Device.Os_type)
	log.Println(inner_data.Req)

	return
}
