package dict

import (
	"code.google.com/p/goprotobuf/proto"
	"context"
	"errors"
	"io/ioutil"
	"jesgoo_protocol"
	"os"
	"utils"
)

type IPDictModule struct {
	IPDict jesgoo_protocol.IPTable
}

func (this *IPDictModule) Init(conf *context.GlobalContext) (err error) {
	ipdict_path := conf.Dict.IPDictPath
	utils.DebugLog.Write("get ipdict_path [%s]", ipdict_path)
	var file *os.File
	file, err = os.OpenFile(ipdict_path, os.O_RDONLY, 0777)
	if err != nil {
		utils.FatalLog.Write("open ipdict[%s] fail. err[%s]", ipdict_path, err.Error())
		return
	}
	defer file.Close()
	data := make([]byte, 0)
	var length int
	data, err = ioutil.ReadAll(file)
	if err != nil {
		utils.FatalLog.Write("Read ipdict fail ! err[%s]", err.Error())
		return
	}
	utils.DebugLog.Write("len of ipdict data is [%d]", length)
	err = proto.Unmarshal(data, &this.IPDict)
	if err != nil {
		utils.FatalLog.Write("parse ipdict fail. err[%s]", err.Error())
		return
	}
	//	utils.DebugLog.Write("ipdict [%s]", this.IPDict.String())
	return

}

func (this *IPDictModule) ipbsearch(data []*jesgoo_protocol.IPTable_IPSection, ipint uint32) (sec jesgoo_protocol.IPTable_IPSection, err error) {
	min := 0
	max := len(data)
	var mid int
	for min <= max {
		mid = (min + max) / 2
		ipsec := data[mid]
		if *ipsec.Start > ipint {
			max = mid - 1
			continue
		}
		if *ipsec.End < ipint {
			min = mid + 1
			continue
		}
		break
	}
	if min > max {
		err = errors.New("not find")
		return
	}
	ipsec := data[mid]
	if *ipsec.Start <= ipint && *ipsec.End >= ipint {
		sec = *data[mid]
		return
	} else {
		err = errors.New("not find")
	}
	return
}

func (this *IPDictModule) Search(ip uint32) (country uint32, province uint32, city uint32) {
	head := ip >> 24
	country = 0
	province = 0
	city = 0
	if len(this.IPDict.Classa) < int(head) {
		utils.FatalLog.Write("IPDict.Classa len is less than [%d] len of classa[%d]", head, len(this.IPDict.Classa))
		return
	}
	iptable := this.IPDict.Classa[head]
	sec, err := this.ipbsearch(iptable.IpSections, ip)
	if err == nil {
		country = *sec.Country
		province = *sec.Province
		city = *sec.City
		utils.WarningLog.Write("ip not find in dict .")
		return
	}
	return
}

var IPDict *IPDictModule
