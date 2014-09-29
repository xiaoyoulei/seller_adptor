// Code generated by protoc-gen-go.
// source: jesgoo_interface.proto
// DO NOT EDIT!

/*
Package jesgoo_protocol is a generated protocol buffer package.

Jesgoo API

It is generated from these files:
	jesgoo_interface.proto

It has these top-level messages:
	Version
	Size
	App
	Site
	Browser
	Media
	WiFi
	Network
	DeviceID
	Device
	Geo
	Client
	AdSlot
	AdNativeMaterial
	SellerRequest
	SellerResponse
*/
package jesgoo_protocol

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// 设备类型
type DeviceType int32

const (
	DeviceType_DEV_UNKNOWN DeviceType = 0
	DeviceType_PC          DeviceType = 1
	DeviceType_DEV_PHONE   DeviceType = 2
	DeviceType_TABLET      DeviceType = 3
	DeviceType_TV          DeviceType = 4
)

var DeviceType_name = map[int32]string{
	0: "DEV_UNKNOWN",
	1: "PC",
	2: "DEV_PHONE",
	3: "TABLET",
	4: "TV",
}
var DeviceType_value = map[string]int32{
	"DEV_UNKNOWN": 0,
	"PC":          1,
	"DEV_PHONE":   2,
	"TABLET":      3,
	"TV":          4,
}

func (x DeviceType) Enum() *DeviceType {
	p := new(DeviceType)
	*p = x
	return p
}
func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}
func (x *DeviceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeviceType_value, data, "DeviceType")
	if err != nil {
		return err
	}
	*x = DeviceType(value)
	return nil
}

// 媒体类型
type MediaType int32

const (
	MediaType_APP MediaType = 1
	MediaType_WEB MediaType = 2
	MediaType_WAP MediaType = 3
)

var MediaType_name = map[int32]string{
	1: "APP",
	2: "WEB",
	3: "WAP",
}
var MediaType_value = map[string]int32{
	"APP": 1,
	"WEB": 2,
	"WAP": 3,
}

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}
func (x MediaType) String() string {
	return proto.EnumName(MediaType_name, int32(x))
}
func (x *MediaType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MediaType_value, data, "MediaType")
	if err != nil {
		return err
	}
	*x = MediaType(value)
	return nil
}

// 操作系统类型
type OSType int32

const (
	OSType_OS_UNKNOWN OSType = 0
	OSType_ANDROID    OSType = 1
	OSType_IOS        OSType = 2
	OSType_WP         OSType = 3
)

var OSType_name = map[int32]string{
	0: "OS_UNKNOWN",
	1: "ANDROID",
	2: "IOS",
	3: "WP",
}
var OSType_value = map[string]int32{
	"OS_UNKNOWN": 0,
	"ANDROID":    1,
	"IOS":        2,
	"WP":         3,
}

func (x OSType) Enum() *OSType {
	p := new(OSType)
	*p = x
	return p
}
func (x OSType) String() string {
	return proto.EnumName(OSType_name, int32(x))
}
func (x *OSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OSType_value, data, "OSType")
	if err != nil {
		return err
	}
	*x = OSType(value)
	return nil
}

// 设备ID类型
type DeviceIDType int32

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

var DeviceIDType_name = map[int32]string{
	1: "IMEI",
	2: "MAC",
	3: "IDFA",
	4: "AAID",
	5: "OPENUDID",
	6: "ANDROIDID",
	7: "UDID",
	8: "ODIN",
	9: "DUID",
}
var DeviceIDType_value = map[string]int32{
	"IMEI":      1,
	"MAC":       2,
	"IDFA":      3,
	"AAID":      4,
	"OPENUDID":  5,
	"ANDROIDID": 6,
	"UDID":      7,
	"ODIN":      8,
	"DUID":      9,
}

func (x DeviceIDType) Enum() *DeviceIDType {
	p := new(DeviceIDType)
	*p = x
	return p
}
func (x DeviceIDType) String() string {
	return proto.EnumName(DeviceIDType_name, int32(x))
}
func (x *DeviceIDType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeviceIDType_value, data, "DeviceIDType")
	if err != nil {
		return err
	}
	*x = DeviceIDType(value)
	return nil
}

