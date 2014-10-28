package rank

import (
	"context"
	"math/rand"
	"sort"
	"utils"
)

type RankModule struct {
}

type sortad []context.AdInfo

func (this sortad) Len() int {
	return len(this)
}

func (this sortad) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// return false means swap(i,j)
func (this sortad) Less(i, j int) bool {
	if this[i].MatchAdSlotType > this[i].MatchAdSlotType {
		return false
	}
	return this[i].Cpm > this[j].Cpm
}

func (this *RankModule) Init(inner_data *context.GlobalContext) (err error) {
	return

}

func (this *RankModule) abs(a int) int {
	if a < 0 {
		return (0 - a)
	}
	return a
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
	for i := 0; i < len(inner_resp.Ads); i++ {
		inner_resp.Ads[i].MatchAdSlotType = uint32(this.abs(int(inner_data.Req.AdSlot.AdSlotType) - int(inner_resp.Ads[i].AdSlotType)))
		inner_resp.Ads[i].Cpm = rand.Int63()
	}
	sort.Sort(sortad(inner_resp.Ads))

	return
}
