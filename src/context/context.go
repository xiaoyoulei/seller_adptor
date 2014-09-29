package context

type AppInfo struct {
	PackageName string
}
type MediaInfo struct {
	Appsid    string
	ChannelId string
	App       AppInfo
}
type SizeInfo struct {
	Width  int32
	Height int32
}
type AdSlotInfo struct {
	Slotid string
	Size   SizeInfo
}
type NetworkInfo struct {
	Ip          string
	NetworkType int
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
)

type DeviceID struct {
	DevIDType DeviceIDType
	ID        string
}

type Device struct {
	OSTypeT   OSType
	OSVersion string
	DevID     []DeviceID
}
type InnerReq struct {
	MediaT   MediaInfo
	AdSlotT  AdSlotInfo
	NetworkT NetworkInfo
	DeviceT  Device
}

type AdInfo struct {
	Adid          int64
	Groupid       int64
	Planid        int64
	Userid        int64
	Title         string
	Bid           int64
	Price         int64
	Ctr           int64
	Cpm           int64
	WuliaoType    int32
	Description1  string
	Description2  string
	ImageUrl      string
	ImageSize     SizeInfo
	LogoUrl       string
	LogoSize      SizeInfo
	ClickUrl      string
	ImpressionUrl string
}
type InnerResp struct {
	Ads []AdInfo
}

type Context struct {
	Searchid string
	ReqBody  []byte
	Req      InnerReq
	Resp     InnerResp
	RespBody []byte
}
