// Code generated by protoc-gen-go.
// source: uilog.proto
// DO NOT EDIT!

/*
Package jesgoo_uilog is a generated protocol buffer package.

It is generated from these files:
	uilog.proto

It has these top-level messages:
	AppInfo
	Size
	Media
	Adslot
	Network
	Version
	DeviceId
	Device
	AdInfo
	AdDspRet
	NoticeLogBody
*/
package jesgoo_uilog

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MediaType int32

const (
	MediaType_UNKNOWN MediaType = 0
	MediaType_APP     MediaType = 1
	MediaType_WEB     MediaType = 2
	MediaType_WAP     MediaType = 3
)

var MediaType_name = map[int32]string{
	0: "UNKNOWN",
	1: "APP",
	2: "WEB",
	3: "WAP",
}
var MediaType_value = map[string]int32{
	"UNKNOWN": 0,
	"APP":     1,
	"WEB":     2,
	"WAP":     3,
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

type AdslotType int32

const (
	AdslotType_BANNER         AdslotType = 1
	AdslotType_OFFERWALL      AdslotType = 2
	AdslotType_RECOMMEND      AdslotType = 3
	AdslotType_INITIALIZATION AdslotType = 4
)

var AdslotType_name = map[int32]string{
	1: "BANNER",
	2: "OFFERWALL",
	3: "RECOMMEND",
	4: "INITIALIZATION",
}
var AdslotType_value = map[string]int32{
	"BANNER":         1,
	"OFFERWALL":      2,
	"RECOMMEND":      3,
	"INITIALIZATION": 4,
}

func (x AdslotType) Enum() *AdslotType {
	p := new(AdslotType)
	*p = x
	return p
}
func (x AdslotType) String() string {
	return proto.EnumName(AdslotType_name, int32(x))
}
func (x *AdslotType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AdslotType_value, data, "AdslotType")
	if err != nil {
		return err
	}
	*x = AdslotType(value)
	return nil
}

type NetworkType int32

const (
	NetworkType_NET_UNKNOWN NetworkType = 0
	NetworkType_NET_WIFI    NetworkType = 1
	NetworkType_NET_2G      NetworkType = 2
	NetworkType_NET_3G      NetworkType = 3
	NetworkType_NET_4G      NetworkType = 4
)

var NetworkType_name = map[int32]string{
	0: "NET_UNKNOWN",
	1: "NET_WIFI",
	2: "NET_2G",
	3: "NET_3G",
	4: "NET_4G",
}
var NetworkType_value = map[string]int32{
	"NET_UNKNOWN": 0,
	"NET_WIFI":    1,
	"NET_2G":      2,
	"NET_3G":      3,
	"NET_4G":      4,
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

type OSType int32

const (
	OSType_OS_UNKNOWN OSType = 0
	OSType_OS_ANDROID OSType = 1
	OSType_OS_IOS     OSType = 2
	OSType_OS_WP      OSType = 3
)

var OSType_name = map[int32]string{
	0: "OS_UNKNOWN",
	1: "OS_ANDROID",
	2: "OS_IOS",
	3: "OS_WP",
}
var OSType_value = map[string]int32{
	"OS_UNKNOWN": 0,
	"OS_ANDROID": 1,
	"OS_IOS":     2,
	"OS_WP":      3,
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

type DeviceIdType int32

const (
	DeviceIdType_DEVID_UNKNOWN DeviceIdType = 0
	DeviceIdType_DEVID_IMEI    DeviceIdType = 1
	DeviceIdType_DEVID_MAC     DeviceIdType = 2
	DeviceIdType_DEVID_IDFA    DeviceIdType = 3
)

var DeviceIdType_name = map[int32]string{
	0: "DEVID_UNKNOWN",
	1: "DEVID_IMEI",
	2: "DEVID_MAC",
	3: "DEVID_IDFA",
}
var DeviceIdType_value = map[string]int32{
	"DEVID_UNKNOWN": 0,
	"DEVID_IMEI":    1,
	"DEVID_MAC":     2,
	"DEVID_IDFA":    3,
}

func (x DeviceIdType) Enum() *DeviceIdType {
	p := new(DeviceIdType)
	*p = x
	return p
}
func (x DeviceIdType) String() string {
	return proto.EnumName(DeviceIdType_name, int32(x))
}
func (x *DeviceIdType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeviceIdType_value, data, "DeviceIdType")
	if err != nil {
		return err
	}
	*x = DeviceIdType(value)
	return nil
}

type AdType int32

const (
	AdType_TEXT      AdType = 1
	AdType_IMAGE     AdType = 2
	AdType_HTML      AdType = 3
	AdType_VIDEO     AdType = 4
	AdType_TEXT_ICON AdType = 5
)

var AdType_name = map[int32]string{
	1: "TEXT",
	2: "IMAGE",
	3: "HTML",
	4: "VIDEO",
	5: "TEXT_ICON",
}
var AdType_value = map[string]int32{
	"TEXT":      1,
	"IMAGE":     2,
	"HTML":      3,
	"VIDEO":     4,
	"TEXT_ICON": 5,
}

func (x AdType) Enum() *AdType {
	p := new(AdType)
	*p = x
	return p
}
func (x AdType) String() string {
	return proto.EnumName(AdType_name, int32(x))
}
func (x *AdType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AdType_value, data, "AdType")
	if err != nil {
		return err
	}
	*x = AdType(value)
	return nil
}

