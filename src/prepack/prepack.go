package prepack

import (
	"code.google.com/p/goprotobuf/proto"
	"context"
	"encoding/base64"
	"html/template"
	"jesgoo_protocol"
	"log"
	"math/rand"
	"strings"
	"time"
)

type PrePackModule struct {
	click_head string
	win_head   string
	coder      [64]*base64.Encoding
	basecoder  *base64.Encoding
	texttpl    *template.Template
	imgtpl     *template.Template
}

func (this *PrePackModule) gentexthtml(ad *context.AdInfo, inner_data *context.Context) (err error) {
	err = this.texttpl.Execute(&ad.HtmlSnippet, ad)
	if err != nil {
		return
	}
	return
}

func (this *PrePackModule) genimghtml(ad *context.AdInfo, inner_data *context.Context) (err error) {
	err = this.imgtpl.Execute(&ad.HtmlSnippet, ad)
	if err != nil {
		return
	}
	return
}

func (this *PrePackModule) geniconhtml(ad *context.AdInfo, inner_data *context.Context) (err error) {
	err = this.texttpl.Execute(&ad.HtmlSnippet, ad)
	if err != nil {
		return
	}
	return
}
func (this *PrePackModule) gencurl(ad *context.AdInfo, inner_data *context.Context) (curl string, err error) {
	var event_head jesgoo_protocol.Event_Head
	event_head.CryptoType = new(jesgoo_protocol.Event_Head_CryptoType)
	*event_head.CryptoType = jesgoo_protocol.Event_Head_JESGOO_BASE64
	event_head.CryptoParam = new(uint32)
	*event_head.CryptoParam = uint32(rand.Intn(len(this.coder)))
	coder := this.coder[*event_head.CryptoParam]
	var event_body jesgoo_protocol.Event_Body
	event_body.Type = new(jesgoo_protocol.Event_Body_EventType)
	*event_body.Type = jesgoo_protocol.Event_Body_CLICK
	event_body.SearchId = new(string)
	*event_body.SearchId = inner_data.Searchid
	event_body.SearchTimestamp = new(uint32)
	*event_body.SearchTimestamp = uint32(time.Now().Unix())
	event_body.Media = new(jesgoo_protocol.Event_Body_Media)
	event_body.Media.MediaId = new(string)
	event_body.Media.ChannelId = new(string)
	*event_body.Media.MediaId = inner_data.Req.Media.Appsid
	*event_body.Media.ChannelId = inner_data.Req.Media.ChannelId
	//	event_body.Region = new(jesgoo_protocol.Event_Body_Region)
	event_body.Ad = new(jesgoo_protocol.Event_Body_Ad)
	event_body.Ad.UserId = new(uint32)
	*event_body.Ad.UserId = uint32(ad.Userid)
	event_body.Ad.PlanId = new(uint32)
	*event_body.Ad.PlanId = uint32(ad.Planid)
	event_body.Ad.GroupId = new(uint32)
	*event_body.Ad.GroupId = uint32(ad.Groupid)
	event_body.Ad.UnitId = new(uint32)
	*event_body.Ad.UnitId = uint32(ad.Adid)
	event_body.Charge = new(jesgoo_protocol.Event_Body_Charge)
	event_body.Charge.Type = new(jesgoo_protocol.ChargeType)
	*event_body.Charge.Type = jesgoo_protocol.ChargeType_CPC
	event_body.Charge.Price = new(uint32)
	*event_body.Charge.Price = uint32(ad.Price)
	event_body.Action = new(jesgoo_protocol.Event_Body_Action)
	event_body.Action.TargetUrl = new(string)
	*event_body.Action.TargetUrl = ad.ClickUrl

	var head_enc []byte
	head_enc, err = proto.Marshal(&event_head)
	if err != nil {
		log.Println("marshal click_head fail")
		return
	}
	head_str := this.basecoder.EncodeToString(head_enc)

	var body_enc []byte
	body_enc, err = proto.Marshal(&event_body)
	if err != nil {
		log.Println("marshal click_body fail")
		return
	}
	/////////test code //////////////
	/*	log.Printf("body [%s]", event_body.String())

		test_file := "encode.txt"
		var fout *os.File
		fout, err = os.Create(test_file)
		defer fout.Close()
		if err != nil {
			log.Printf("open file fail\n")
		} else {
			fout.Write(body_enc)
		}
	*/
	/////////////////////////////////
	body_str := coder.EncodeToString(body_enc)
	curl = this.click_head + strings.Replace(head_str, "=", "", -1) + "." + strings.Replace(body_str, "=", "", -1)
	log.Printf("curl is [%s]", curl)
	return
}