// 广告位类型
type AdSlotType int32

const (
	AdSlotType_BANNER          AdSlotType = 1
	AdSlotType_OFFERWALL       AdSlotType = 2
	AdSlotType_RECOMMEND       AdSlotType = 3
	AdSlotType_INTERSTITIAL    AdSlotType = 4
	AdSlotType_REALTIME_SPLASH AdSlotType = 5
	AdSlotType_CACHED_SPLASH   AdSlotType = 6
	AdSlotType_FEED            AdSlotType = 7
)

var AdSlotType_name = map[int32]string{
	1: "BANNER",
	2: "OFFERWALL",
	3: "RECOMMEND",
	4: "INTERSTITIAL",
	5: "REALTIME_SPLASH",
	6: "CACHED_SPLASH",
	7: "FEED",
}
var AdSlotType_value = map[string]int32{
	"BANNER":          1,
	"OFFERWALL":       2,
	"RECOMMEND":       3,
	"INTERSTITIAL":    4,
	"REALTIME_SPLASH": 5,
	"CACHED_SPLASH":   6,
	"FEED":            7,
}

func (x AdSlotType) Enum() *AdSlotType {
	p := new(AdSlotType)
	*p = x
	return p
}
func (x AdSlotType) String() string {
	return proto.EnumName(AdSlotType_name, int32(x))
}
func (x *AdSlotType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AdSlotType_value, data, "AdSlotType")
	if err != nil {
		return err
	}
	*x = AdSlotType(value)
	return nil
}

// 广告风格
type AdStyle int32

const (
	AdStyle_TEXT      AdStyle = 1
	AdStyle_IMAGE     AdStyle = 2
	AdStyle_ICON_TEXT AdStyle = 3
	AdStyle_SMART_AD  AdStyle = 4
	AdStyle_VIDEO     AdStyle = 5
)

var AdStyle_name = map[int32]string{
	1: "TEXT",
	2: "IMAGE",
	3: "ICON_TEXT",
	4: "SMART_AD",
	5: "VIDEO",
}
var AdStyle_value = map[string]int32{
	"TEXT":      1,
	"IMAGE":     2,
	"ICON_TEXT": 3,
	"SMART_AD":  4,
	"VIDEO":     5,
}

func (x AdStyle) Enum() *AdStyle {
	p := new(AdStyle)
	*p = x
	return p
}
func (x AdStyle) String() string {
	return proto.EnumName(AdStyle_name, int32(x))
}
func (x *AdStyle) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AdStyle_value, data, "AdStyle")
	if err != nil {
		return err
	}
	*x = AdStyle(value)
	return nil
}

// 广告推广类型
type PromotionType int32

const (
	PromotionType_NOACTION  PromotionType = 0
	PromotionType_LP        PromotionType = 1
	PromotionType_DOWNLOAD  PromotionType = 2
	PromotionType_PRO_PHONE PromotionType = 3
	PromotionType_SMS       PromotionType = 4
	PromotionType_EMAIL     PromotionType = 5
)

var PromotionType_name = map[int32]string{
	0: "NOACTION",
	1: "LP",
	2: "DOWNLOAD",
	3: "PRO_PHONE",
	4: "SMS",
	5: "EMAIL",
}
var PromotionType_value = map[string]int32{
	"NOACTION":  0,
	"LP":        1,
	"DOWNLOAD":  2,
	"PRO_PHONE": 3,
	"SMS":       4,
	"EMAIL":     5,
}

func (x PromotionType) Enum() *PromotionType {
	p := new(PromotionType)
	*p = x
	return p
}
func (x PromotionType) String() string {
	return proto.EnumName(PromotionType_name, int32(x))
}
func (x *PromotionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PromotionType_value, data, "PromotionType")
	if err != nil {
		return err
	}
	*x = PromotionType(value)
	return nil
}

// 网络类型
type NetworkType int32

const (
	NetworkType_WIFI             NetworkType = 1
	NetworkType_CELLULAR_UNKNOWN NetworkType = 2
	NetworkType_CELLULAR_2G      NetworkType = 3
	NetworkType_CELLULAR_3G      NetworkType = 4
	NetworkType_CELLULAR_4G      NetworkType = 5
)

