# It's the interface between UI && BS
# all rights reserved

namespace cpp internal

struct Media {
	1: required string appsid,
	2: required string channel_id,
}
struct Size {
	1: i32 width ,
	2: i32 height
}

// 广告位类型
enum AdSlotType {
	BANNER = 1,
	OFFERWALL = 2,
	RECOMMEND = 3,
	INTERSTITIAL = 4
}
// 广告位风格
enum AdSlotStyle {
	TEXT = 1,
	IMAGE = 2,
	ICON_TEXT = 3,
	SMART_AD = 4,
	VIDEO = 5
}
struct AdSlot {
	1: string id,
	2: AdSlotType type,
	3: list<AdSlotStyle> style,
	4: Size size,
	5: optional set<i32> templates,
	6: optional i32 ad_count = 1,
}
// 操作系统类型
enum OSType {
	UNKNOWN = 0,  // 未知或其他系统
	ANDROID = 1,  // 安卓
	IOS = 2,  // iOS
	WP = 3,  // Windows Phone
}

// 设备ID类型
enum DeviceIDType {
	IMEI = 1,
	MAC = 2,
	IDFA = 3,
	AAID = 4,
	OPENUDID = 5,
	ANDROIDID = 6,
	UDID = 7,
	ODIN = 8,
	DUID = 9,
}
// 设备ID
struct DeviceID {
	1: DeviceIDType type,  // 设备ID类型
	2: string id,  // 设备ID
	3: bool compact,  // 是否精简编码（对MAC等有辅助字符的ID生效）
	4: bool md5  // 是否使用MD5签名
}
struct Device {
	1: required OSType os,
	2: optional string osv,
	3: list<DeviceID> dev_id
}
enum NetworkType {
	WIFI = 1,
	CELLULAR_UNKNOWN = 2,
	CELLULAR_2G = 3,
	CELLULAR_3G = 4,
	CELLULAR_4G = 5
}
struct Network {
	1: string ip,
	2: optional NetworkType type,
	3: optional i32 cellular_operator,
	4: optional string cellular_id,
}
struct BSRequest {
	1: required string searchid,
	2: required Media  media,
	3: required AdSlot adslot,
	4: required Device device,
}

struct Ad {
	1: optional i64 adid,
	2: optional i64 groupid,
	3: optional i64 planid,
	4: optional i64 userid,
	5: optional i64 bid,
	6: optional i32 wuliao_type,
	7: optional i32 width,
	8: optional i32 height,
	9: optional string title,
	10:optional string desc,
	11:optional string img_url,
	12:optional string target_url,
	13:optional string show_url,
	14:optional string appname,
	15:optional string app_id,
	16:optional string app_logo,
	17:optional i64 app_size
}
struct BSResponse {
	1: optional i32 res_flag,
	2: optional list<Ad> ads
}

service BSService {
	BSResponse search(1:BSRequest req)
}
