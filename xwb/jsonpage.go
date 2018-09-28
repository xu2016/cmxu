package xwb

import (
	"encoding/json"
	"net/http"
)

//JSONPage 返回Json格式的页面
func JSONPage(cnt interface{}, w http.ResponseWriter) {
	data, _ := json.Marshal(cnt)
	w.Write(data)
	return
}