var NetworkType_name = map[int32]string{
	1: "WIFI",
	2: "CELLULAR_UNKNOWN",
	3: "CELLULAR_2G",
	4: "CELLULAR_3G",
	5: "CELLULAR_4G",
}
var NetworkType_value = map[string]int32{
	"WIFI":             1,
	"CELLULAR_UNKNOWN": 2,
	"CELLULAR_2G":      3,
	"CELLULAR_3G":      4,
	"CELLULAR_4G":      5,
}

func (x NetworkType) Enum() *NetworkType {
	p := new(NetworkType)
	*p = x
	return p
}
func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}
func (x *NetworkType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NetworkType_value, data, "NetworkType")
	if err != nil {
		return err
	}
	*x = NetworkType(value)
	return nil
}

// 地理坐标类型
type GeoType int32

const (
	GeoType_WGS84 GeoType = 1
	GeoType_GCJ02 GeoType = 2
	GeoType_BD09  GeoType = 3
)

var GeoType_name = map[int32]string{
	1: "WGS84",
	2: "GCJ02",
	3: "BD09",
}
var GeoType_value = map[string]int32{
	"WGS84": 1,
	"GCJ02": 2,
	"BD09":  3,
}

func (x GeoType) Enum() *GeoType {
	p := new(GeoType)
	*p = x
	return p
}
func (x GeoType) String() string {
	return proto.EnumName(GeoType_name, int32(x))
}
func (x *GeoType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GeoType_value, data, "GeoType")
	if err != nil {
		return err
	}
	*x = GeoType(value)
	return nil
}

// 广告客户端类型
type ClientType int32

const (
	ClientType_NATIVESDK ClientType = 1
	ClientType_JSSDK     ClientType = 2
	ClientType_OPENAPI   ClientType = 3
)

var ClientType_name = map[int32]string{
	1: "NATIVESDK",
	2: "JSSDK",
	3: "OPENAPI",
}
var ClientType_value = map[string]int32{
	"NATIVESDK": 1,
	"JSSDK":     2,
	"OPENAPI":   3,
}

func (x ClientType) Enum() *ClientType {
	p := new(ClientType)
	*p = x
	return p
}
func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}
func (x *ClientType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClientType_value, data, "ClientType")
	if err != nil {
		return err
	}
	*x = ClientType(value)
	return nil
}

// 地理定位来源
type GeoSource int32

const (
	GeoSource_NATIVE GeoSource = 1
	GeoSource_BAIDU  GeoSource = 2
)

var GeoSource_name = map[int32]string{
	1: "NATIVE",
	2: "BAIDU",
}
var GeoSource_value = map[string]int32{
	"NATIVE": 1,
	"BAIDU":  2,
}

func (x GeoSource) Enum() *GeoSource {
	p := new(GeoSource)
	*p = x
	return p
}
func (x GeoSource) String() string {
	return proto.EnumName(GeoSource_name, int32(x))
}
func (x *GeoSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GeoSource_value, data, "GeoSource")
	if err != nil {
		return err
	}
	*x = GeoSource(value)
	return nil
}

type SellerResponse_Ad_MaterialType int32

const (
	SellerResponse_Ad_DYNAMIC SellerResponse_Ad_MaterialType = 0
	SellerResponse_Ad_NATIVE  SellerResponse_Ad_MaterialType = 1
)

var SellerResponse_Ad_MaterialType_name = map[int32]string{
	0: "DYNAMIC",
	1: "NATIVE",
}
var SellerResponse_Ad_MaterialType_value = map[string]int32{
	"DYNAMIC": 0,
	"NATIVE":  1,
}

func (x SellerResponse_Ad_MaterialType) Enum() *SellerResponse_Ad_MaterialType {
	p := new(SellerResponse_Ad_MaterialType)
	*p = x
	return p
}
func (x SellerResponse_Ad_MaterialType) String() string {
	return proto.EnumName(SellerResponse_Ad_MaterialType_name, int32(x))
}
func (x *SellerResponse_Ad_MaterialType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SellerResponse_Ad_MaterialType_value, data, "SellerResponse_Ad_MaterialType")
	if err != nil {
		return err
	}
	*x = SellerResponse_Ad_MaterialType(value)
	return nil
}

