package pack

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"context"
	"html/template"
	"jesgoo_protocol"
	"log"
)

type PackJesgooResponseModule struct {
}

func (this PackJesgooResponseModule) Run(inner_data *context.Context) (err error) {

	var temp_response jesgoo_protocol.SellerResponse
	temp_response.Success = new(bool)
	*temp_response.Success = true
	temp_response.SearchId = new(string)
	*temp_response.SearchId = inner_data.Searchid
	var temp_ad jesgoo_protocol.SellerResponse_Ad
	temp_ad.AdslotId = new(string)
	*temp_ad.AdslotId = "123"
	temp_ad.MaterialType = new(jesgoo_protocol.SellerResponse_Ad_MaterialType)
	*temp_ad.MaterialType = jesgoo_protocol.SellerResponse_Ad_DYNAMIC
	var tpl *template.Template
	tpl, err = template.ParseFiles("template/ads.html")
	if err != nil || tpl == nil {
		log.Printf("parse template fail %s", err.Error())
		return
	}
	var temp_html bytes.Buffer
	if len(inner_data.Resp.Ads) > 0 {
		err = tpl.Execute(&temp_html, inner_data.Resp.Ads[0])
		if err != nil {
			log.Printf("execute template fail %s", err.Error())
			return
		}
		temp_ad.HtmlSnippet = make([]byte, 0)
		temp_ad.HtmlSnippet = temp_html.Bytes()
	}
	temp_response.Ads = make([]*jesgoo_protocol.SellerResponse_Ad, 0)
	temp_response.Ads = append(temp_response.Ads, &temp_ad)
	inner_data.RespBody, err = proto.Marshal(&temp_response)
	return
}

func (this PackJesgooResponseModule) Init(inner_data *context.GlobalContext) (err error) {
	return
}
