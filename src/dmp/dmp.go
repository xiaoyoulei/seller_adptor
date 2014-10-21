package dmp

import (
	"context"
	"dict"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type DMPModule struct {
	ipdict *dict.IPDictModule
}

func (this *DMPModule) Init(global_conf *context.GlobalContext) (err error) {
	this.ipdict = dict.IPDict
	if this.ipdict == nil {
		err = errors.New("ipdict is nil")
		return
	}
	return
}

func (this *DMPModule) ip2int(ipstr string) (ipint uint32, err error) {
	ips := strings.Split(ipstr, ".")
	if len(ips) != 4 {
		str := fmt.Sprintf("invalid ip [%s]", ipstr)
		err = errors.New(str)
		return
	}
	ipint = 0
	var tempint int
	for i := 0; i < 4; i++ {
		tempint, err = strconv.Atoi(ips[i])
		if err != nil {
			return
		}
		ipint += uint32(tempint) << uint(24-i*8)
	}
	return
}

func (this *DMPModule) Run(inner_data *context.Context) (err error) {
	var ipint uint32
	ipint, err = this.ip2int(inner_data.Req.Network.Ip)
	loc := &inner_data.Req.Location
	if err != nil {
		loc.Country = 0
		loc.Province = 0
		loc.City = 0
	} else {
		loc.Country, loc.Province, loc.City = this.ipdict.Search(ipint)
	}
	utils.DebugLog.Write("find location country[%d] pro[%d] city[%d]", loc.Country, loc.Province, loc.City)
	return
}