// 通用版本类型
type Version struct {
	Major            *uint32 `protobuf:"varint,1,req,name=major" json:"major,omitempty"`
	Minor            *uint32 `protobuf:"varint,2,opt,name=minor,def=0" json:"minor,omitempty"`
	Micro            *uint32 `protobuf:"varint,3,opt,name=micro,def=0" json:"micro,omitempty"`
	Build            *uint32 `protobuf:"varint,4,opt,name=build,def=0" json:"build,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}

const Default_Version_Minor uint32 = 0
const Default_Version_Micro uint32 = 0
const Default_Version_Build uint32 = 0

func (m *Version) GetMajor() uint32 {
	if m != nil && m.Major != nil {
		return *m.Major
	}
	return 0
}

func (m *Version) GetMinor() uint32 {
	if m != nil && m.Minor != nil {
		return *m.Minor
	}
	return Default_Version_Minor
}

func (m *Version) GetMicro() uint32 {
	if m != nil && m.Micro != nil {
		return *m.Micro
	}
	return Default_Version_Micro
}

func (m *Version) GetBuild() uint32 {
	if m != nil && m.Build != nil {
		return *m.Build
	}
	return Default_Version_Build
}

// 通用尺寸类型
type Size struct {
	Width            *uint32 `protobuf:"varint,1,req,name=width" json:"width,omitempty"`
	Height           *uint32 `protobuf:"varint,2,req,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Size) Reset()         { *m = Size{} }
func (m *Size) String() string { return proto.CompactTextString(m) }
func (*Size) ProtoMessage()    {}

func (m *Size) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *Size) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

// 应用信息
// 联盟自有流量请求可不携带分类信息
type App struct {
	PackageName      *string  `protobuf:"bytes,1,req,name=package_name" json:"package_name,omitempty"`
	Categories       []uint32 `protobuf:"varint,2,rep,name=categories" json:"categories,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}

func (m *App) GetPackageName() string {
	if m != nil && m.PackageName != nil {
		return *m.PackageName
	}
	return ""
}

func (m *App) GetCategories() []uint32 {
	if m != nil {
		return m.Categories
	}
	return nil
}

// 站点信息
// 联盟自有流量请求可不携带分类信息
type Site struct {
	Domain           []byte   `protobuf:"bytes,1,req,name=domain" json:"domain,omitempty"`
	Categories       []uint32 `protobuf:"varint,2,rep,name=categories" json:"categories,omitempty"`
	Url              []byte   `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Site) Reset()         { *m = Site{} }
func (m *Site) String() string { return proto.CompactTextString(m) }
func (*Site) ProtoMessage()    {}

func (m *Site) GetDomain() []byte {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *Site) GetCategories() []uint32 {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Site) GetUrl() []byte {
	if m != nil {
		return m.Url
	}
	return nil
}

