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
	"net"
	"net/http"
	"time"
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
	version_tmp.Minor = new(uint32)
	version_tmp.Micro = new(uint32)
	*version_tmp.Major = 4
	*version_tmp.Minor = 0
	*version_tmp.Micro = 0

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
	*version_tmp.Major = inner_data.Req.Device.OSVersion.Major
	*version_tmp.Minor = inner_data.Req.Device.OSVersion.Minor
	*version_tmp.Micro = inner_data.Req.Device.OSVersion.Micro
	device_tmp.OsVersion = new(mobads_api.Version)
	device_tmp.OsVersion.Major = new(uint32)
	device_tmp.OsVersion.Minor = new(uint32)
	device_tmp.OsVersion.Micro = new(uint32)
	*device_tmp.OsVersion.Major = inner_data.Req.Device.OSVersion.Major
	*device_tmp.OsVersion.Minor = inner_data.Req.Device.OSVersion.Minor
	*device_tmp.OsVersion.Micro = inner_data.Req.Device.OSVersion.Micro
	device_tmp.Udid = new(mobads_api.Device_UdId)
	device_udid := device_tmp.Udid
	for i := 0; i < len(inner_data.Req.Device.DevID); i++ {
		switch inner_data.Req.Device.DevID[0].DevIDType {
		case context.DeviceIDType_IMEI:
			device_udid.Imei = new(string)
			*device_udid.Imei = inner_data.Req.Device.DevID[0].ID
		case context.DeviceIDType_MAC:
			device_udid.Mac = new(string)
			*device_udid.Mac = inner_data.Req.Device.DevID[0].ID
		case context.DeviceIDType_IDFA:
			device_udid.Idfa = new(string)
			*device_udid.Idfa = inner_data.Req.Device.DevID[0].ID
		default:
			device_udid.Imei = new(string)
			*device_udid.Imei = inner_data.Req.Device.DevID[0].ID
		}
	}
	device_tmp.Vendor = new(string)
	*device_tmp.Vendor = inner_data.Req.Device.Brand
	device_tmp.Model = new(string)
	*device_tmp.Model = inner_data.Req.Device.Model

	//network
	request.Network = new(mobads_api.Network)
	network_tmp := request.Network
	network_tmp.Ipv4 = new(string)
	*network_tmp.Ipv4 = inner_data.Req.Network.Ip
	log.Printf("req qiushi inner_ip [%s]", inner_data.Req.Network.Ip)

	//adslot
	var adslot_tmp mobads_api.AdSlot
	adslot_tmp.Id = new(string)
	*adslot_tmp.Id = "L0000041"
	var size_tmp mobads_api.Size
	size_tmp.Width = new(uint32)
	size_tmp.Height = new(uint32)
	if inner_data.Req.AdSlot.Size.Width != 0 {
		*size_tmp.Width = uint32(inner_data.Req.AdSlot.Size.Width)
		*size_tmp.Height = uint32(inner_data.Req.AdSlot.Size.Height)
	} else {
		*size_tmp.Width = 320
		*size_tmp.Height = 48
	}
	adslot_tmp.Size = new(mobads_api.Size)
	*adslot_tmp.Size = size_tmp
	request.Adslots = make([]*mobads_api.AdSlot, 0)
	request.Adslots = append(request.Adslots, &adslot_tmp)
	return
}

func (this *ReqQiushiModule) convert_ad(inad *context.AdInfo, bsad *mobads_api.Ad) (err error) {
	if bsad.AdId != nil {
		inad.Adid = int64(*bsad.AdId)
	}
	//	inad.Groupid = bsad.Groupid
	//	inad.Planid = bsad.Groupid
	//	inad.Userid = bsad.Userid
	admeta := bsad.MaterialMeta
	if admeta != nil {

		inad.AdSrc = context.AdSrc_BAIDU

		if admeta.CreativeType != nil {
			switch *admeta.CreativeType {
			case mobads_api.CreativeType_TEXT:
				inad.AdType = context.TEXT
			case mobads_api.CreativeType_IMAGE:
				inad.AdType = context.IMAGE
			case mobads_api.CreativeType_HTML:
				inad.AdType = context.HTML
			case mobads_api.CreativeType_VIDEO:
				inad.AdType = context.VIDEO
			case mobads_api.CreativeType_TEXT_ICON:
				inad.AdType = context.TEXT_ICON
			}
		}
		if admeta.InteractionType != nil {
			switch *admeta.InteractionType {
			case mobads_api.InteractionType_SURFING:
				inad.InteractionType = context.SURFING
			case mobads_api.InteractionType_DOWNLOAD:
				inad.InteractionType = context.DOWNLOAD
			case mobads_api.InteractionType_DIALING:
				inad.InteractionType = context.DIALING
			case mobads_api.InteractionType_MESSAGE:
				inad.InteractionType = context.MESSAGE
			case mobads_api.InteractionType_MAIL:
				inad.InteractionType = context.MAIL
			default:
				inad.InteractionType = context.NO_INTERACT

			}
		}
		if admeta.Title != nil {
			inad.Title = *admeta.Title
		}
		if admeta.Description1 != nil {
			inad.Description1 = *admeta.Description1
		}
		if admeta.Description2 != nil {
			inad.Description2 = *admeta.Description2
		}
		if admeta.MediaUrl != nil {
			inad.ImageUrl = *admeta.MediaUrl
		}
		for i := 0; i < len(admeta.WinNoticeUrl); i++ {
			inad.ImpressionUrl = append(inad.ImpressionUrl, admeta.WinNoticeUrl[i])
		}
		if admeta.ClickUrl != nil {
			inad.ClickUrl = *admeta.ClickUrl
		}
	}
	inad.Bid = 0
	inad.Price = 0
	inad.Ctr = 0
	inad.Cpm = 0
	//	inad.Wuliao_type = bsad.WuliaoType
	//	inad.LogoUrl = bsad.AppLogo
	return
}

func (this *ReqQiushiModule) parse_resp(response *mobads_api.BidResponse, inner_data *context.Context) (err error) {
	log.Println("baidu_response")
	log.Println(response)
	if response.ErrorCode != nil {
		log.Printf("request qiushi fail . error_code is %u\n", *response.ErrorCode)
		err = errors.New("request qiushi fail .")
		return
	}
	for i := 0; i < len(response.Ads); i++ {
		var inner_ad context.AdInfo
		err = this.convert_ad(&inner_ad, response.Ads[i])
		if err != nil {
			continue
		}
		inner_data.BaiduAds = append(inner_data.BaiduAds, inner_ad)
	}
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
	request, err = http.NewRequest("POST", "http://mobads.baidu.com/api", bytes.NewBuffer(request_byte))
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
		return
	}
	err = proto.Unmarshal(response_byte, &response_body)
	if err != nil {
		log.Printf("error occur %s\n", err.Error())
		return
	}
	err = this.parse_resp(&response_body, inner_data)

	return
}

func (this *ReqQiushiModule) Init(inner_data *context.GlobalContext) (err error) {
	this.client = &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Millisecond*500)
				if err != nil {
					log.Println("dail timeout", err)
					return nil, err
				}
				return c, nil
			},
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Millisecond * 500,
		},
	}
	return
}

/*
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var temp context.Context
	temp.Req.Network.Ip = "220.181.111.85"
	var module *ReqQiushiModule
	module = new(ReqQiushiModule)
	var temp_global context.GlobalContext
	module.Init(&temp_global)
	module.Run(&temp)
}
*/