type AdSrc int32

const (
	AdSrc_JESGOO AdSrc = 1
	AdSrc_BAIDU  AdSrc = 2
)

var AdSrc_name = map[int32]string{
	1: "JESGOO",
	2: "BAIDU",
}
var AdSrc_value = map[string]int32{
	"JESGOO": 1,
	"BAIDU":  2,
}

func (x AdSrc) Enum() *AdSrc {
	p := new(AdSrc)
	*p = x
	return p
}
func (x AdSrc) String() string {
	return proto.EnumName(AdSrc_name, int32(x))
}
func (x *AdSrc) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AdSrc_value, data, "AdSrc")
	if err != nil {
		return err
	}
	*x = AdSrc(value)
	return nil
}

type InteractionType int32

const (
	InteractionType_NO_INTERACT InteractionType = 0
	InteractionType_SURFING     InteractionType = 1
	InteractionType_DOWNLOAD    InteractionType = 2
	InteractionType_DIALING     InteractionType = 3
	InteractionType_MESSAGE     InteractionType = 4
	InteractionType_MAIL        InteractionType = 5
)

var InteractionType_name = map[int32]string{
	0: "NO_INTERACT",
	1: "SURFING",
	2: "DOWNLOAD",
	3: "DIALING",
	4: "MESSAGE",
	5: "MAIL",
}
var InteractionType_value = map[string]int32{
	"NO_INTERACT": 0,
	"SURFING":     1,
	"DOWNLOAD":    2,
	"DIALING":     3,
	"MESSAGE":     4,
	"MAIL":        5,
}

func (x InteractionType) Enum() *InteractionType {
	p := new(InteractionType)
	*p = x
	return p
}
func (x InteractionType) String() string {
	return proto.EnumName(InteractionType_name, int32(x))
}
func (x *InteractionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InteractionType_value, data, "InteractionType")
	if err != nil {
		return err
	}
	*x = InteractionType(value)
	return nil
}

type AppInfo struct {
	Packagename      *string `protobuf:"bytes,1,opt,name=packagename" json:"packagename,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AppInfo) Reset()         { *m = AppInfo{} }
func (m *AppInfo) String() string { return proto.CompactTextString(m) }
func (*AppInfo) ProtoMessage()    {}

func (m *AppInfo) GetPackagename() string {
	if m != nil && m.Packagename != nil {
		return *m.Packagename
	}
	return ""
}

type Size struct {
	Height           *uint32 `protobuf:"varint,1,req,name=height" json:"height,omitempty"`
	Width            *uint32 `protobuf:"varint,2,req,name=width" json:"width,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Size) Reset()         { *m = Size{} }
func (m *Size) String() string { return proto.CompactTextString(m) }
func (*Size) ProtoMessage()    {}

func (m *Size) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *Size) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

type Media struct {
	Type             *MediaType `protobuf:"varint,1,req,name=type,enum=jesgoo.uilog.MediaType" json:"type,omitempty"`
	Appsid           *string    `protobuf:"bytes,2,req,name=appsid" json:"appsid,omitempty"`
	Channelid        *string    `protobuf:"bytes,3,opt,name=channelid" json:"channelid,omitempty"`
	App              *AppInfo   `protobuf:"bytes,4,opt,name=app" json:"app,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Media) Reset()         { *m = Media{} }
func (m *Media) String() string { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()    {}

func (m *Media) GetType() MediaType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MediaType_UNKNOWN
}

func (m *Media) GetAppsid() string {
	if m != nil && m.Appsid != nil {
		return *m.Appsid
	}
	return ""
}

func (m *Media) GetChannelid() string {
	if m != nil && m.Channelid != nil {
		return *m.Channelid
	}
	return ""
}

func (m *Media) GetApp() *AppInfo {
	if m != nil {
		return m.App
	}
	return nil
}

type Adslot struct {
	Id               *string     `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Type             *AdslotType `protobuf:"varint,2,req,name=type,enum=jesgoo.uilog.AdslotType" json:"type,omitempty"`
	Size             *Size       `protobuf:"bytes,3,req,name=size" json:"size,omitempty"`
	Capacity         *uint32     `protobuf:"varint,4,req,name=capacity" json:"capacity,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Adslot) Reset()         { *m = Adslot{} }
func (m *Adslot) String() string { return proto.CompactTextString(m) }
func (*Adslot) ProtoMessage()    {}

func (m *Adslot) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Adslot) GetType() AdslotType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return AdslotType_BANNER
}

func (m *Adslot) GetSize() *Size {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *Adslot) GetCapacity() uint32 {
	if m != nil && m.Capacity != nil {
		return *m.Capacity
	}
	return 0
}

type Network struct {
	Type             *NetworkType `protobuf:"varint,1,req,name=type,enum=jesgoo.uilog.NetworkType" json:"type,omitempty"`
	Ip               *string      `protobuf:"bytes,2,req,name=ip" json:"ip,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}

func (m *Network) GetType() NetworkType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return NetworkType_NET_UNKNOWN
}

