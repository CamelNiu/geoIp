package httpServer

import (
	"geoIp/ipInfo"
	"net/http"
	"geoIp/logFmt"
)

func Run()  {
	http.HandleFunc("/",handles)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		LogFmt.SetVisitLog("Listen error")
	}
}

/**
连接成功,执行
 */
func handles (w http.ResponseWriter, r *http.Request) {
	ipAdr := GetUrlArg(r,"ip")
	LogFmt.SetVisitLog("visit. ip is ["+ipAdr+"] ")
	ipInfo := ipInfo.GetInfo(ipAdr)
	w.Write([]byte(ipInfo))
}

/**
获取get参数
 */
func GetUrlArg(r *http.Request,name string)string{
	var arg string
	values := r.URL.Query()
	arg=values.Get(name)
	return arg
}