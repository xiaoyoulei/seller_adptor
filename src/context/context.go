package context

import "bytes"

type AppInfo struct {
	PackageName string
}

type MediaType int64

const (
	MediaType_APP MediaType = 1
	MediaType_WEB MediaType = 2
	MediaType_WAP MediaType = 3
)

type MediaInfo struct {
	Appsid    string
	ChannelId string
	MediaType MediaType
	App       AppInfo
}
type SizeInfo struct {
	Width  int32
	Height int32
}
type AdSlotType int64

const (
	AdSlotType_BANNER    AdSlotType = 1
	AdSlotType_OFFERWALL AdSlotType = 2
	AdSlotType_RECOMMEND AdSlotType = 3
)

type AdSlotInfo struct {
	Slotid     string
	AdSlotType AdSlotType
	Size       SizeInfo
	Capacity   uint32
}
type NetworkType int64

const (
	NetworkType_WIFI    NetworkType = 1
	NetworkType_UNKNOWN NetworkType = 2
	NetworkType_2G      NetworkType = 3
	NetworkType_3G      NetworkType = 4
	NetworkType_4G      NetworkType = 5
)

type NetworkInfo struct {
	Ip          string
	NetworkType NetworkType
}
type OSType int64

const (
	OSType_UNKNOWN OSType = 0
	OSType_ANDROID OSType = 1
	OSType_IOS     OSType = 2
	OSType_WP      OSType = 3
)

type DeviceIDType int64

const (
	DeviceIDType_IMEI      DeviceIDType = 1
	DeviceIDType_MAC       DeviceIDType = 2
	DeviceIDType_IDFA      DeviceIDType = 3
	DeviceIDType_AAID      DeviceIDType = 4
	DeviceIDType_OPENUDID  DeviceIDType = 5
	DeviceIDType_ANDROIDID DeviceIDType = 6
	DeviceIDType_UDID      DeviceIDType = 7
	DeviceIDType_ODIN      DeviceIDType = 8
	DeviceIDType_DUID      DeviceIDType = 9
	DeviceIDType_UNKNOWN   DeviceIDType = 10
)

type DeviceID struct {
	DevIDType DeviceIDType
	ID        string
}

type Version struct {
	Major uint32
	Minor uint32
	Micro uint32
	Build uint32
}

type Device struct {
	OSType    OSType
	OSVersion Version
	DevID     []DeviceID
	Brand     string
	Model     string
}
type InnerReq struct {
	Media   MediaInfo
	AdSlot  AdSlotInfo
	Network NetworkInfo
	Device  Device
}

type AdType int64

const (
	TEXT      AdType = 1
	IMAGE     AdType = 2
	HTML      AdType = 3
	VIDEO     AdType = 4
	TEXT_ICON AdType = 5
)

type AdSrc int64

const (
	AdSrc_JESGOO AdSrc = 1
	AdSrc_BAIDU  AdSrc = 2
)

type InteractionType int32

const (
	NO_INTERACT InteractionType = 0
	SURFING     InteractionType = 1
	DOWNLOAD    InteractionType = 2
	DIALING     InteractionType = 3
	MESSAGE     InteractionType = 4
	MAIL        InteractionType = 5
)

type AdInfo struct {
	AdType          AdType
	AdSrc           AdSrc
	InteractionType InteractionType
	Adid            int64
	Groupid         int64
	Planid          int64
	Userid          int64
	Title           string
	Bid             int64
	Price           int64
	Ctr             int64
	Cpm             int64
	WuliaoType      int32
	Description1    string
	Description2    string
	ImageUrl        string
	ImageSize       SizeInfo
	LogoUrl         string
	LogoSize        SizeInfo
	ClickUrl        string
	ImpressionUrl   []string
	HtmlSnippet     bytes.Buffer
}
type InnerResp struct {
	Ads []AdInfo
}

type Context struct {
	Searchid  string
	ReqBody   []byte
	Req       InnerReq
	Resp      InnerResp
	BaiduAds  []AdInfo
	JesgooAds []AdInfo
	RespBody  []byte
}

// 服务器全局信息
type GlobalContext struct {
}
