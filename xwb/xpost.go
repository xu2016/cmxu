package xwb

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	//PostSuc XPost成功
	PostSuc = iota
	//PostERR Post提交失败
	PostERR
	//ReadERR 读取HTML内容失败
	ReadERR
	//UJsonERR 解析JSON失败
	UJsonERR
)

//XPost post方法提交
func XPost(urlstr, qstr, qcstr, qtstr string, rr interface{}) (flag int, err error) {
	flag = PostSuc
	resp, err := http.PostForm(urlstr, url.Values{"qstr": {qstr}, "qcstr": {qcstr}, "qtstr": {qtstr}})
	if err != nil {
		log.Println("XPost post err:", err)
		flag = PostERR
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("XPost ReadAll err:", err)
		flag = ReadERR
		return
	}
	err = UJSON(body, rr)
	if err != nil {
		flag = UJsonERR
	}
	return
}
