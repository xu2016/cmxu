package xtsz

import "net/http"

//Rt 路由
var Rt map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Rt = make(map[string]func(http.ResponseWriter, *http.Request))
	//系统设置
	Rt["/login"] = login   //登陆
	Rt["/logout"] = logout //用户注销
	Rt["/reg"] = reg       //注册
	Rt["/mmxg"] = mmxg     //密码修改
	Rt["/zlxg"] = zlxg     //资料修改
}