// 浏览器
type Browser struct {
	UserAgent        *string `protobuf:"bytes,1,opt,name=user_agent" json:"user_agent,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Browser) Reset()         { *m = Browser{} }
func (m *Browser) String() string { return proto.CompactTextString(m) }
func (*Browser) ProtoMessage()    {}

func (m *Browser) GetUserAgent() string {
	if m != nil && m.UserAgent != nil {
		return *m.UserAgent
	}
	return ""
}

// 媒体信息
type Media struct {
	Id               *string    `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	ChannelId        *string    `protobuf:"bytes,2,req,name=channel_id" json:"channel_id,omitempty"`
	Type             *MediaType `protobuf:"varint,3,req,name=type,enum=jesgoo.protocol.MediaType" json:"type,omitempty"`
	App              *App       `protobuf:"bytes,4,opt,name=app" json:"app,omitempty"`
	Site             *Site      `protobuf:"bytes,5,opt,name=site" json:"site,omitempty"`
	Browser          *Browser   `protobuf:"bytes,6,opt,name=browser" json:"browser,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Media) Reset()         { *m = Media{} }
func (m *Media) String() string { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()    {}

func (m *Media) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Media) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

func (m *Media) GetType() MediaType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MediaType_APP
}

func (m *Media) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Media) GetSite() *Site {
	if m != nil {
		return m.Site
	}
	return nil
}

func (m *Media) GetBrowser() *Browser {
	if m != nil {
		return m.Browser
	}
	return nil
}

// WiFi热点
type WiFi struct {
	Mac              *string `protobuf:"bytes,1,req,name=mac" json:"mac,omitempty"`
	Rssi             *int32  `protobuf:"varint,2,req,name=rssi" json:"rssi,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WiFi) Reset()         { *m = WiFi{} }
func (m *WiFi) String() string { return proto.CompactTextString(m) }
func (*WiFi) ProtoMessage()    {}

func (m *WiFi) GetMac() string {
	if m != nil && m.Mac != nil {
		return *m.Mac
	}
	return ""
}

func (m *WiFi) GetRssi() int32 {
	if m != nil && m.Rssi != nil {
		return *m.Rssi
	}
	return 0
}

// 网络环境
type Network struct {
	Ip               *string      `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Type             *NetworkType `protobuf:"varint,2,opt,name=type,enum=jesgoo.protocol.NetworkType" json:"type,omitempty"`
	CellularOperator *uint32      `protobuf:"varint,3,opt,name=cellular_operator" json:"cellular_operator,omitempty"`
	CellularId       *string      `protobuf:"bytes,4,opt,name=cellular_id" json:"cellular_id,omitempty"`
	Wifis            []*WiFi      `protobuf:"bytes,5,rep,name=wifis" json:"wifis,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}

func (m *Network) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Network) GetType() NetworkType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return NetworkType_WIFI
}

func (m *Network) GetCellularOperator() uint32 {
	if m != nil && m.CellularOperator != nil {
		return *m.CellularOperator
	}
	return 0
}

func (m *Network) GetCellularId() string {
	if m != nil && m.CellularId != nil {
		return *m.CellularId
	}
	return ""
}

func (m *Network) GetWifis() []*WiFi {
	if m != nil {
		return m.Wifis
	}
	return nil
}

// 设备ID
// 精简ID针对有辅助字符的ID类型，如MAC和IDFA
type DeviceID struct {
	Type             *DeviceIDType `protobuf:"varint,1,req,name=type,enum=jesgoo.protocol.DeviceIDType" json:"type,omitempty"`
	Id               []byte        `protobuf:"bytes,2,req,name=id" json:"id,omitempty"`
	Compact          *bool         `protobuf:"varint,3,opt,name=compact,def=0" json:"compact,omitempty"`
	Md5              *bool         `protobuf:"varint,4,opt,name=md5,def=0" json:"md5,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *DeviceID) Reset()         { *m = DeviceID{} }
func (m *DeviceID) String() string { return proto.CompactTextString(m) }
func (*DeviceID) ProtoMessage()    {}

const Default_DeviceID_Compact bool = false
const Default_DeviceID_Md5 bool = false

func (m *DeviceID) GetType() DeviceIDType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return DeviceIDType_IMEI
}

func (m *DeviceID) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DeviceID) GetCompact() bool {
	if m != nil && m.Compact != nil {
		return *m.Compact
	}
	return Default_DeviceID_Compact
}

func (m *DeviceID) GetMd5() bool {
	if m != nil && m.Md5 != nil {
		return *m.Md5
	}
	return Default_DeviceID_Md5
}

// 设备信息
type Device struct {
	Type             *DeviceType `protobuf:"varint,1,req,name=type,enum=jesgoo.protocol.DeviceType" json:"type,omitempty"`
	Ids              []*DeviceID `protobuf:"bytes,2,rep,name=ids" json:"ids,omitempty"`
	OsType           *OSType     `protobuf:"varint,3,req,name=os_type,enum=jesgoo.protocol.OSType" json:"os_type,omitempty"`
	OsVersion        *Version    `protobuf:"bytes,4,req,name=os_version" json:"os_version,omitempty"`
	Brand            *string     `protobuf:"bytes,5,opt,name=brand" json:"brand,omitempty"`
	Model            *string     `protobuf:"bytes,6,opt,name=model" json:"model,omitempty"`
	ScreenSize       *Size       `protobuf:"bytes,7,opt,name=screen_size" json:"screen_size,omitempty"`
	ScreenDensity    *float64    `protobuf:"fixed64,8,opt,name=screen_density" json:"screen_density,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}

func (m *Device) GetType() DeviceType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return DeviceType_DEV_UNKNOWN
}

