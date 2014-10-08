package parser

import (
	"code.google.com/p/goprotobuf/proto"
	"context"
	"jesgoo_protocol"
	"log"
	"utils"
)

type ParseJesgooRequestModule struct {
}

func (this ParseJesgooRequestModule) Run(inner_data *context.Context) (err error) {
	var temp_req jesgoo_protocol.SellerRequest
	proto.Unmarshal(inner_data.ReqBody, &temp_req)
	if err != nil {
		log.Println("err is not null " + err.Error())
	}
	var inner_req *context.InnerReq
	inner_req = &inner_data.Req
	var inner_media *context.MediaInfo
	inner_media = &inner_req.Media
	inner_media.Appsid = *temp_req.Media.Id
	inner_media.ChannelId = *temp_req.Media.ChannelId
	var inner_device *context.Device
	inner_device = &inner_req.Device
	switch *temp_req.Device.OsType {
	case jesgoo_protocol.OSType_ANDROID:
		inner_device.OSType = context.OSType_ANDROID
	case jesgoo_protocol.OSType_IOS:
		inner_device.OSType = context.OSType_IOS
	case jesgoo_protocol.OSType_WP:
		inner_device.OSType = context.OSType_WP
	default:
		inner_device.OSType = context.OSType_UNKNOWN
	}
	inner_data.Req = *inner_req
	inner_data.Searchid = utils.GenSearchid("123")

	return
}

func (this ParseJesgooRequestModule) Init(inner_data *context.GlobalContext) (err error) {
	return
}
