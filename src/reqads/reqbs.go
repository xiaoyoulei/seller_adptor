package reqads

import (
	"context"
	//	"math/rand"
	"utils"
)

type AdType int

const (
	Banner       AdType = 0
	Initlization AdType = 1
	Insert       AdType = 2
	OfferWall    AdType = 3
	Recommed     AdType = 4
	MaxAdType    AdType = 5
)

type IBsModule interface {
	Init(golal_conf *context.GlobalContext) (err error)
	Run(inner_data *context.Context, bschan *chan bool)
}
type ReqBsModule struct {
	bsmodule []IBsModule
	timeout  int
}

func (this *ReqBsModule) Init(global_conf *context.GlobalContext) (err error) {
	this.bsmodule = make([]IBsModule, 0)
	this.bsmodule = append(this.bsmodule, &ReqQiushiModule{})
	this.bsmodule = append(this.bsmodule, &ReqJesgooModule{})

	for i := 0; i < len(this.bsmodule); i++ {
		err = this.bsmodule[i].Init(global_conf)
		if err != nil {
			utils.FatalLog.Write("Init reqbs module fail . err[%s]", err.Error())
			return
		}
	}

	this.timeout = global_conf.Dsp.Timeout
	return
}

func (this *ReqBsModule) strategy(inner_data *context.Context, bsflag *[]bool) (err error) {

	switch inner_data.Req.AdSlot.AdSlotType {
	case context.AdSlotType_BANNER:
		(*bsflag)[0] = true
	case context.AdSlotType_INITIALIZATION:
		(*bsflag)[0] = true
	case context.AdSlotType_INSERT:
		(*bsflag)[0] = true
	case context.AdSlotType_RECOMMEND:
		(*bsflag)[1] = true
	case context.AdSlotType_OFFERWALL:
		(*bsflag)[1] = true
	default:
		(*bsflag)[0] = true
	}
	return
}

func (this *ReqBsModule) Run(inner_data *context.Context) (err error) {
	var bsflag []bool
	var bsans []bool
	bsflag = make([]bool, len(this.bsmodule))
	bsans = make([]bool, len(this.bsmodule))
	for i := 0; i < len(bsflag); i++ {
		bsflag[i] = false
		bsans[i] = true
	}
	var bschan []chan bool
	bschan = make([]chan bool, len(this.bsmodule))
	for i := 0; i < len(bschan); i++ {
		bschan[i] = make(chan bool)
	}
	utils.DebugLog.Write("len of bschan is [%d]", len(bschan))

	err = this.strategy(inner_data, &bsflag)
	if err != nil {
		return
	}
	for i := 0; i < len(bsflag); i++ {
		if bsflag[i] == true {
			go this.bsmodule[i].Run(inner_data, &bschan[i])
			utils.DebugLog.Write("req bs[%d]", i)
		}
	}
	for i := 0; i < len(bsflag); i++ {
		if bsflag[i] == true {
			utils.DebugLog.Write("waiting index[%d] chan", i)
			bsans[i] = <-bschan[i]
		}
	}
	// 补余逻辑, 如果有bsans 为false的时候，说明有请求没有返回，则把所有bsflag 取反，用剩余的去补余
	// 前提条件，所有的下游都支持相同的请求
	var need_buyu bool
	need_buyu = false
	for i := 0; i < len(bsans); i++ {
		if bsans[i] == false {
			need_buyu = true
		}
	}
	if need_buyu == true {
		utils.WarningLog.Write("buyu open, appsid[%s], channelid[%s]", inner_data.Req.Media.Appsid, inner_data.Req.Media.ChannelId)
		for i := 0; i < len(bsflag); i++ {
			bsflag[i] = !bsflag[i]
		}
		for i := 0; i < len(bsflag); i++ {
			if bsflag[i] == true {
				go this.bsmodule[i].Run(inner_data, &bschan[i])
				utils.DebugLog.Write("buyu req bs[%d]", i)
			}
		}
		for i := 0; i < len(bsflag); i++ {
			if bsflag[i] == true {
				utils.DebugLog.Write("waiting index[%d] chan", i)
				<-bschan[i]
			}
		}
	}
	return
}
