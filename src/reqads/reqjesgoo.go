package reqads

import (
	"connpool"
	"context"
	"errors"
	//	"github.com/apache/thrift/lib/go/thrift"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"time"
	"ui2bs"
	"utils"
)

type ReqJesgooModule struct {
	pool    *connpool.ConnPool
	host    string
	port    string
	timeout int
	redis   ReqRedisModule
}

func (this *ReqJesgooModule) Init(global_conf *context.GlobalContext) (err error) {

	err = this.redis.Init(global_conf)
	if err != nil {
		utils.FatalLog.Write("init redis fail, err[%s]", err.Error())
		return
	}

	this.host = global_conf.JesgooBs.Host
	this.port = global_conf.JesgooBs.Port
	this.timeout = global_conf.JesgooBs.Timeout

	this.pool = &connpool.ConnPool{
		Dial: func() (interface{}, error) {
			transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
			protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
			transport, _ := thrift.NewTSocket(net.JoinHostPort(this.host, this.port))
			usetransport := transportFactory.GetTransport(transport)
			client := ui2bs.NewBSServiceClientFactory(usetransport, protocolFactory)
			err := client.Transport.Open()
			if err != nil {
				utils.FatalLog.Write("new jesgoo bs client fail , [%s]", err.Error())
				return nil, err
			}
			return client, err
		},
		Close: func(c interface{}) error {
			err := c.(*ui2bs.BSServiceClient).Transport.Close()
			return err
		},
		Alive: func(c interface{}) bool {
			_, err := c.(*ui2bs.BSServiceClient).Ping()
			if err != nil {
				utils.DebugLog.Write("transport is closed")
				return false
			}
			return true
		},
		MaxIdle: 1024,
	}
	return
}

func (this *ReqJesgooModule) pack_req(inner_data *context.Context, bs_req *ui2bs.BSRequest) (err error) {
	if bs_req == nil {
		utils.FatalLog.Write("bs_req is null")
		err = errors.New("bs_req is null")
		return
	}

	bs_req.Searchid = inner_data.Searchid

	bs_req.Media = new(ui2bs.Media)
	bs_req.Media.Appsid = inner_data.Req.Media.Appsid
	bs_req.Media.ChannelId = inner_data.Req.Media.ChannelId

	bs_req.Adslot = new(ui2bs.AdSlot)
	bs_req.Adslot.Id = inner_data.Req.AdSlot.Slotid
	bs_req.Adslot.Size = new(ui2bs.Size)
	bs_req.Adslot.Size.Width = inner_data.Req.AdSlot.Size.Width
	bs_req.Adslot.Size.Height = inner_data.Req.AdSlot.Size.Height
	//	bs_req.Adslot.TypeA1 = new(ui2bs.Adslot.AdSlotType)
	switch inner_data.Req.AdSlot.AdSlotType {
	case context.AdSlotType_BANNER:
		bs_req.Adslot.TypeA1 = ui2bs.AdSlotType_BANNER
	case context.AdSlotType_OFFERWALL:
		bs_req.Adslot.TypeA1 = ui2bs.AdSlotType_OFFERWALL
	case context.AdSlotType_RECOMMEND:
		bs_req.Adslot.TypeA1 = ui2bs.AdSlotType_RECOMMEND
	case context.AdSlotType_INITIALIZATION:
		bs_req.Adslot.TypeA1 = ui2bs.AdSlotType_INITIALIZATION
	default:
		bs_req.Adslot.TypeA1 = ui2bs.AdSlotType_BANNER
	}

	bs_req.Device = new(ui2bs.Device)
	switch inner_data.Req.Device.OSType {
	case context.OSType_ANDROID:
		bs_req.Device.Os = ui2bs.OSType_ANDROID
	case context.OSType_IOS:
		bs_req.Device.Os = ui2bs.OSType_IOS
	case context.OSType_WP:
		bs_req.Device.Os = ui2bs.OSType_WP
	default:
		bs_req.Device.Os = ui2bs.OSType_UNKNOWN
	}
	//	bs_req.Device.Osv = inner_data.Req.Device.OSVersion
	var temp_device_id *ui2bs.DeviceID
	temp_device_id = new(ui2bs.DeviceID)
	if len(inner_data.Req.Device.DevID) > 0 {
		switch inner_data.Req.Device.DevID[0].DevIDType {
		case context.DeviceIDType_IMEI:
			temp_device_id.TypeA1 = ui2bs.DeviceIDType_IMEI
		case context.DeviceIDType_MAC:
			temp_device_id.TypeA1 = ui2bs.DeviceIDType_MAC
		case context.DeviceIDType_IDFA:
			temp_device_id.TypeA1 = ui2bs.DeviceIDType_IDFA
		default:
			// 临时方案，之后会增加不同类型ID
			temp_device_id.TypeA1 = ui2bs.DeviceIDType_AAID
		}
		temp_device_id.Id = inner_data.Req.Device.DevID[0].ID
	}
	bs_req.Device.DevId = make([]*ui2bs.DeviceID, 0)
	bs_req.Device.DevId = append(bs_req.Device.DevId, temp_device_id)

	bs_req.Network = new(ui2bs.Network)
	bs_req.Location = new(ui2bs.Location)
	bs_req.Location.Province = int32(inner_data.Req.Location.Province)
	bs_req.Location.City = int32(inner_data.Req.Location.City)
	bs_req.Location.Country = int32(inner_data.Req.Location.Country)

	return

}
func (this *ReqJesgooModule) convert_resp_ad(inad *context.AdInfo, bsad *ui2bs.Ad, adtype AdType) {
	inad.AdSrc = context.AdSrc_JESGOO
	switch bsad.Adtype {
	case ui2bs.AdType_TEXT:
		inad.AdType = context.TEXT
	case ui2bs.AdType_IMAGE:
		inad.AdType = context.IMAGE
	case ui2bs.AdType_HTML:
		inad.AdType = context.HTML
	case ui2bs.AdType_VIDEO:
		inad.AdType = context.VIDEO
	case ui2bs.AdType_TEXT_ICON:
		inad.AdType = context.TEXT_ICON
	}
	switch bsad.InteractionType {
	case ui2bs.Interaction_SURFING:
		inad.InteractionType = context.SURFING
	case ui2bs.Interaction_DOWNLOAD:
		inad.InteractionType = context.DOWNLOAD
	case ui2bs.Interaction_DIALING:
		inad.InteractionType = context.DIALING
	case ui2bs.Interaction_MESSAGE:
		inad.InteractionType = context.MESSAGE
	case ui2bs.Interaction_MAIL:
		inad.InteractionType = context.MAIL
	}

	switch adtype {
	case Banner:
		inad.AdSlotType = context.AdSlotType_BANNER
	case Initlization:
		inad.AdSlotType = context.AdSlotType_INITIALIZATION
	case Insert:
		inad.AdSlotType = context.AdSlotType_INSERT
	case OfferWall:
		inad.AdSlotType = context.AdSlotType_OFFERWALL
	case Recommed:
		inad.AdSlotType = context.AdSlotType_RECOMMEND
	default:
		inad.AdSlotType = context.AdSlotType_BANNER
	}

	inad.Adid = bsad.Adid
	inad.Groupid = bsad.Groupid
	inad.Planid = bsad.Groupid
	inad.Userid = bsad.Userid
	inad.Bid = bsad.Bid
	inad.Price = 0
	inad.Ctr = 0
	inad.Cpm = 0

}

