package pack

import (
	"code.google.com/p/goprotobuf/proto"
	"context"
	"jesgoo_interface"
	"strconv"
	"utils"
)

type PackJesgooResponseModule struct {
}

func (this *PackJesgooResponseModule) pack_native_ad(ad *jesgoo_interface.SellerResponse_Ad, inner_ad *context.AdInfo) (err error) {
	ad.NativeMaterial = new(jesgoo_interface.AdNativeMaterial)
	admaterial := ad.NativeMaterial
	admaterial.Id = new(string)
	*admaterial.Id = strconv.Itoa(int(inner_ad.Adid))
	if len(inner_ad.Title) > 0 {
		admaterial.Title = new(string)
		*admaterial.Title = inner_ad.Title
	}
	if len(inner_ad.Description1) > 0 {
		admaterial.Description1 = new(string)
		*admaterial.Description1 = inner_ad.Description1
	}
	if len(inner_ad.Description2) > 0 {
		admaterial.Description2 = new(string)
		*admaterial.Description2 = inner_ad.Description1
	}
	if len(inner_ad.ImageUrl) > 0 {
		admaterial.ImageUrl = new(string)
		*admaterial.ImageUrl = inner_ad.ImageUrl
	}
	if len(inner_ad.LogoUrl) > 0 {
		admaterial.LogoUrl = new(string)
		*admaterial.LogoUrl = inner_ad.LogoUrl
	}
	admaterial.ClickUrl = new(string)
	*admaterial.ClickUrl = inner_ad.ClickUrl
	admaterial.ImpressionLogUrl = make([]string, 0)
	for i := 0; i < len(inner_ad.ImpressionUrl); i++ {
		admaterial.ImpressionLogUrl = append(admaterial.ImpressionLogUrl, inner_ad.ImpressionUrl[i])
	}

	return
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
		switch inner_data.Req.AdSlot.AdSlotType {
		case context.AdSlotType_BANNER:
			*temp_ad.MaterialType = jesgoo_interface.SellerResponse_Ad_DYNAMIC
			temp_ad.HtmlSnippet = make([]byte, 0)
			temp_ad.HtmlSnippet = inner_data.Resp.Ads[i].HtmlSnippet.Bytes()
		case context.AdSlotType_INITIALIZATION:
			*temp_ad.MaterialType = jesgoo_interface.SellerResponse_Ad_NATIVE
			err = this.pack_native_ad(&temp_ad, &inner_data.Resp.Ads[i])
		}
		temp_response.Ads = append(temp_response.Ads, &temp_ad)
	}
	inner_data.RespBody, err = proto.Marshal(&temp_response)
	return
}

func (this PackJesgooResponseModule) Init(inner_data *context.GlobalContext) (err error) {
	return
}