func (m *Network) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

type Version struct {
	Major            *uint32 `protobuf:"varint,1,req,name=major" json:"major,omitempty"`
	Minor            *uint32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Micro            *uint32 `protobuf:"varint,3,opt,name=micro" json:"micro,omitempty"`
	Build            *uint32 `protobuf:"varint,4,opt,name=build" json:"build,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}

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
	return 0
}

func (m *Version) GetMicro() uint32 {
	if m != nil && m.Micro != nil {
		return *m.Micro
	}
	return 0
}

func (m *Version) GetBuild() uint32 {
	if m != nil && m.Build != nil {
		return *m.Build
	}
	return 0
}

type DeviceId struct {
	Type             *DeviceIdType `protobuf:"varint,1,req,name=type,enum=jesgoo.uilog.DeviceIdType" json:"type,omitempty"`
	Id               *string       `protobuf:"bytes,2,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *DeviceId) Reset()         { *m = DeviceId{} }
func (m *DeviceId) String() string { return proto.CompactTextString(m) }
func (*DeviceId) ProtoMessage()    {}

func (m *DeviceId) GetType() DeviceIdType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return DeviceIdType_DEVID_UNKNOWN
}

func (m *DeviceId) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type Device struct {
	Os               *OSType     `protobuf:"varint,1,req,name=os,enum=jesgoo.uilog.OSType" json:"os,omitempty"`
	Osversion        *Version    `protobuf:"bytes,2,req,name=osversion" json:"osversion,omitempty"`
	Ids              []*DeviceId `protobuf:"bytes,3,rep,name=ids" json:"ids,omitempty"`
	Brand            *string     `protobuf:"bytes,4,opt,name=brand" json:"brand,omitempty"`
	Model            *string     `protobuf:"bytes,5,opt,name=model" json:"model,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}

func (m *Device) GetOs() OSType {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return OSType_OS_UNKNOWN
}

func (m *Device) GetOsversion() *Version {
	if m != nil {
		return m.Osversion
	}
	return nil
}

func (m *Device) GetIds() []*DeviceId {
	if m != nil {
		return m.Ids
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

type AdInfo struct {
	Type             *AdType          `protobuf:"varint,1,req,name=type,enum=jesgoo.uilog.AdType" json:"type,omitempty"`
	Src              *AdSrc           `protobuf:"varint,2,req,name=src,enum=jesgoo.uilog.AdSrc" json:"src,omitempty"`
	Interaction      *InteractionType `protobuf:"varint,3,req,name=interaction,enum=jesgoo.uilog.InteractionType" json:"interaction,omitempty"`
	Adid             *uint32          `protobuf:"varint,4,opt,name=adid" json:"adid,omitempty"`
	Groupid          *uint32          `protobuf:"varint,5,opt,name=groupid" json:"groupid,omitempty"`
	Planid           *uint32          `protobuf:"varint,6,opt,name=planid" json:"planid,omitempty"`
	Userid           *uint32          `protobuf:"varint,7,opt,name=userid" json:"userid,omitempty"`
	Bid              *uint32          `protobuf:"varint,8,opt,name=bid" json:"bid,omitempty"`
	Price            *uint32          `protobuf:"varint,9,opt,name=price" json:"price,omitempty"`
	Ctr              *uint64          `protobuf:"varint,10,opt,name=ctr" json:"ctr,omitempty"`
	Cpm              *uint64          `protobuf:"varint,11,opt,name=cpm" json:"cpm,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *AdInfo) Reset()         { *m = AdInfo{} }
func (m *AdInfo) String() string { return proto.CompactTextString(m) }
func (*AdInfo) ProtoMessage()    {}

func (m *AdInfo) GetType() AdType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return AdType_TEXT
}

