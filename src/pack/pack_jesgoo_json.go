package pack

import (
	"context"
	"encoding/json"
	"strconv"
	"utils"
)

type PackJesgooResponseJsonModule struct {
}

type Ad struct {
	Adslot_id       string
	Material_type   int
	Html_snippet    string
	Native_material NativeMaterial
}

type NativeMaterial struct {
	Id                 string
	Type               int
	Interaction_type   int
	Title              string
	Description1       string
	Description2       string
	Image_url          string
	Logo_url           string
	Click_url          string
	Impression_log_url []string
}

type SellerReponse struct {
	Success   bool
	Ads       []Ad
	Search_id string
}

func (this *PackJesgooResponseJsonModule) Init(inner_data *context.GlobalContext) (err error) {
	return
}

func (this *PackJesgooResponseJsonModule) pack_native_ad(ad *Ad, inner_ad *context.AdInfo) (err error) {
	admaterial := &ad.Native_material
	switch inner_ad.AdType {
	case context.TEXT:
		admaterial.Type = 0
	case context.IMAGE:
		admaterial.Type = 1
	case context.TEXT_ICON:
		admaterial.Type = 2
	default:
		admaterial.Type = 0
	}
	switch inner_ad.InteractionType {
	case context.NO_INTERACT:
		admaterial.Interaction_type = 0
	case context.SURFING:
		admaterial.Interaction_type = 1
	case context.DOWNLOAD:
		admaterial.Interaction_type = 2
	case context.DIALING:
		admaterial.Interaction_type = 3
	case context.MESSAGE:
		admaterial.Interaction_type = 4
	case context.MAIL:
		admaterial.Interaction_type = 5
	}
	admaterial.Id = strconv.Itoa(int(inner_ad.Adid))
	admaterial.Title = inner_ad.Title
	admaterial.Image_url = inner_ad.ImageUrl
	admaterial.Description1 = inner_ad.Description1
	admaterial.Description2 = inner_ad.Description2
	admaterial.Click_url = inner_ad.ClickUrl
	for i := 0; i < len(inner_ad.ImpressionUrl); i++ {
		admaterial.Impression_log_url = append(admaterial.Impression_log_url, inner_ad.ImpressionUrl[i])
	}
	utils.DebugLog.Write("fill kaiping ad success")
	return
}

func (this *PackJesgooResponseJsonModule) Run(inner_data *context.Context) (err error) {
	var temp_resp SellerReponse
	temp_resp.Success = true
	temp_resp.Search_id = inner_data.Searchid
	need_ad := inner_data.Req.AdSlot.Capacity
	ad_num := len(inner_data.Resp.Ads)
	var pack_num int32
	if int32(need_ad) < int32(ad_num) {
		pack_num = int32(need_ad)
	} else {
		pack_num = int32(ad_num)
	}
	utils.DebugLog.Write("pack_num is %d, need_ad %d, ad_num %d", pack_num, need_ad, ad_num)
	var i int32
	for i = 0; i < pack_num; i++ {
		var temp_ad Ad
		temp_ad.Adslot_id = inner_data.Req.AdSlot.Slotid
		switch inner_data.Req.AdSlot.AdSlotType {
		case context.AdSlotType_BANNER:
			temp_ad.Html_snippet = inner_data.Resp.Ads[i].HtmlSnippet.String()
			temp_ad.Material_type = 0
		case context.AdSlotType_INITIALIZATION:
			err = this.pack_native_ad(&temp_ad, &inner_data.Resp.Ads[i])
			temp_ad.Material_type = 1
		case context.AdSlotType_INSERT:
			err = this.pack_native_ad(&temp_ad, &inner_data.Resp.Ads[i])
			temp_ad.Material_type = 1
		default:
			temp_ad.Html_snippet = inner_data.Resp.Ads[i].HtmlSnippet.String()
			temp_ad.Material_type = 0
		}
		temp_resp.Ads = append(temp_resp.Ads, temp_ad)
	}
	utils.DebugLog.Write("json resp [%s]", temp_resp)
	inner_data.RespBody, err = json.Marshal(temp_resp)
	return
}
