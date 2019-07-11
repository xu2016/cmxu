package xuser

import (
	"cmxu/xcm"
	"errors"
	"strconv"
	"strings"
	"time"
)

//GUKey 全局UniKey管理器
var GUKey *UniKeyManager
var keyStr24 = [24]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}
var keyStr60 = [60]string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09",
	"10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
	"20", "21", "22", "23", "24", "25", "26", "27", "28", "29",
	"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
	"40", "41", "42", "43", "44", "45", "46", "47", "48", "49",
	"50", "51", "52", "53", "54", "55", "56", "57", "58", "59"}

//UniKeyManager 唯一key管理器
type UniKeyManager struct {
	uks map[string]chan UniKey
}

//UniKey 唯一key定义
type UniKey struct {
	RN      int
	USecond int64
}

//NewUniKeyManager 参加一个Session管理器
func NewUniKeyManager(uknames []string) (*UniKeyManager, error) {
	ukm := UniKeyManager{uks: make(map[string]chan UniKey)}
	for _, v := range uknames {
		ukm.uks[v] = make(chan UniKey)
	}
	return &ukm, nil
}

//RunUniKey 新建一个UniKey生成器
func (ukm *UniKeyManager) RunUniKey(ukname string, min, max int) {
	sunix := time.Now().Unix()
	for i := min; ; i++ {
		if sunix != time.Now().Unix() {
			sunix = time.Now().Unix()
			i = min
		}
		if i <= max {
			ukm.uks[ukname] <- UniKey{RN: i, USecond: sunix}
		}
	}
}

/*GetUniKey 获取生成的唯一的编号，编号格式如下：
类型（位数具体确定）+年（4位）月（1位）日（1位）时（1位）分（2位）秒（2位）+随机位（位数具体由生成NewUniKey确定）
key的总长度:len(idtype)+11+len(max)
*/
func (ukm *UniKeyManager) GetUniKey(ukname, idtype string) (id string, err error) {
	rnx, ok := <-ukm.uks[ukname]
	if !ok {
		err = errors.New("get unikey error")
		return
	}
	tm := time.Unix(rnx.USecond, 0)
	yy, mm, dd, hh, ii, ss := xcm.GetTime(tm)
	id = strings.ToUpper(idtype) + strconv.Itoa(yy) + keyStr24[mm] + keyStr24[dd] +
		keyStr24[hh] + keyStr60[ii] + keyStr60[ss] + strconv.Itoa(rnx.RN)
	return
}
