package reqads

import (
	"context"
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

	(*bsflag)[0] = true
	//	(*bsflag)[1] = true

	return
}

func (this *ReqBsModule) Run(inner_data *context.Context) (err error) {
	var bsflag []bool
	bsflag = make([]bool, len(this.bsmodule))
	for i := 0; i < len(bsflag); i++ {
		bsflag[i] = false
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
		}
	}
	for i := 0; i < len(bsflag); i++ {
		if bsflag[i] == true {
			utils.DebugLog.Write("waiting index[%d] chan", i)
			<-bschan[i]
		}
	}

	return
}