func (m *Device) GetIds() []*DeviceID {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Device) GetOsType() OSType {
	if m != nil && m.OsType != nil {
		return *m.OsType
	}
	return OSType_OS_UNKNOWN
}

func (m *Device) GetOsVersion() *Version {
	if m != nil {
		return m.OsVersion
	}
	return nil
}

func (m *Device) GetBrand() string {
	if m != nil && m.Brand != nil {
		return *m.Brand
	}
	return ""
}

func (m *Device) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *Device) GetScreenSize() *Size {
	if m != nil {
		return m.ScreenSize
	}
	return nil
}

func (m *Device) GetScreenDensity() float64 {
	if m != nil && m.ScreenDensity != nil {
		return *m.ScreenDensity
	}
	return 0
}

// 设备
type Geo struct {
	Type             *GeoType   `protobuf:"varint,1,req,name=type,enum=jesgoo.protocol.GeoType" json:"type,omitempty"`
	Longitude        *float64   `protobuf:"fixed64,2,req,name=longitude" json:"longitude,omitempty"`
	Latitude         *float64   `protobuf:"fixed64,3,req,name=latitude" json:"latitude,omitempty"`
	Timestampe       *uint32    `protobuf:"varint,4,opt,name=timestampe" json:"timestampe,omitempty"`
	Source           *GeoSource `protobuf:"varint,5,opt,name=source,enum=jesgoo.protocol.GeoSource" json:"source,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Geo) Reset()         { *m = Geo{} }
func (m *Geo) String() string { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()    {}

func (m *Geo) GetType() GeoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return GeoType_WGS84
}

func (m *Geo) GetLongitude() float64 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *Geo) GetLatitude() float64 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *Geo) GetTimestampe() uint32 {
	if m != nil && m.Timestampe != nil {
		return *m.Timestampe
	}
	return 0
}

func (m *Geo) GetSource() GeoSource {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return GeoSource_NATIVE
}

// 广告客户端
type Client struct {
	Type             *ClientType `protobuf:"varint,1,req,name=type,enum=jesgoo.protocol.ClientType" json:"type,omitempty"`
	Version          *Version    `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}

func (m *Client) GetType() ClientType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ClientType_NATIVESDK
}

func (m *Client) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

// 广告位信息
type AdSlot struct {
	Id               *string         `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Type             *AdSlotType     `protobuf:"varint,2,req,name=type,enum=jesgoo.protocol.AdSlotType" json:"type,omitempty"`
	Size             *Size           `protobuf:"bytes,3,req,name=size" json:"size,omitempty"`
	Styles           []AdStyle       `protobuf:"varint,4,rep,name=styles,enum=jesgoo.protocol.AdStyle" json:"styles,omitempty"`
	Capacity         *uint32         `protobuf:"varint,5,opt,name=capacity,def=1" json:"capacity,omitempty"`
	Templates        []uint32        `protobuf:"varint,6,rep,name=templates" json:"templates,omitempty"`
	Promotions       []PromotionType `protobuf:"varint,7,rep,name=promotions,enum=jesgoo.protocol.PromotionType" json:"promotions,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *AdSlot) Reset()         { *m = AdSlot{} }
func (m *AdSlot) String() string { return proto.CompactTextString(m) }
func (*AdSlot) ProtoMessage()    {}

const Default_AdSlot_Capacity uint32 = 1

func (m *AdSlot) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *AdSlot) GetType() AdSlotType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return AdSlotType_BANNER
}

func (m *AdSlot) GetSize() *Size {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *AdSlot) GetStyles() []AdStyle {
	if m != nil {
		return m.Styles
	}
	return nil
}

