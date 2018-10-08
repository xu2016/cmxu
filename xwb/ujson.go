package xwb

import (
	"encoding/json"
	"log"
)

//UJSON 解析JSON
func UJSON(body []byte, rp interface{}) (err error) {
	err = json.Unmarshal(body, &rp)
	if err != nil {
		log.Println("XPost Unmarshal err:", err)
		return
	}
	return
}
