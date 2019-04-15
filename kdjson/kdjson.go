package kdjson

//RJSON 仅仅返回信息
type RJSON struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
