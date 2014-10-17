package pack

import (
	"code.google.com/p/goprotobuf/proto"
	"context"
	"jesgoo_interface"
	"utils"
)

type PackJesgooResponseModule struct {
}

func (this PackJesgooResponseModule) Run(inner_data *context.Context) (err error) {

	var temp_response jesgoo_interface.SellerResponse
	temp_response.Success = new(bool)
	*temp_response.Success = true
	temp_response.SearchId = new(string)
	*temp_response.SearchId = inner_data.Searchid
	temp_response.Ads = make([]*jesgoo_interface.SellerResponse_Ad, 0)
	need_ad := inner_data.Req.AdSlot.Capacity
	ad_num := len(inner_data.Resp.Ads)
	var pack_num int32
	if int32(need_ad) < int32(ad_num) {
		pack_num = int32(need_ad)
	} else {
		pack_num = int32(ad_num)
	}
	utils.DebugLog.Write("pack_num is %d", pack_num)
	var i int32
	for i = 0; i < pack_num; i++ {
		var temp_ad jesgoo_interface.SellerResponse_Ad
		temp_ad.AdslotId = new(string)
		*temp_ad.AdslotId = inner_data.Req.AdSlot.Slotid
		temp_ad.MaterialType = new(jesgoo_interface.SellerResponse_Ad_MaterialType)
		*temp_ad.MaterialType = jesgoo_interface.SellerResponse_Ad_DYNAMIC
		temp_ad.HtmlSnippet = make([]byte, 0)
		temp_ad.HtmlSnippet = inner_data.Resp.Ads[i].HtmlSnippet.Bytes()
		temp_response.Ads = append(temp_response.Ads, &temp_ad)
	}
	inner_data.RespBody, err = proto.Marshal(&temp_response)
	return
}

func (this PackJesgooResponseModule) Init(inner_data *context.GlobalContext) (err error) {
	return
}