func (m *AdSlot) GetCapacity() uint32 {
	if m != nil && m.Capacity != nil {
		return *m.Capacity
	}
	return Default_AdSlot_Capacity
}

func (m *AdSlot) GetTemplates() []uint32 {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *AdSlot) GetPromotions() []PromotionType {
	if m != nil {
		return m.Promotions
	}
	return nil
}

// 广告原生信息
type AdNativeMaterial struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Title            *string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Description1     *string `protobuf:"bytes,3,opt,name=description1" json:"description1,omitempty"`
	Description2     *string `protobuf:"bytes,4,opt,name=description2" json:"description2,omitempty"`
	ImageUrl         *string `protobuf:"bytes,5,opt,name=image_url" json:"image_url,omitempty"`
	ImageSize        *Size   `protobuf:"bytes,6,opt,name=image_size" json:"image_size,omitempty"`
	LogoUrl          *string `protobuf:"bytes,7,opt,name=logo_url" json:"logo_url,omitempty"`
	LogoSize         *Size   `protobuf:"bytes,8,opt,name=logo_size" json:"logo_size,omitempty"`
	ClickUrl         *string `protobuf:"bytes,9,opt,name=click_url" json:"click_url,omitempty"`
	ImpressionLogUrl *string `protobuf:"bytes,10,opt,name=impression_log_url" json:"impression_log_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AdNativeMaterial) Reset()         { *m = AdNativeMaterial{} }
func (m *AdNativeMaterial) String() string { return proto.CompactTextString(m) }
func (*AdNativeMaterial) ProtoMessage()    {}

func (m *AdNativeMaterial) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *AdNativeMaterial) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *AdNativeMaterial) GetDescription1() string {
	if m != nil && m.Description1 != nil {
		return *m.Description1
	}
	return ""
}

func (m *AdNativeMaterial) GetDescription2() string {
	if m != nil && m.Description2 != nil {
		return *m.Description2
	}
	return ""
}

func (m *AdNativeMaterial) GetImageUrl() string {
	if m != nil && m.ImageUrl != nil {
		return *m.ImageUrl
	}
	return ""
}

func (m *AdNativeMaterial) GetImageSize() *Size {
	if m != nil {
		return m.ImageSize
	}
	return nil
}

func (m *AdNativeMaterial) GetLogoUrl() string {
	if m != nil && m.LogoUrl != nil {
		return *m.LogoUrl
	}
	return ""
}

func (m *AdNativeMaterial) GetLogoSize() *Size {
	if m != nil {
		return m.LogoSize
	}
	return nil
}

func (m *AdNativeMaterial) GetClickUrl() string {
	if m != nil && m.ClickUrl != nil {
		return *m.ClickUrl
	}
	return ""
}

func (m *AdNativeMaterial) GetImpressionLogUrl() string {
	if m != nil && m.ImpressionLogUrl != nil {
		return *m.ImpressionLogUrl
	}
	return ""
}

// 卖方请求
type SellerRequest struct {
	Media   *Media    `protobuf:"bytes,1,req,name=media" json:"media,omitempty"`
	Device  *Device   `protobuf:"bytes,2,req,name=device" json:"device,omitempty"`
	Network *Network  `protobuf:"bytes,3,req,name=network" json:"network,omitempty"`
	Client  *Client   `protobuf:"bytes,4,req,name=client" json:"client,omitempty"`
	Geo     *Geo      `protobuf:"bytes,5,opt,name=geo" json:"geo,omitempty"`
	Adslots []*AdSlot `protobuf:"bytes,6,rep,name=adslots" json:"adslots,omitempty"`
	// 特殊标记字段，从100开始编号
	Debug            *bool  `protobuf:"varint,101,opt,name=debug" json:"debug,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SellerRequest) Reset()         { *m = SellerRequest{} }
func (m *SellerRequest) String() string { return proto.CompactTextString(m) }
func (*SellerRequest) ProtoMessage()    {}

func (m *SellerRequest) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func (m *SellerRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *SellerRequest) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *SellerRequest) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *SellerRequest) GetGeo() *Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *SellerRequest) GetAdslots() []*AdSlot {
	if m != nil {
		return m.Adslots
	}
	return nil
}

