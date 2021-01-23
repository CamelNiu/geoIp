package ipInfo

import (
	"encoding/json"
	"github.com/oschwald/geoip2-golang"
	"net"
	"geoIp/logFmt"
)

const (
	/**
	ip库
	 */
	GEOIPPATHLOCAL = "/Users/niushaogang/backups/GeoLite2-City-20180703.mmdb"
	GEOIPPATHALI = "/data/www/download/GeoLite2-City-20180703.mmdb"
)

/**
ip详情整理之后的结构体
 */
type ipInfo struct {
	Continent map[string]string `json:"continent"`
	Country map[string]string `json:"country"`
	Subdivisions interface{} `json:"subdivisions"`
	City map[string]string `json:"city"`
	Location interface{} `json:"location"`
}

/**
ip详情待输出数据
 */
type resData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func GetInfo(ip string) string {

	var res string

	resDatas := &resData{
		Code: 0,
		Msg: "ok",
		Data: nil,
	}

	db,err := geoip2.Open(GEOIPPATHLOCAL)

	if err != nil {
		errMsg := "Open ip_db error ["+ip+"]"
		resDatas.Code = -1
		resDatas.Msg = errMsg
		res = formatReturn(resDatas)
		LogFmt.SetErrLog(errMsg)
		return res
	}

	defer db.Close()
	ipFmt := net.ParseIP(ip)

	if ipFmt == nil {
		errMsg := "This Ip Format error ["+ip+"]"
		resDatas.Code = -2
		resDatas.Msg = errMsg
		res = formatReturn(resDatas)
		LogFmt.SetErrLog(errMsg)
		return res
	}

	record,err := db.City(ipFmt)

	if err != nil {
		errMsg := "Read ip_db error ["+ip+"]"
		resDatas.Code = -3
		resDatas.Msg = errMsg
		res = formatReturn(resDatas)
		LogFmt.SetErrLog(errMsg)
		return res
	}

	ipInfo := &ipInfo{
		City:record.City.Names,
		Country: record.Country.Names,
		Continent:record.Continent.Names,
		Location:record.Location,
		Subdivisions:record.Subdivisions,
	}
	resDatas.Data = *ipInfo
	res = formatReturn(resDatas)
	return res
}


func formatReturn( r *resData ) string {
	resDataJson,_ := json.Marshal(r)
	resDataStr := string(resDataJson)
	return resDataStr
}

