package xwb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type InfoJSON struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//XPost post方法提交
func XPost(urlstr, qstr, qcstr, qtstr string, rp interface{}) (err error) {
	resp, err := http.PostForm(urlstr, url.Values{"qstr": {qstr}, "qcstr": {qcstr}, "qtstr": {qtstr}})
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal([]byte(body), &rp)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
