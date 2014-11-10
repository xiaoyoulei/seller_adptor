package main

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"io/ioutil"
	"jesgoo_interface"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var w sync.WaitGroup

func main() {
	arg_num := len(os.Args)
	if arg_num != 2 {
		fmt.Println("get args error . input num will be sent")
		return
	}
	coutn, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("convert args fail . err[%s]\n", err.Error())
	}
	for i := 0; i < coutn; i++ {
		time.Sleep(time.Millisecond * 1)
		w.Add(1)
		go run()
	}
	w.Wait()
}

func run() {

	var req_body jesgoo_interface.SellerRequest
	var resp_body jesgoo_interface.SellerResponse
	var media_tmp jesgoo_interface.Media
	media_tmp.Id = new(string)
	*media_tmp.Id = "123"
	media_tmp.ChannelId = new(string)
	*media_tmp.ChannelId = "123"
	media_tmp.Type = new(jesgoo_interface.MediaType)
	*media_tmp.Type = jesgoo_interface.MediaType_APP
	req_body.Media = new(jesgoo_interface.Media)
	*req_body.Media = media_tmp
	var deviceid_tmp jesgoo_interface.DeviceID
	deviceid_tmp.Type = new(jesgoo_interface.DeviceIDType)
	*deviceid_tmp.Type = jesgoo_interface.DeviceIDType_IMEI
	deviceid_tmp.Id = []byte("351806037915050")
	//	deviceid_tmp.Id = make([]byte, 0)
	//	*deviceid_tmp.Id = []byte("351806037915050")
	//	deviceid_tmp.Id = append(deviceid_tmp.Id, []byte("351806037915050"))
	var device_tmp jesgoo_interface.Device
	device_tmp.Ids = make([]*jesgoo_interface.DeviceID, 0)
	device_tmp.Ids = append(device_tmp.Ids, &deviceid_tmp)
	device_tmp.Type = new(jesgoo_interface.DeviceType)
	*device_tmp.Type = jesgoo_interface.DeviceType_DEV_PHONE
	device_tmp.OsType = new(jesgoo_interface.OSType)
	*device_tmp.OsType = jesgoo_interface.OSType_ANDROID
	device_tmp.OsVersion = new(jesgoo_interface.Version)
	device_tmp.OsVersion.Major = new(uint32)
	*device_tmp.OsVersion.Major = 1
	req_body.Device = new(jesgoo_interface.Device)
	*req_body.Device = device_tmp
	req_body.Network = new(jesgoo_interface.Network)
	req_body.Network.Ip = new(string)
	*req_body.Network.Ip = "61.171.44.61"
	req_body.Adslots = make([]*jesgoo_interface.AdSlot, 0)
	var adslot jesgoo_interface.AdSlot
	adslot.Id = new(string)
	*adslot.Id = "123"
	adslot.Type = new(jesgoo_interface.AdSlotType)
	*adslot.Type = jesgoo_interface.AdSlotType_BANNER
	adslot.Size = new(jesgoo_interface.Size)
	adslot.Size.Width = new(uint32)
	*adslot.Size.Width = 0
	adslot.Size.Height = new(uint32)
	*adslot.Size.Height = 0
	req_body.Adslots = append(req_body.Adslots, &adslot)

	var client_tmp jesgoo_interface.Client
	client_tmp.Type = new(jesgoo_interface.ClientType)
	*client_tmp.Type = jesgoo_interface.ClientType_NATIVESDK
	client_tmp.Version = new(jesgoo_interface.Version)
	client_tmp.Version.Major = new(uint32)
	*client_tmp.Version.Major = 1
	req_body.Client = new(jesgoo_interface.Client)
	*req_body.Client = client_tmp

	fmt.Println(req_body)
	b, err := proto.Marshal(&req_body)
	if err != nil {
		fmt.Println("format err:", err)
	}

	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post("http://localhost:8081/v1/protobuf", "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("response header is: %d\n", res.StatusCode)
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = proto.Unmarshal(result, &resp_body)
	fmt.Println("response is: " + resp_body.String())
	w.Done()
}