func (this *PrePackModule) genwinnotice(ad *context.AdInfo, inner_data *context.Context) (winnotice string, err error) {
	var event_head jesgoo_protocol.Event_Head
	event_head.CryptoType = new(jesgoo_protocol.Event_Head_CryptoType)
	*event_head.CryptoType = jesgoo_protocol.Event_Head_JESGOO_BASE64
	event_head.CryptoParam = new(uint32)
	*event_head.CryptoParam = uint32(rand.Intn(len(this.coder)))
	coder := this.coder[*event_head.CryptoParam]
	var event_body jesgoo_protocol.Event_Body
	event_body.Type = new(jesgoo_protocol.Event_Body_EventType)
	*event_body.Type = jesgoo_protocol.Event_Body_IMPRESSION
	event_body.SearchId = new(string)
	*event_body.SearchId = inner_data.Searchid
	event_body.SearchTimestamp = new(uint32)
	*event_body.SearchTimestamp = uint32(time.Now().Unix())
	event_body.Media = new(jesgoo_protocol.Event_Body_Media)
	event_body.Media.MediaId = new(string)
	event_body.Media.ChannelId = new(string)
	*event_body.Media.MediaId = inner_data.Req.Media.Appsid
	*event_body.Media.ChannelId = inner_data.Req.Media.ChannelId
	//	event_body.Region = new(jesgoo_protocol.Event_Body_Region)
	event_body.Ad = new(jesgoo_protocol.Event_Body_Ad)
	event_body.Ad.UserId = new(uint32)
	*event_body.Ad.UserId = uint32(ad.Userid)
	event_body.Ad.PlanId = new(uint32)
	*event_body.Ad.PlanId = uint32(ad.Planid)
	event_body.Ad.GroupId = new(uint32)
	*event_body.Ad.GroupId = uint32(ad.Groupid)
	event_body.Ad.UnitId = new(uint32)
	*event_body.Ad.UnitId = uint32(ad.Adid)
	event_body.Charge = new(jesgoo_protocol.Event_Body_Charge)
	event_body.Charge.Type = new(jesgoo_protocol.ChargeType)
	*event_body.Charge.Type = jesgoo_protocol.ChargeType_CPC
	event_body.Charge.Price = new(uint32)
	*event_body.Charge.Price = uint32(ad.Price)
	//	event_body.Action = new(jesgoo_protocol.Event_Body_Action)
	//	event_body.Action.TargetUrl = new(string)
	//	*event_body.Action.TargetUrl = ad.ClickUrl

	var head_enc []byte
	head_enc, err = proto.Marshal(&event_head)
	if err != nil {
		log.Println("marshal click_head fail")
		return
	}
	head_str := this.basecoder.EncodeToString(head_enc)

	var body_enc []byte
	body_enc, err = proto.Marshal(&event_body)
	if err != nil {
		log.Println("marshal click_body fail")
		return
	}
	body_str := coder.EncodeToString(body_enc)
	winnotice = this.win_head + strings.Replace(head_str, "=", "", -1) + "." + strings.Replace(body_str, "=", "", -1)
	return
}

