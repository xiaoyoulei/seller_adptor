package reqads

import (
	"context"
	"errors"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
	"ui2bs"
)

type ReqBSModule struct {
}

var transportFactory thrift.TTransportFactory
var protocolFactory *thrift.TBinaryProtocolFactory
var transport *thrift.TSocket
var useTransport thrift.TTransport
var client *ui2bs.BSServiceClient
var err error

func (this *ReqBSModule) InitReqBs(host string, port string) {

	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	transport, err = thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal("create transport fail")
	}

	useTransport = transportFactory.GetTransport(transport)
	client = ui2bs.NewBSServiceClientFactory(useTransport, protocolFactory)
	err := transport.Open()
	if err != nil {
		log.Fatal("open transport fail")
	}
	return

}

func (this *ReqBSModule) pack_req(inner_data *context.Context, bs_req *ui2bs.BSRequest) (err error) {
	if bs_req == nil || bs_req.Media == nil || bs_req.Device == nil || bs_req.Adslot == nil {
		log.Println("bs_req is null")
		err = errors.New("bs_req is null")
		return
	}
	bs_req.Searchid = inner_data.Searchid
	bs_req.Media.Appsid = inner_data.Req.Media.Appsid
	bs_req.Media.ChannelId = inner_data.Req.Media.ChannelId
	bs_req.Adslot.Id = inner_data.Req.AdSlot.Slotid
	bs_req.Adslot.Size = new(ui2bs.Size)
	bs_req.Adslot.Size.Width = inner_data.Req.AdSlot.Size.Width
	bs_req.Adslot.Size.Height = inner_data.Req.AdSlot.Size.Height
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
	log.Println(bs_req)
	return

}
func (this *ReqBSModule) convert_resp_ad(inad *context.AdInfo, bsad *ui2bs.Ad) {
	inad.Adid = bsad.Adid
	inad.Groupid = bsad.Groupid
	inad.Planid = bsad.Groupid
	inad.Userid = bsad.Userid
	inad.Title = bsad.Title
	inad.Bid = bsad.Bid
	inad.Price = 0
	inad.Ctr = 0
	inad.Cpm = 0
	//	inad.Wuliao_type = bsad.WuliaoType
	inad.Description1 = bsad.Desc
	inad.Description2 = ""
	inad.ImageUrl = bsad.ImgUrl
	inad.LogoUrl = bsad.AppLogo
	inad.ClickUrl = bsad.TargetUrl
	inad.AdType = context.TEXT
	//inad.ImpressionUrl = ""
}

func (this *ReqBSModule) parse_resp(inner_data *context.Context, bs_resp *ui2bs.BSResponse) {
	inner_data.JesgooAds = make([]context.AdInfo, 0)
	var ad_num int
	ad_num = len(bs_resp.Ads)
	for i := 0; i < ad_num; i++ {
		var tmpad context.AdInfo
		this.convert_resp_ad(&tmpad, bs_resp.Ads[i])
		inner_data.JesgooAds = append(inner_data.JesgooAds, tmpad)
	}

}

func (this *ReqBSModule) ReqBs(inner_data *context.Context) (err error) {

	bs_req := new(ui2bs.BSRequest)
	bs_req.Media = new(ui2bs.Media)
	bs_req.Device = new(ui2bs.Device)
	bs_req.Adslot = new(ui2bs.AdSlot)

	bs_resp := new(ui2bs.BSResponse)
	this.InitReqBs("218.244.131.175", "8900")
	transport.Open()
	defer transport.Close()
	err = this.pack_req(inner_data, bs_req)
	log.Println(bs_req)
	if err != nil {
		log.Println(err)
		return
	}
	bs_resp, err = client.Search(bs_req)
	if err != nil {
		log.Println("request bs fail")
		log.Println(err)
	}
	this.parse_resp(inner_data, bs_resp)
	log.Println("bs return :")
	log.Println(bs_resp)
	return
}

func (this *ReqBSModule) Run(inner_data *context.Context) (err error) {
	err = this.ReqBs(inner_data)
	return
}
func (this *ReqBSModule) Init(inner_data *context.GlobalContext) (err error) {
	this.InitReqBs("218.244.131.175", "8900")
	return
}
