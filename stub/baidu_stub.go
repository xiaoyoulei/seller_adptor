package main

import (
	"code.google.com/p/goprotobuf/proto"
	"io/ioutil"
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

	reqbody, err := ioutil.ReadAll(req.Body)
	var temp_req mobads_api.BidRequest
	err = proto.Unmarshal(reqbody, &temp_req)

	var title string
	var desc string
	switch *temp_req.App.Id {
	case "10042c1f":
		title = "I am Banner ad title"
		desc = "I am Banner ad desc"
	case "10044934":
		title = "I am kaiping ad title"
		desc = "I am kaiping ad desc"
	case "10044933":
		title = "I am chaping ad title"
		desc = "I am chaping ad desc"

	}

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
	//	*material.CreativeType = mobads_api.CreativeType_IMAGE
	*material.CreativeType = mobads_api.CreativeType_TEXT
	material.InteractionType = new(mobads_api.InteractionType)
	*material.InteractionType = mobads_api.InteractionType_SURFING
	material.Title = new(string)
	*material.Title = title
	material.Description1 = new(string)
	*material.Description1 = desc
	material.WinNoticeUrl = make([]string, 0)
	material.WinNoticeUrl = append(material.WinNoticeUrl, "http://192.168.1.5:8123/a.gif")
	material.MediaUrl = new(string)
	*material.MediaUrl = "http://192.168.1.5:8123/image/splash.png"
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
