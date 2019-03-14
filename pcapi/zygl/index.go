package zygl

import "net/http"

//Rt 路由
var Rt map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Rt = make(map[string]func(http.ResponseWriter, *http.Request))
	//资源管理
	Rt["/zycx"] = zycx //资源查询
	Rt["/zyxg"] = zyxg //资源修改
	Rt["/zyxz"] = zyxz //资源下载
}
