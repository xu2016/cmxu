package xcm

import (
	"encoding/json"
	"log"
	"net/http"
)

//XWErrorJSON 页面函数访问错误返回JSON
type XWErrorJSON struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//XWError Web返回错误
func XWError(w http.ResponseWriter, errstr string, err error) (b bool) {
	rp := &XWErrorJSON{Code: 1, Msg: errstr}
	if err != nil {
		b = true
		log.Println(errstr, err)
		data, _ := json.Marshal(rp)
		w.Write(data)
	}
	return
}

//XError 普通返回错误
func XError(errstr string, err error) (b bool) {
	if err != nil {
		b = true
		log.Println(errstr, err)
	}
	return
}