func (this *PrePackModule) Run(inner_data *context.Context) (err error) {
	for i := 0; i < len(inner_data.Resp.Ads); i++ {
		var temp_click string
		var temp_winnotice string
		temp_click, err = this.gencurl(&inner_data.Resp.Ads[i], inner_data)
		if err != nil {
			log.Printf("make curl fail [%s]\n", err.Error())
			return
		}
		temp_winnotice, err = this.genwinnotice(&inner_data.Resp.Ads[i], inner_data)
		if err != nil {
			log.Printf("make winnotice fail [%s]\n", err.Error())
			return
		}
		inner_data.Resp.Ads[i].ClickUrl = temp_click
		inner_data.Resp.Ads[i].ImpressionUrl = append(inner_data.Resp.Ads[i].ImpressionUrl, temp_winnotice)
		switch inner_data.Resp.Ads[i].AdType {
		case context.TEXT:
			err = this.gentexthtml(&inner_data.Resp.Ads[i], inner_data)
		case context.IMAGE:
			err = this.genimghtml(&inner_data.Resp.Ads[i], inner_data)
		default:
			err = this.gentexthtml(&inner_data.Resp.Ads[i], inner_data)
		}
		if err != nil {
			log.Printf("make html fail [%s]\n", err.Error())
			return
		}
	}
	return
}

