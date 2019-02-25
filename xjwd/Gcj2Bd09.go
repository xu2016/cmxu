package xjwd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type bd09Json struct {
	Status int                  `json:"status"`
	Result []map[string]float64 `json:"result"`
}

//Gcj2Bd09 把GCJ02坐标系转换成Bd09坐标系
//url=http://api.map.baidu.com/geoconv/v1/?from=1&to=5&coords=114.21892734521,29.575429778924&ak=你的密钥
//from:3、GCJ02,5、BD09ll,to:3、GCJ02,5、BD09ll
func Gcj2Bd09(url string, lng float64, lat float64, key string) (bd09lng, bd09lat float64, err error) {
	var bd bd09Json
	bd.Status = 1
	resp, err := http.Get(url + "&coords=" + strconv.FormatFloat(lng, 'f', -1, 64) + "," + strconv.FormatFloat(lat, 'f', -1, 64) + "&ak=" + key)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(body), &bd)
	if bd.Status == 0 {
		bd09lng = bd.Result[0]["x"]
		bd09lat = bd.Result[0]["y"]
		err = nil
	}
	return
}