func (this *ReqJesgooModule) parse_resp(ret_ads *[]context.AdInfo, bs_resp *ui2bs.BSResponse, adtype AdType) {
	if bs_resp.Ads == nil {
		return
	}
	var ad_num int
	ad_num = len(bs_resp.Ads)
	utils.DebugLog.Write("get jesgoo ad num [%d]", ad_num)
	for i := 0; i < ad_num; i++ {
		var tmpad context.AdInfo
		tmpad.MaterialReady = false
		this.convert_resp_ad(&tmpad, bs_resp.Ads[i], adtype)
		*ret_ads = append(*ret_ads, tmpad)
	}

}

func (this *ReqJesgooModule) ReqBs(inner_data *context.Context, ret_ads *[]context.AdInfo, ch *chan bool, adtype AdType) {
	defer func() {
		*ch <- true
	}()
	var err error
	bs_req := new(ui2bs.BSRequest)

	bs_resp := new(ui2bs.BSResponse)

	err = this.pack_req(inner_data, bs_req)
	if err != nil {
		utils.WarningLog.Write("reqbs pack req fail [%s]", err.Error())
		return
	}
	var client interface{}
	// 用完之后已经要放回连接池
	client, err = this.pool.Get()
	if err != nil {
		utils.FatalLog.Write("get free sock fail ! err[%s]", err.Error())
		return
	}
	utils.DebugLog.Write("get a new client ")
	bs_resp, err = client.(*ui2bs.BSServiceClient).Search(bs_req)
	if err != nil {
		utils.WarningLog.Write("request bs fail [%s]", err.Error())
		return
	}
	this.pool.Put(client)
	this.parse_resp(ret_ads, bs_resp, adtype)
	if len(*ret_ads) > 0 {
		err = this.redis.GetMaterial(ret_ads)
		if err != nil {
			utils.WarningLog.Write("get jesgoo material fail. err[%s]", err.Error())
			return
		}
	} else {
		utils.WarningLog.Write("jesgoo bs return ad num is 0")
	}
	return
}

func (this *ReqJesgooModule) Run(inner_data *context.Context, bschan *chan bool) {
	defer func() {
		if len(inner_data.JesgooAds) == 0 {
			*bschan <- false
		} else {
			*bschan <- true
		}
	}()

	ch := make(chan bool)
	ret_ads := make([]context.AdInfo, 0)
	go this.ReqBs(inner_data, &ret_ads, &ch, Banner)

	inner_data.JesgooAds = make([]context.AdInfo, 0)
	select {
	case <-ch:
		for i := 0; i < len(ret_ads); i++ {
			if ret_ads[i].MaterialReady == true {
				inner_data.JesgooAds = append(inner_data.JesgooAds, ret_ads[i])
			}
		}
	case <-time.After(time.Millisecond * time.Duration(this.timeout)):
		utils.WarningLog.Write("req jesgoo bs timeout")
	}
	return
}