func (this *PrePackModule) Init(inner_data *context.GlobalContext) (err error) {
	this.click_head = "http://rcv.jesgoo.com/click?"
	this.win_head = "http://rcv.jesgoo.com/show?"
	base64_code := [64]string{"0Ge5Q6wnuBML3Wg-8s7kAjOr2xohRHif1CczEXDKq_VaUPJdYpy9TSNvIFmbZ4tl", "1RPspa70TiZum_6tKNIFOVJ4o3gBGkdLlvhCnQHWDxEA5cq8fYzj2XU9MwrS-ybe", "2txfHr4BFaqSOTVn_CuYozeygvbcL1RNKs8MkWUPG0JpmXd5l6hDIAjw7i9E-Q3Z", "3c5gtSlVFoEXsehLaMJwI0Rv6U7xmWuzCYABP_TDp-QfdKniNrq2Ob419jHkGZy8", "456tTGIcm3EK7_RQpBA2uezWXjqnHPsi9bV-Zx8vSgkwaMlCoJ1DyfLFr0UYNhdO", "5zDkOUe_q7a2TN-Y3fvrnsCHZxhBIdcESPj8GuWXLgib4QwJo6yVtMF0Kp1mA9Rl", "6bgkOJjx8mPI9Y_hWANypSoFtnDadivQ31crfR0uC-K2qzE4UslMXGBTL5Vw7HeZ", "7zXOg9UZp0sa2vWGJP3TCBYR-6eIfHVnmdKwMrbuSlo8LFNthAkqcDy4QxE5ji1_", "8iMsWOQuApUEhoITJfjbYdLqrcF3m17XwNa_B4lVSC-Deny0KPtG6z2k5gx9vZHR", "9gdocC-fpejPiYnHF7l0XrWUAaS3OI_Nm2uDkyZbVGRJq5vt1QzhM6xTLw4BsK8E", "aJEr4IWdD3Zxw8ClmNzGPHFL9usbKTvXU-2eAkiV16ScBOy7pqgQnR0_j5MYftho", "bZqml65BeDJETLrNpMAC-tIOhVXsnf1w9_Szv7gP2xaWKR0cQ4F3Uuo8YHGdjiky", "cIJjl64KtC_doMuUeEq0fv831F2kDnNshAbzVaZTprYHB-mWXwygQ7RGxi9PL5SO", "dZgpqwv-321tRL5ITib0X8Uy64kKjoQuOsrGENxMBfY9WDea7HcmJnAC_PzSlFVh", "eUOghPNMEVoxkiG7YQL0Zmjc_IWsdl2FtRufySH4DvATBwbXCK9z5-par8J61nq3", "fP1pUQ3_moNy-7FXHAGT8uKL4wszIVMjaRdhi25E9O0kWntrBvScl6DZgqbxYJCe", "g43jISvAuK6UND0qYma2QXxRPyBib1l_zCtke5VpsEJOMf7LdZHGcFnw-hW9o8Tr", "humQGil2xCkHY391US7DvZVEbe_PR4NXWqKgrFfJApd60Bt8yno5IMawsj-OLzcT", "i2MxZ56X0LWSQGzpONHwPje_lA1nftYIvRTr-9UJs4DkydcgoVqBCu3bE7Kh8mFa", "jk9A_YBn6aqWt1yfm38SLJ7z-c42erUEHxhb5dROGgNoIusvKCDXQMTp0lZiPwVF", "kCBHjyTZmhSEwfY8JPdX9MpoKWuL61FQDtOnzab-s3INrc4qi0e7lA_gxUV5R2Gv", "lToWfMpEHAcir_JLRdBDSkaU5hPnmOCqVZ2KXNzI4uwGx01Ys379egtQ8yF-vbj6", "m36NTQ82RhD4E-tMfxjs0CPoLbirI_WBZHgnFcypezvkadXSGwO95lAKYJ7Uu1qV", "niOyHQ6oFbWzeM1qmu4tV9grTkldhGj20NELD3pRU8ZSAa7-wxKJsYcPBI5_vCXf", "odnysI_h-FGKmJQ2NvUwgVLcaOHCbSDAqlxp85Zurjzi4Mft793RYPe6TE01XkWB", "p0tGYHmhNTfdFkMQbUc1axCq4y5viXD89suS3BAIPL67VKrWolOgzwRZ_eE2jJ-n", "qrw0jSf_3pxth-LZnPoBITkGAvl7gQKdF69CVyWaD5izE4mN8ecMRHYOu2UJs1Xb", "rjUafJvNocE3D8TiQYVAHLp7WbSnzemqugBy5lIFk10_2CdM-Gh69Ksx4XRZtOwP", "sPWiyNlL3rGBI8gtfxpDbAoUe_5uhzMjZd09cTCamEwYJF6k14QvnVK7Oq-R2HXS", "tHir6_aKGnqf1RQABh3IzSwYysp4kWdUPCce2l7ZV8XJO9MmLbg50NT-vFuEoxjD", "uvfDMaj4s7hUr0GBSxWqyV6z-gnLH9QCFoJZwPK8_1lXYON2RpATId5bimEtce3k", "vy01f3dLjbn4RQuSpiEhKV_XaMB6l2cgxFI-8oUWPAY95wtmZkJNzHOCqrTe7sGD", "wDmOP9vQr_tyguqxobLHzjMC50-IAJiB4EZf1nh36sl8VWTXYGF2d7UaeSkcKpRN", "xyqhnozlE_1cV8mukNLTtHfW6r5gw2Se09A4BC3OvJdIXFpsRDiM7YbPGUajKQZ-", "y6D3Z4lYsGaxOd0nE9irmRthB7ILkPW-NbCjQpSUe12gcFK8uvofzXq5_TAHMJVw", "zH-ExpCWvMa4iPKkDdeTwq3_moAIQFXSUYGgBbL07NftO2r8nJ6lsRyu5cZhV91j", "ApKPs5S8LviIeqzyGMEmkrUhJdRgx-Q12VZFlTt9b7Y3NoBc_uOaD0CfWXHn4jw6", "BTUcJEezrnZg1YKoD3Iy_aGNOPCQtiljm90LvVWXkx48p5AHhf-dqSu67MsFwb2R", "Cfxmu205MWqpTElyPLQhODzjG36JwdsYVINZvSr471tAk_nea-R89oHbFKBicgXU", "De9UpodQBKy12HC_f6xRkYOgXv-7PNEWbuLA0qaGTScs54inrwJmZzjMF3IlVh8t", "E1RkI23NpneclxgmPyZ685Twv4aGrfCQsib-0Md9JhDO7F_VUSHqKYouAWjtBXLz", "Fu1SjpgwfAm5hU3DGdYnEsIX-kl_oaN62KWctPb098CezHZqOR4vVB7yQLixTrJM", "Gv7QT10t6edVH59gcLWOEAblzRkPyFCYNXf4I82iMmwUB_3Z-axKSspnojuJhqrD", "H-lduv6Xih9fQmF3O4JxMY7pcIjTnoatSs0CZBe15U_RyrLgPN8EwkAKGbDzW2Vq", "I9h-UfpFRlVGb_O10Zai3sq7dALjyDte2SzvMCBcYuxK5H8JwNXTPkEWr6gQno4m", "J-fYWMd6jFomtR5SP9lcqEn4aHAVx0ONILrCsyQi_pwTUhg73BvKDz2u1ZkXe8bG", "KboJvRSdBLe_r9EDz8HO-Z5GI0PW7cQuNns63gAhtYilTkqwyCF1xUVaM2m4Xfpj", "LmPDzctkaYWfSZn4yRCIBl253dw8sb9VxJEAXU6erN_Ou0GvMQg-Ko1iT7FhqpjH", "M1DFQXKTH-fJtUwBpngi04odZV3eYO9GuW86NrLh7xEslRP2z_aCkIjAcym5qSbv", "NuzOPiFag3dWUewMmyTEB5kYhcxDL2fbr_V8nRC9HGI0ZQ67KXo-lAqvJSs41jtp", "OXE5d2WJIKelZS0wj3tN6xvhP1QMrF8-nDka7HTyLi_YgVCGAzubomUf49BRcqps", "PIoQ0Vd-DgF9aYLHyGtvBcX61x8mJ_AKMU5bR3sruwTlfh42knSejNiWEqp7OZCz", "QrfzshiqxbpXZGuPV4_HL62on8YwSWITAc-eyO3djvKgNa5mE7tD10BFCMkJUl9R", "R3kYn_jgNAOJCdSwHol84aQbD7U-IhBciV20EuKM9tfrFxZ6XqzvmpLW1GPT5sey", "SLfn20Wt7oU1zMqrYkJcFZ9BmaOueV3sPTKQ5N-xbiI8DC6yp4gAvXGHlwhdRE_j", "T3JinH-K2NPsx4MuacYrXvdmWZjeOAQfp9C8tDz075qE1yVUBRb_L6whSGFoklIg", "UCvz-5K34p9OBSL8tkImg21wGYFfl6NaciZnDx7sRW0b_eyoAEQPrXqMjhHTduJV", "V4miLp8wCY0QGx7oB6_FvjWJeRs2dZK5XDNOAkhUbuMcHanT3zEIlgyr9tf-1SPq", "Wxd0EnjzS6DXcIR1qUoJMmF_fGa8uHs5h9A3VkbTiCZvNplLryKwQY4BOPg7et-2", "XYc4wekTRE8ZoxQ9dtiyO-3aAfuIjVsCFh_7Jp20rqm5SHWL1lBbnUNzDKvGPMg6", "YHS9KRLhdqbMtaeX6wG4QNp21sBFrIV_y7DlCzOx80UiumWAnT5-PkgocJEf3jZv", "Z_FmazrYTKXHW-4sSbk86D5VPyfN13gBu9deJE7q2ontjQOGhxwiCclUAMvLp0IR", "_rJb0Up6R2hNDWFAH3LswnYd7zPQo1M8KxyIjBugESqTGZtmleifcVkCX9v5-aO4", "-6mTYvy7KRxCr2elEiWbL5qMoc9G1_3untsdpjwDAzFVhZSgBN0aIOU4QkH8fJXP"}
	for i := 0; i < len(base64_code); i++ {
		this.coder[i] = base64.NewEncoding(base64_code[i])
	}
	code_char := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	this.basecoder = base64.NewEncoding(code_char)
	this.texttpl, err = template.ParseFiles("template/ads_text.html")
	if err != nil || this.texttpl == nil {
		return
	}
	this.imgtpl, err = template.ParseFiles("template/ads_img.html")
	if err != nil || this.imgtpl == nil {
		return
	}
	return
}
