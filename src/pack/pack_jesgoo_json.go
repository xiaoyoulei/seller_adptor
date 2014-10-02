package pack

import (
	"bytes"
	"context"
	"encoding/json"
	"html/template"
	"log"
)

type PackJesgooResponseJsonModule struct {
}

type Ad struct {
	Adslot_id     string
	Material_type int
	Html_snippet  string
}

type SellerReponse struct {
	Success   bool
	Ads       []Ad
	Search_id string
}

func (this PackJesgooResponseJsonModule) Init(inner_data *context.Context) (err error) {
	return
}

func (this PackJesgooResponseJsonModule) Run(inner_data *context.Context) (err error) {
	var temp_resp SellerReponse
	temp_resp.Success = true
	temp_resp.Search_id = inner_data.Searchid
	var temp_ad Ad
	if len(inner_data.Resp.Ads) > 0 {
		temp_ad.Material_type = 1
		var tpl *template.Template
		tpl, err = template.ParseFiles("template/ads.html")
		if err != nil || tpl == nil {
			log.Printf("parse template fail %s", err.Error())
			return
		}
		var temp_html bytes.Buffer
		err = tpl.Execute(&temp_html, inner_data.Resp.Ads[0])
		if err != nil {
			log.Printf("execute template fail %s", err.Error())
			return
		}
		temp_ad.Html_snippet = temp_html.String()
		temp_resp.Ads = append(temp_resp.Ads, temp_ad)
	}
	inner_data.RespBody, err = json.Marshal(temp_resp)
	return
}