func (m *AdInfo) GetSrc() AdSrc {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return AdSrc_JESGOO
}

func (m *AdInfo) GetInteraction() InteractionType {
	if m != nil && m.Interaction != nil {
		return *m.Interaction
	}
	return InteractionType_NO_INTERACT
}

func (m *AdInfo) GetAdid() uint32 {
	if m != nil && m.Adid != nil {
		return *m.Adid
	}
	return 0
}

func (m *AdInfo) GetGroupid() uint32 {
	if m != nil && m.Groupid != nil {
		return *m.Groupid
	}
	return 0
}

func (m *AdInfo) GetPlanid() uint32 {
	if m != nil && m.Planid != nil {
		return *m.Planid
	}
	return 0
}

func (m *AdInfo) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *AdInfo) GetBid() uint32 {
	if m != nil && m.Bid != nil {
		return *m.Bid
	}
	return 0
}

func (m *AdInfo) GetPrice() uint32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *AdInfo) GetCtr() uint64 {
	if m != nil && m.Ctr != nil {
		return *m.Ctr
	}
	return 0
}

func (m *AdInfo) GetCpm() uint64 {
	if m != nil && m.Cpm != nil {
		return *m.Cpm
	}
	return 0
}

type AdDspRet struct {
	Src              *AdSrc  `protobuf:"varint,1,req,name=src,enum=jesgoo.uilog.AdSrc" json:"src,omitempty"`
	Adnum            *uint32 `protobuf:"varint,2,req,name=adnum" json:"adnum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AdDspRet) Reset()         { *m = AdDspRet{} }
func (m *AdDspRet) String() string { return proto.CompactTextString(m) }
func (*AdDspRet) ProtoMessage()    {}

func (m *AdDspRet) GetSrc() AdSrc {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return AdSrc_JESGOO
}

func (m *AdDspRet) GetAdnum() uint32 {
	if m != nil && m.Adnum != nil {
		return *m.Adnum
	}
	return 0
}

type NoticeLogBody struct {
	Searchid         *string     `protobuf:"bytes,1,req,name=searchid" json:"searchid,omitempty"`
	Timestamp        *uint32     `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	Media            *Media      `protobuf:"bytes,3,opt,name=media" json:"media,omitempty"`
	Adslot           []*Adslot   `protobuf:"bytes,4,rep,name=adslot" json:"adslot,omitempty"`
	Device           *Device     `protobuf:"bytes,5,opt,name=device" json:"device,omitempty"`
	Ads              []*AdInfo   `protobuf:"bytes,6,rep,name=ads" json:"ads,omitempty"`
	Dspret           []*AdDspRet `protobuf:"bytes,7,rep,name=dspret" json:"dspret,omitempty"`
	Debug            *bool       `protobuf:"varint,8,opt,name=debug" json:"debug,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *NoticeLogBody) Reset()         { *m = NoticeLogBody{} }
func (m *NoticeLogBody) String() string { return proto.CompactTextString(m) }
func (*NoticeLogBody) ProtoMessage()    {}

func (m *NoticeLogBody) GetSearchid() string {
	if m != nil && m.Searchid != nil {
		return *m.Searchid
	}
	return ""
}

func (m *NoticeLogBody) GetTimestamp() uint32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *NoticeLogBody) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func (m *NoticeLogBody) GetAdslot() []*Adslot {
	if m != nil {
		return m.Adslot
	}
	return nil
}

func (m *NoticeLogBody) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *NoticeLogBody) GetAds() []*AdInfo {
	if m != nil {
		return m.Ads
	}
	return nil
}

func (m *NoticeLogBody) GetDspret() []*AdDspRet {
	if m != nil {
		return m.Dspret
	}
	return nil
}

func (m *NoticeLogBody) GetDebug() bool {
	if m != nil && m.Debug != nil {
		return *m.Debug
	}
	return false
}

func init() {
	proto.RegisterEnum("jesgoo.uilog.MediaType", MediaType_name, MediaType_value)
	proto.RegisterEnum("jesgoo.uilog.AdslotType", AdslotType_name, AdslotType_value)
	proto.RegisterEnum("jesgoo.uilog.NetworkType", NetworkType_name, NetworkType_value)
	proto.RegisterEnum("jesgoo.uilog.OSType", OSType_name, OSType_value)
	proto.RegisterEnum("jesgoo.uilog.DeviceIdType", DeviceIdType_name, DeviceIdType_value)
	proto.RegisterEnum("jesgoo.uilog.AdType", AdType_name, AdType_value)
	proto.RegisterEnum("jesgoo.uilog.AdSrc", AdSrc_name, AdSrc_value)
	proto.RegisterEnum("jesgoo.uilog.InteractionType", InteractionType_name, InteractionType_value)
}