func (m *SellerRequest) GetDebug() bool {
	if m != nil && m.Debug != nil {
		return *m.Debug
	}
	return false
}

// 卖方应答
type SellerResponse struct {
	Success          *bool                `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	Ads              []*SellerResponse_Ad `protobuf:"bytes,2,rep,name=ads" json:"ads,omitempty"`
	SearchId         *string              `protobuf:"bytes,3,opt,name=search_id" json:"search_id,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SellerResponse) Reset()         { *m = SellerResponse{} }
func (m *SellerResponse) String() string { return proto.CompactTextString(m) }
func (*SellerResponse) ProtoMessage()    {}

func (m *SellerResponse) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *SellerResponse) GetAds() []*SellerResponse_Ad {
	if m != nil {
		return m.Ads
	}
	return nil
}

func (m *SellerResponse) GetSearchId() string {
	if m != nil && m.SearchId != nil {
		return *m.SearchId
	}
	return ""
}

// 广告内容
type SellerResponse_Ad struct {
	AdslotId         *string                         `protobuf:"bytes,1,req,name=adslot_id" json:"adslot_id,omitempty"`
	MaterialType     *SellerResponse_Ad_MaterialType `protobuf:"varint,2,req,name=material_type,enum=jesgoo.protocol.SellerResponse_Ad_MaterialType" json:"material_type,omitempty"`
	HtmlSnippet      []byte                          `protobuf:"bytes,3,opt,name=html_snippet" json:"html_snippet,omitempty"`
	NativeMaterial   *AdNativeMaterial               `protobuf:"bytes,4,opt,name=native_material" json:"native_material,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *SellerResponse_Ad) Reset()         { *m = SellerResponse_Ad{} }
func (m *SellerResponse_Ad) String() string { return proto.CompactTextString(m) }
func (*SellerResponse_Ad) ProtoMessage()    {}

func (m *SellerResponse_Ad) GetAdslotId() string {
	if m != nil && m.AdslotId != nil {
		return *m.AdslotId
	}
	return ""
}

func (m *SellerResponse_Ad) GetMaterialType() SellerResponse_Ad_MaterialType {
	if m != nil && m.MaterialType != nil {
		return *m.MaterialType
	}
	return SellerResponse_Ad_DYNAMIC
}

func (m *SellerResponse_Ad) GetHtmlSnippet() []byte {
	if m != nil {
		return m.HtmlSnippet
	}
	return nil
}

func (m *SellerResponse_Ad) GetNativeMaterial() *AdNativeMaterial {
	if m != nil {
		return m.NativeMaterial
	}
	return nil
}

func init() {
	proto.RegisterEnum("jesgoo.protocol.DeviceType", DeviceType_name, DeviceType_value)
	proto.RegisterEnum("jesgoo.protocol.MediaType", MediaType_name, MediaType_value)
	proto.RegisterEnum("jesgoo.protocol.OSType", OSType_name, OSType_value)
	proto.RegisterEnum("jesgoo.protocol.DeviceIDType", DeviceIDType_name, DeviceIDType_value)
	proto.RegisterEnum("jesgoo.protocol.AdSlotType", AdSlotType_name, AdSlotType_value)
	proto.RegisterEnum("jesgoo.protocol.AdStyle", AdStyle_name, AdStyle_value)
	proto.RegisterEnum("jesgoo.protocol.PromotionType", PromotionType_name, PromotionType_value)
	proto.RegisterEnum("jesgoo.protocol.NetworkType", NetworkType_name, NetworkType_value)
	proto.RegisterEnum("jesgoo.protocol.GeoType", GeoType_name, GeoType_value)
	proto.RegisterEnum("jesgoo.protocol.ClientType", ClientType_name, ClientType_value)
	proto.RegisterEnum("jesgoo.protocol.GeoSource", GeoSource_name, GeoSource_value)
	proto.RegisterEnum("jesgoo.protocol.SellerResponse_Ad_MaterialType", SellerResponse_Ad_MaterialType_name, SellerResponse_Ad_MaterialType_value)
}
