package parser

import (
	"code.google.com/p/goprotobuf/proto"
	"context"
	"errors"
	"jesgoo_interface"
	"log"
	"utils"
)

type ParseJesgooRequestModule struct {
}

func (this ParseJesgooRequestModule) Run(inner_data *context.Context) (err error) {
	var temp_req jesgoo_interface.SellerRequest
	proto.Unmarshal(inner_data.ReqBody, &temp_req)
	if err != nil {
		log.Println("err is not null " + err.Error())
	}
	var inner_req *context.InnerReq
	inner_req = &inner_data.Req

	//media
	inner_media := &inner_req.Media
	if temp_req.Media == nil {
		err = errors.New("request media is null")
		return
	}
	if temp_req.Media.Id == nil || temp_req.Media.ChannelId == nil || temp_req.Media.Type == nil {
		err = errors.New("request media.id or media.channelid is nil")
		return
	}
	inner_media.Appsid = *temp_req.Media.Id
	inner_media.ChannelId = *temp_req.Media.ChannelId
	switch *temp_req.Media.Type {
	case jesgoo_interface.MediaType_APP:
		inner_media.MediaType = context.MediaType_APP
	case jesgoo_interface.MediaType_WAP:
		inner_media.MediaType = context.MediaType_WAP
	case jesgoo_interface.MediaType_WEB:
		inner_media.MediaType = context.MediaType_WEB
	default:
		inner_media.MediaType = context.MediaType_APP
	}

	//device
	inner_device := &inner_req.Device
	if temp_req.Device != nil {
		if temp_req.Device.OsType != nil {
			switch *temp_req.Device.OsType {
			case jesgoo_interface.OSType_ANDROID:
				inner_device.OSType = context.OSType_ANDROID
			case jesgoo_interface.OSType_IOS:
				inner_device.OSType = context.OSType_IOS
			case jesgoo_interface.OSType_WP:
				inner_device.OSType = context.OSType_WP
			default:
				inner_device.OSType = context.OSType_UNKNOWN
			}
		}
		if len(temp_req.Device.Ids) > 0 {
			var device_id context.DeviceID
			switch *temp_req.Device.Ids[0].Type {
			case jesgoo_interface.DeviceIDType_IMEI:
				device_id.DevIDType = context.DeviceIDType_IMEI
			case jesgoo_interface.DeviceIDType_MAC:
				device_id.DevIDType = context.DeviceIDType_MAC
			case jesgoo_interface.DeviceIDType_IDFA:
				device_id.DevIDType = context.DeviceIDType_IDFA
			default:
				device_id.DevIDType = context.DeviceIDType_IMEI
			}
			device_id.ID = string(temp_req.Device.Ids[0].Id)
			inner_device.DevID = append(inner_device.DevID, device_id)
		}
	}

	//network
	inner_network := &inner_req.Network
	if temp_req.Network != nil {
		if temp_req.Network.Ip != nil {
			inner_network.Ip = *temp_req.Network.Ip
		}
		if temp_req.Network.Type != nil {
			switch *temp_req.Network.Type {
			case jesgoo_interface.NetworkType_WIFI:
				inner_network.NetworkType = context.NetworkType_WIFI
			case jesgoo_interface.NetworkType_CELLULAR_UNKNOWN:
				inner_network.NetworkType = context.NetworkType_UNKNOWN
			case jesgoo_interface.NetworkType_CELLULAR_2G:
				inner_network.NetworkType = context.NetworkType_2G
			case jesgoo_interface.NetworkType_CELLULAR_3G:
				inner_network.NetworkType = context.NetworkType_3G
			case jesgoo_interface.NetworkType_CELLULAR_4G:
				inner_network.NetworkType = context.NetworkType_4G
			default:
				inner_network.NetworkType = context.NetworkType_UNKNOWN
			}
		}
	}

	//adslot
	if len(temp_req.Adslots) > 0 {
		inner_adslot := &inner_req.AdSlot
		temp_adslot := *temp_req.Adslots[0]
		if temp_adslot.Id != nil {
			inner_adslot.Slotid = *temp_adslot.Id
		} else {
			inner_adslot.Slotid = "0"
		}
		if temp_adslot.Type != nil {
			switch *temp_adslot.Type {
			case jesgoo_interface.AdSlotType_BANNER:
				inner_adslot.AdSlotType = context.AdSlotType_BANNER
			case jesgoo_interface.AdSlotType_OFFERWALL:
				inner_adslot.AdSlotType = context.AdSlotType_OFFERWALL
			case jesgoo_interface.AdSlotType_RECOMMEND:
				inner_adslot.AdSlotType = context.AdSlotType_RECOMMEND
			default:
				inner_adslot.AdSlotType = context.AdSlotType_BANNER
			}
		} else {
			inner_adslot.AdSlotType = context.AdSlotType_BANNER
		}
		if temp_adslot.Capacity != nil {
			inner_adslot.Capacity = *temp_adslot.Capacity
		} else {
			inner_adslot.Capacity = 1
		}
	} else {
		err = errors.New("no adslot in request")
		return
	}

	//searchid
	if len(inner_data.Req.Device.DevID) > 0 {
		inner_data.Searchid = utils.GenSearchid(inner_data.Req.Device.DevID[0].ID)
	} else {
		inner_data.Searchid = utils.GenSearchid("default")
	}

	return
}

func (this ParseJesgooRequestModule) Init(inner_data *context.GlobalContext) (err error) {
	return
}
