package xuser

import (
	"cmxu/xcm"
	"strconv"
	"sync"
)

//UniKey 唯一key管理器
type UniKey struct {
	lock sync.Mutex // protects session
	cnt  int        //8位计数器，每maxlifetime返回10000001
}

//NewUniKey 新建一个UniKey构建器
func NewUniKey() (*UniKey, error) {
	return &UniKey{cnt: 100000}, nil
}

//Reset 重置计数器
func (uk *UniKey) Reset() {
	uk.lock.Lock()
	defer uk.lock.Unlock()
	uk.cnt = 100000
	return
}

//Get 获取计数器
func (uk *UniKey) Get() (unikey string) {
	uk.lock.Lock()
	defer uk.lock.Unlock()
	yy, mm, dd, hh, mn, ss := xcm.ReturnTime()
	tms, _ := xcm.NumCodeStr([]int{yy % 100, mm, dd, hh, mn, ss})
	uk.cnt++
	unikey = tms + strconv.Itoa(uk.cnt)
	return
}
