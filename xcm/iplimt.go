package xcm

import (
	"log"
	"net/http"
	"strings"
)

//IPLimt 判断访问的客户端IP是不是指定的IP地址
func IPLimt(IPARR []string, r *http.Request) (b bool) {
	fwip := strings.Split(r.RemoteAddr, ":")[0]
	if IPARR == nil {
		log.Println("xcm.IPLimt err:IPARR is nil")
		return
	}
	for _, v := range IPARR {
		if v == fwip {
			b = true
			return
		}
	}
	return
}
