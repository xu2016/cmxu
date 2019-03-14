package ddgl

import (
	"net/http"
)

//Rt 路由
var Rt map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Rt = make(map[string]func(http.ResponseWriter, *http.Request))
	//订单管理
	Rt["/ddcx"] = ddcx //订单查询
	Rt["/ddjk"] = ddjk //订单监控
	Rt["/ddps"] = ddps //订单派送
	Rt["/ddxg"] = ddxg //订单修改
	Rt["/ddxz"] = ddxz //订单下载
	Rt["/htxd"] = htxd //后台下单
}
