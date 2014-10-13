package parser

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"utils"
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

type Size struct {
	Width  int32
	Height int32
}
type Adslot struct {
	Id       string
	Type     int
	Size     Size
	Capacity uint32
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

func (this *ParseJesgooJsonRequestModule) parse(inner_data *context.Context) (err error) {

	var temp_req SellerRequest
	err = json.Unmarshal(inner_data.ReqBody, &temp_req)
	if err != nil {
		log.Println("error occur " + err.Error())
		return
	}

	// media
	inner_media := &inner_data.Req.Media
	inner_media.Appsid = temp_req.Media.Id
	inner_media.ChannelId = temp_req.Media.ChannelId
	switch temp_req.Media.Type {
	case 1:
		inner_media.MediaType = context.MediaType_APP
	case 2:
		inner_media.MediaType = context.MediaType_WEB
	case 3:
		inner_media.MediaType = context.MediaType_WAP
	default:
		inner_media.MediaType = context.MediaType_WAP
	}

	//device
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
				switch device_id.Type {
				case 1:
					inner_device_id.DevIDType = context.DeviceIDType_IMEI
				case 2:
					inner_device_id.DevIDType = context.DeviceIDType_MAC
				case 3:
					inner_device_id.DevIDType = context.DeviceIDType_IDFA
				case 4:
					inner_device_id.DevIDType = context.DeviceIDType_AAID
				default:
					inner_device_id.DevIDType = context.DeviceIDType_IMEI
				}
				inner_device_id.ID = device_id.Id
				inner_device.DevID = append(inner_device.DevID, inner_device_id)
			}
		}

		//network
		temp_req_network := &temp_req.Network
		if temp_req_network == nil {
			log.Println("request has no network")
		} else {
			inner_network := &inner_data.Req.Network

			if len(temp_req_network.Ip) > 6 {
				inner_network.Ip = temp_req_network.Ip
				log.Printf("client ip is %s", temp_req_network.Ip)
			}
			switch temp_req_network.Type {
			case 1:
				inner_network.NetworkType = context.NetworkType_WIFI
			case 2:
				inner_network.NetworkType = context.NetworkType_UNKNOWN
			case 3:
				inner_network.NetworkType = context.NetworkType_2G
			case 4:
				inner_network.NetworkType = context.NetworkType_3G
			case 5:
				inner_network.NetworkType = context.NetworkType_4G
			}
		}

		//adslot
		if len(temp_req.Adslots) > 0 {
			temp_req_adslot := &temp_req.Adslots[0]
			inner_adslot := &inner_data.Req.AdSlot
			inner_adslot.Slotid = temp_req_adslot.Id
			switch temp_req_adslot.Type {
			case 1:
				inner_adslot.AdSlotType = context.AdSlotType_BANNER
			case 2:
				inner_adslot.AdSlotType = context.AdSlotType_OFFERWALL
			case 3:
				inner_adslot.AdSlotType = context.AdSlotType_RECOMMEND
			default:
				inner_adslot.AdSlotType = context.AdSlotType_BANNER
			}
			inner_adslot.Size.Width = temp_req_adslot.Size.Width
			inner_adslot.Size.Height = temp_req_adslot.Size.Height
			if temp_req_adslot.Capacity != 0 {
				inner_adslot.Capacity = temp_req_adslot.Capacity
			} else {
				inner_adslot.Capacity = 1
			}
		} else {
			err = errors.New("no adslot info in request")
			return
		}

		//searchid
		if len(inner_data.Req.Device.DevID) > 0 {
			inner_data.Searchid = utils.GenSearchid(inner_data.Req.Device.DevID[0].ID)
		} else {
			inner_data.Searchid = utils.GenSearchid("default")
		}
	}
	return
}

func (this *ParseJesgooJsonRequestModule) Run(inner_data *context.Context) (err error) {
	err = this.parse(inner_data)
	if err != nil {
		log.Printf("parse jesgoo json fail [%s]", err.Error())
	}
	return
}
