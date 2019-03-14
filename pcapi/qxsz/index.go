package qxsz

import "net/http"

//Rt 路由
var Rt map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Rt = make(map[string]func(http.ResponseWriter, *http.Request))
	//权限管理
	Rt["/cpgl"] = cpgl //产品管理
	Rt["/yhgl"] = yhgl //用户管理
}
