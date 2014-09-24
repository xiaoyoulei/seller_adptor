# It's the interface between UI && BS
# all rights reserved

namespace cpp internal

struct BSRequest {
	1: optional string searchid,
	2: optional string userid,
	3: optional string sequence_id,
	4: optional string cookie,
	5: optional string ad_block_id,
	6: optional i32	width,
	7: optional i32 height,
	8: optional string os,
	9: optional string osv,
	10:optional string ip
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
	void search2(),
	BSResponse search(1:BSRequest req)
}
