package searchlog

import (
	"code.google.com/p/goprotobuf/proto"
	"context"
	"encoding/base64"
	"jesgoo_uilog"
	"time"
	"utils"
)

type SearchLogModule struct {
	basecoder *base64.Encoding
}

func (this *SearchLogModule) Init(conf *context.GlobalContext) (err error) {
	code_char := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	this.basecoder = base64.NewEncoding(code_char)
	return
}

func (this *SearchLogModule) fill_notice(noticelog *jesgoo_uilog.NoticeLogBody, inner_data *context.Context) (err error) {
	noticelog.Searchid = new(string)
	*noticelog.Searchid = inner_data.Searchid
	noticelog.Timestamp = new(uint32)
	*noticelog.Timestamp = uint32(time.Now().Unix())

	//media
	noticelog.Media = new(jesgoo_uilog.Media)
	media := noticelog.Media
	media.Type = new(jesgoo_uilog.MediaType)
	switch inner_data.Req.Media.MediaType {
	case context.MediaType_APP:
		*media.Type = jesgoo_uilog.MediaType_APP
	case context.MediaType_WEB:
		*media.Type = jesgoo_uilog.MediaType_WEB
	case context.MediaType_WAP:
		*media.Type = jesgoo_uilog.MediaType_WAP
	default:
		*media.Type = jesgoo_uilog.MediaType_UNKNOWN
	}
	media.Appsid = new(string)
	*media.Appsid = inner_data.Req.Media.Appsid
	media.Channelid = new(string)
	*media.Channelid = inner_data.Req.Media.ChannelId
	media.App = new(jesgoo_uilog.AppInfo)
	media.App.Packagename = new(string)
	*media.App.Packagename = inner_data.Req.Media.App.PackageName

	//adslot
	noticelog.Adslot = make([]*jesgoo_uilog.Adslot, 0)
	var adslot *jesgoo_uilog.Adslot
	adslot = new(jesgoo_uilog.Adslot)
	inner_adslot := &inner_data.Req.AdSlot
	adslot.Id = new(string)
	*adslot.Id = inner_adslot.Slotid
	adslot.Type = new(jesgoo_uilog.AdslotType)
	switch inner_adslot.AdSlotType {
	case context.AdSlotType_BANNER:
		*adslot.Type = jesgoo_uilog.AdslotType_BANNER
	case context.AdSlotType_OFFERWALL:
		*adslot.Type = jesgoo_uilog.AdslotType_OFFERWALL
	case context.AdSlotType_RECOMMEND:
		*adslot.Type = jesgoo_uilog.AdslotType_RECOMMEND
	default:
		*adslot.Type = jesgoo_uilog.AdslotType_BANNER
	}
	adslot.Size = new(jesgoo_uilog.Size)
	adslot.Size.Width = new(uint32)
	*adslot.Size.Width = uint32(inner_adslot.Size.Width)
	adslot.Size.Height = new(uint32)
	*adslot.Size.Height = uint32(inner_adslot.Size.Height)
	adslot.Capacity = new(uint32)
	*adslot.Capacity = inner_adslot.Capacity
	noticelog.Adslot = append(noticelog.Adslot, adslot)

	//device
	noticelog.Device = new(jesgoo_uilog.Device)
	device := noticelog.Device
	device.Os = new(jesgoo_uilog.OSType)
	switch inner_data.Req.Device.OSType {
	case context.OSType_ANDROID:
		*device.Os = jesgoo_uilog.OSType_OS_ANDROID
	case context.OSType_IOS:
		*device.Os = jesgoo_uilog.OSType_OS_IOS
	case context.OSType_WP:
		*device.Os = jesgoo_uilog.OSType_OS_WP
	default:
		*device.Os = jesgoo_uilog.OSType_OS_UNKNOWN
	}
	device.Osversion = new(jesgoo_uilog.Version)
	device.Osversion.Major = new(uint32)
	*device.Osversion.Major = inner_data.Req.Device.OSVersion.Major
	// left other version id
	device.Ids = make([]*jesgoo_uilog.DeviceId, 0)
	for i := 0; i < len(inner_data.Req.Device.DevID); i++ {
		var tempid *jesgoo_uilog.DeviceId
		tempid = new(jesgoo_uilog.DeviceId)
		innerid := &inner_data.Req.Device.DevID[i]
		tempid.Type = new(jesgoo_uilog.DeviceIdType)
		switch innerid.DevIDType {
		case context.DeviceIDType_IMEI:
			*tempid.Type = jesgoo_uilog.DeviceIdType_DEVID_IMEI
		case context.DeviceIDType_MAC:
			*tempid.Type = jesgoo_uilog.DeviceIdType_DEVID_MAC
		case context.DeviceIDType_IDFA:
			*tempid.Type = jesgoo_uilog.DeviceIdType_DEVID_IDFA
		default:
			*tempid.Type = jesgoo_uilog.DeviceIdType_DEVID_UNKNOWN
		}
		tempid.Id = new(string)
		*tempid.Id = innerid.ID
		device.Ids = append(device.Ids, tempid)
	}
	device.Brand = new(string)
	*device.Brand = inner_data.Req.Device.Brand
	device.Model = new(string)
	*device.Model = inner_data.Req.Device.Model

	//ads return
	noticelog.Ads = make([]*jesgoo_uilog.AdInfo, 0)
	for i := 0; i < len(inner_data.Resp.Ads); i++ {
		var tempad *jesgoo_uilog.AdInfo
		tempad = new(jesgoo_uilog.AdInfo)
		inner_ad := &inner_data.Resp.Ads[i]
		tempad.Type = new(jesgoo_uilog.AdType)
		switch inner_ad.AdType {
		case context.TEXT:
			*tempad.Type = jesgoo_uilog.AdType_TEXT
		case context.IMAGE:
			*tempad.Type = jesgoo_uilog.AdType_IMAGE
		case context.HTML:
			*tempad.Type = jesgoo_uilog.AdType_HTML
		case context.VIDEO:
			*tempad.Type = jesgoo_uilog.AdType_VIDEO
		case context.TEXT_ICON:
			*tempad.Type = jesgoo_uilog.AdType_TEXT_ICON
		}
		tempad.Src = new(jesgoo_uilog.AdSrc)
		switch inner_ad.AdSrc {
		case context.AdSrc_JESGOO:
			*tempad.Src = jesgoo_uilog.AdSrc_JESGOO
		case context.AdSrc_BAIDU:
			*tempad.Src = jesgoo_uilog.AdSrc_BAIDU
		}
		tempad.Interaction = new(jesgoo_uilog.InteractionType)
		switch inner_ad.InteractionType {
		case context.NO_INTERACT:
			*tempad.Interaction = jesgoo_uilog.InteractionType_NO_INTERACT
		case context.SURFING:
			*tempad.Interaction = jesgoo_uilog.InteractionType_SURFING
		case context.DOWNLOAD:
			*tempad.Interaction = jesgoo_uilog.InteractionType_DOWNLOAD
		case context.DIALING:
			*tempad.Interaction = jesgoo_uilog.InteractionType_DIALING
		case context.MESSAGE:
			*tempad.Interaction = jesgoo_uilog.InteractionType_MESSAGE
		case context.MAIL:
			*tempad.Interaction = jesgoo_uilog.InteractionType_MAIL
		}
		tempad.Adid = new(uint32)
		*tempad.Adid = uint32(inner_ad.Adid)
		tempad.Groupid = new(uint32)
		*tempad.Groupid = uint32(inner_ad.Groupid)
		tempad.Planid = new(uint32)
		*tempad.Planid = uint32(inner_ad.Planid)
		tempad.Userid = new(uint32)
		*tempad.Userid = uint32(inner_ad.Userid)
		tempad.Bid = new(uint32)
		*tempad.Bid = uint32(inner_ad.Bid)
		tempad.Price = new(uint32)
		*tempad.Price = uint32(inner_ad.Price)
		tempad.Ctr = new(uint64)
		*tempad.Ctr = uint64(inner_ad.Ctr)
		tempad.Cpm = new(uint64)
		*tempad.Cpm = uint64(inner_ad.Cpm)
		noticelog.Ads = append(noticelog.Ads, tempad)
	}
	//	noticelog.Dspret = make([]*AdDspRet, 0)
	//	var dspret *jesgoo_uilog.AdDspRet
	//	dspret = new(jesgoo_uilog.AdDspRet)
	noticelog.Debug = new(bool)
	*noticelog.Debug = inner_data.Req.Debug
	return
}

func (this *SearchLogModule) Run(inner_data *context.Context) (err error) {
	var noticelog jesgoo_uilog.NoticeLogBody
	err = this.fill_notice(&noticelog, inner_data)
	if err != nil {
		return
	}
	var body []byte
	body, err = proto.Marshal(&noticelog)
	if err != nil {
		return
	}
	logstr := this.basecoder.EncodeToString(body)
	utils.NoticeLog.Write(logstr)
	return
}
