package sjfx

import "net/http"

//Rt 路由
var Rt map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Rt = make(map[string]func(http.ResponseWriter, *http.Request))
	//数据分析

}
