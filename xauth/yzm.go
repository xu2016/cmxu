package xauth

import (
	"cmxu/xcm"
	"sync"
	"time"
)

//GYzm 全局验证码管理器
var GYzm *Cyzm

func init() {
	GYzm = NewCyzm(600)
	go GYzm.GC()
}

//Cyzm 验证码管理器
type Cyzm struct {
	lock        sync.Mutex           // protects session
	yzm         map[string]time.Time //session id 唯一标示
	maxlifetime int64
}

//NewCyzm 参加一个验证码管理器
func NewCyzm(maxlifetime int64) *Cyzm {
	yzm := make(map[string]time.Time, 0)
	return &Cyzm{yzm: yzm, maxlifetime: maxlifetime}
}

//GC 定时对过期的验证码进行删除
func (yzmm *Cyzm) GC() {
	yzmm.lock.Lock()
	defer yzmm.lock.Unlock()
	for k, v := range yzmm.yzm {
		if (v.Unix() + yzmm.maxlifetime) < time.Now().Unix() {
			delete(yzmm.yzm, k)
		}
	}
	time.AfterFunc(time.Duration(yzmm.maxlifetime*2), func() { yzmm.GC() })
}

//SetYzm 设置验证码md5(yzm+uid)
func (yzmm *Cyzm) SetYzm(yzm, uid string) {
	zsstr := xcm.GetMD5(yzm + uid)
	yzmm.lock.Lock()
	defer yzmm.lock.Unlock()
	yzmm.yzm[zsstr] = time.Now()
	return
}

//GetYzm 验证验证码是否正确
func (yzmm *Cyzm) GetYzm(yzm, uid string) bool {
	yzmm.lock.Lock()
	defer yzmm.lock.Unlock()
	yzmid := xcm.GetMD5(yzm + uid)
	if yzmTime, ok := yzmm.yzm[yzmid]; ok {
		if (yzmTime.Unix() + 500) < time.Now().Unix() {
			delete(yzmm.yzm, yzmid)
			return false
		}
		return true
	}
	return false
}
