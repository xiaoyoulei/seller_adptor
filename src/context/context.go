package context

type AppInfo struct {
	Package_name string
}
type MediaInfo struct {
	Appsid     string
	Channel_id string
	App        AppInfo
}
type SizeInfo struct {
	Width  int
	Height int
}
type AdSlotInfo struct {
	Slotid string
	Size   SizeInfo
}

type NetworkInfo struct {
	Ip           string
	Network_type int
}
type InnerReq struct {
	Media_info MediaInfo
	AdSlot_t   AdSlotInfo
	Network_t  NetworkInfo
}

type AdInfo struct {
	Adid           string
	Title          string
	Description1   string
	Description2   string
	Image_url      string
	Image_size     SizeInfo
	Logo_url       string
	Logo_size      SizeInfo
	Click_url      string
	Impression_url string
}
type InnerResp struct {
	Ads []AdInfo
}

type Context struct {
	Searchid string
	Req      InnerReq
	Resp     InnerResp
}
