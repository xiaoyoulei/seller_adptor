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
type DeviceId struct {
	Type    int
	Id      string
	Compact bool
	Md5     bool
}
type Device struct {
	Type    int
	Os_type int
	Brand   string
	Model   string
	Ids     []DeviceId
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

func (this *ParseJesgooJsonRequestModule) Init(inner_data *context.GlobalContext) (err error) {
	return

}

func (this *ParseJesgooJsonRequestModule) Run(inner_data *context.Context) (err error) {

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
	temp_req_device := &temp_req.Device
	if temp_req_device == nil {
		log.Println("request has no device")
	} else {
		switch temp_req_device.Os_type {
		case 1:
			inner_device.OSType = context.OSType_ANDROID
		case 2:
			inner_device.OSType = context.OSType_IOS
		case 3:
			inner_device.OSType = context.OSType_WP
		default:
			inner_device.OSType = context.OSType_UNKNOWN
		}
		if temp_req_device.Ids != nil {
			var device_id DeviceId
			if len(temp_req_device.Ids) > 0 {
				device_id = temp_req_device.Ids[0]
				var inner_device_id context.DeviceID
				//				inner_device_id.DevIDType = device_id.Type
				inner_device_id.ID = device_id.Id
				inner_device.DevID = append(inner_device.DevID, inner_device_id)
			}

		}

	}
	return
}
