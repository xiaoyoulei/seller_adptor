package reqads

import (
	"code.google.com/p/goprotobuf/proto"
	"context"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"jesgoo_protocol"
	"time"
	"utils"
)

type ReqRedisModule struct {
	pool     *redis.Pool
	location string
	timeout  int
}

func (this *ReqRedisModule) Init(golbal_conf *context.GlobalContext) (err error) {
	this.location = golbal_conf.RedisMaterial.Location
	this.timeout = golbal_conf.RedisMaterial.Timeout
	this.pool = &redis.Pool{
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.Dial("tcp", this.location)
			if err != nil {
				return
			}
			return
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     1024,
		MaxActive:   1024,
		IdleTimeout: time.Duration(this.timeout) * time.Millisecond,
	}
	return
}

func (this *ReqRedisModule) query(v []int64) (ans []interface{}, err error) {
	var c redis.Conn
	c = this.pool.Get()
	var tv []interface{}
	for i := 0; i < len(v); i++ {
		tv = append(tv, interface{}(v[i]))
	}
	var tp interface{}
	tp, err = c.Do("mget", tv...)
	if err != nil {
		utils.FatalLog.Write("request redis fail . err[%s]", err.Error())
		return
	}
	ans = tp.([]interface{})
	return
}

func (this *ReqRedisModule) fill_material(ad *context.AdInfo, material *jesgoo_protocol.Material) (err error) {

	if ad.Adid != int64(*material.Id) {
		err = errors.New("id is not equal")
		return
	}

	switch *material.Type {
	case jesgoo_protocol.Material_TEXT:
		if ad.AdType != context.TEXT {
			str := fmt.Sprintf("adid[%lld] is not TEXT in bs", ad.Adid)
			err = errors.New(str)
			return
		}
	case jesgoo_protocol.Material_IMAGE:
		if ad.AdType != context.IMAGE {
			str := fmt.Sprintf("adid[%lld] is not IMAGE in bs", ad.Adid)
			err = errors.New(str)
			return
		}
	case jesgoo_protocol.Material_ICON_TEXT:
		if ad.AdType != context.TEXT_ICON {
			str := fmt.Sprintf("adid[%lld] is not ICON_TEXT in bs", ad.Adid)
			err = errors.New(str)
			return
		}
	default:
		str := fmt.Sprintf("adid[%lld] is UNKNOWN is bs", ad.Adid)
		err = errors.New(str)
		return
	}
	if material.Title != nil {
		ad.Title = *material.Title
	}
	if material.Description != nil {
		ad.Description1 = *material.Description
	}
	if material.Image != nil {
		ad.ImageUrl = *material.Image.Url
	}
	if material.TargetUrl != nil {
		ad.ClickUrl = *material.TargetUrl
	}
	if material.App != nil {
		ad.Appname = *material.App.Name
		ad.Package = *material.App.PackageName
		ad.LogoUrl = *material.App.Logo.Url
	}
	if material.LongDescription != nil {
		ad.Description2 = *material.LongDescription
	}
	ad.MaterialReady = true
	return
}

func (this *ReqRedisModule) GetMaterial(ads *[]context.AdInfo) (err error) {

	utils.DebugLog.Write("start to get material, ads len[%d]", len(*ads))
	var ids []int64
	for i := 0; i < len(*ads); i++ {
		ids = append(ids, (*ads)[i].Adid)
	}
	var ans []interface{}
	utils.DebugLog.Write("request is [%s]", ids)
	ans, err = this.query(ids)
	if err != nil {
		utils.WarningLog.Write("request redis fail . err[%s]", err.Error())
		return
	}
	var material_tmp jesgoo_protocol.Material
	for i := 0; i < len(ans); i++ {
		tp, err := redis.Bytes(ans[i], nil)
		if err != nil {
			utils.WarningLog.Write("read redis return fail err[%s]", err.Error())
			continue
		}
		err = proto.Unmarshal(tp, &material_tmp)
		if err != nil {
			utils.FatalLog.Write("parse material fail . adid[%d]", (*ads)[i].Adid)
			continue
		}
		err = this.fill_material(&(*ads)[i], &material_tmp)
		if err != nil {
			utils.FatalLog.Write("fill material fail . adid[%lld] err[%s]", (*ads)[i].Adid, err.Error())
		}
		//		utils.DebugLog.Write("get material . [%s]", tp)
	}
	err = nil
	return
}
