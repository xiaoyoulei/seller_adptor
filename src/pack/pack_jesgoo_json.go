package pack

import (
	"context"
	"encoding/json"
	"log"
)

type PackJesgooResponseJsonModule struct {
}

type Ad struct {
	Adslot_id     string
	Material_type int
	Html_snippet  string
}

type SellerReponse struct {
	Success   bool
	Ads       []Ad
	Search_id string
}

func (this PackJesgooResponseJsonModule) Init(inner_data *context.GlobalContext) (err error) {
	return
}

func (this PackJesgooResponseJsonModule) Run(inner_data *context.Context) (err error) {
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
	log.Printf("pack_num is %d, need_ad %d, ad_num %d", pack_num, need_ad, ad_num)
	var i int32
	for i = 0; i < pack_num; i++ {
		var temp_ad Ad
		temp_ad.Html_snippet = inner_data.Resp.Ads[i].HtmlSnippet.String()
		temp_resp.Ads = append(temp_resp.Ads, temp_ad)
	}
	inner_data.RespBody, err = json.Marshal(temp_resp)
	return
}
