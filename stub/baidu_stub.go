package main

import (
	"code.google.com/p/goprotobuf/proto"
	"log"
	"mobads_api"
	"net/http"
)

func main() {

	http.HandleFunc("/api", CallBack)
	err := http.ListenAndServe(":8123", nil)
	if err != nil {
		log.Fatalf("start server fail . err[%s]", err.Error())
	}
	return
}

func CallBack(resp http.ResponseWriter, req *http.Request) {
	var temp_ans mobads_api.BidResponse

	temp_ans.RequestId = new(string)
	*temp_ans.RequestId = "aaaaaa"
	temp_ans.Ads = make([]*mobads_api.Ad, 0)
	var ad *mobads_api.Ad
	ad = new(mobads_api.Ad)
	ad.AdslotId = new(string)
	*ad.AdslotId = "123"
	ad.MaterialMeta = new(mobads_api.Ad_MaterialMeta)
	material := ad.MaterialMeta
	material.CreativeType = new(mobads_api.CreativeType)
	*material.CreativeType = mobads_api.CreativeType_IMAGE
	material.InteractionType = new(mobads_api.InteractionType)
	*material.InteractionType = mobads_api.InteractionType_SURFING
	material.WinNoticeUrl = make([]string, 0)
	material.WinNoticeUrl = append(material.WinNoticeUrl, "http://aaa.com/a.gif")
	material.MediaUrl = new(string)
	*material.MediaUrl = "http://192.168.1.110:8090/splash.png"
	material.ClickUrl = new(string)
	*material.ClickUrl = "http://www.jesgoo.com"
	temp_ans.Ads = append(temp_ans.Ads, ad)

	buf, err := proto.Marshal(&temp_ans)
	if err != nil {
		log.Printf("serialize fail . err[%s]", err.Error())
	}

	resp.Write(buf)

	return
}
