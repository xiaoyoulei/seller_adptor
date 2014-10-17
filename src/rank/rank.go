package rank

import (
	"context"
	"utils"
)

type RankModule struct {
}

func (this *RankModule) Init(inner_data *context.GlobalContext) (err error) {
	return

}

func (this *RankModule) Run(inner_data *context.Context) (err error) {
	inner_resp := &inner_data.Resp
	inner_resp.Ads = make([]context.AdInfo, 0)
	utils.DebugLog.Write("JesgooAds Num is %d", len(inner_data.JesgooAds))
	for i := 0; i < len(inner_data.JesgooAds); i++ {
		inner_resp.Ads = append(inner_resp.Ads, inner_data.JesgooAds[i])
	}
	utils.DebugLog.Write("BaiduAds Num is %d", len(inner_data.BaiduAds))
	for i := 0; i < len(inner_data.BaiduAds); i++ {
		inner_resp.Ads = append(inner_resp.Ads, inner_data.BaiduAds[i])
	}
	return
}
