package xuser

import (
	"cmxu/xcm"
	"errors"
	"sync"
	"time"
)

//Gyzm 全局验证码管理器
var Gyzm *YzmManager

//YzmManager 验证码管理器
type YzmManager struct {
	lock        sync.Mutex      // protects session
	yzms        map[string]CYzm //号码验证码
	maxlifetime int64
}

//CYzm 验证码类
type CYzm struct {
	yzmid        string    //验证码
	timeAccessed time.Time //最后访问时间
}

//NewYzmManager 参加一个Session管理器
func NewYzmManager(maxlifetime int64) (*YzmManager, error) {
	return &YzmManager{yzms: make(map[string]CYzm, 0), maxlifetime: maxlifetime}, nil
}

//GC 定时对过期的验证码进行删除
func (yzm *YzmManager) GC() {
	yzm.lock.Lock()
	defer yzm.lock.Unlock()
	for k, v := range yzm.yzms {
		if (v.timeAccessed.Unix() + yzm.maxlifetime) < time.Now().Unix() {
			delete(yzm.yzms, k)
		}
	}
	time.AfterFunc(time.Duration(yzm.maxlifetime*2), func() { yzm.GC() })
}

//Set 添加验证码
func (yzm *YzmManager) Set(phone string) (yzmid string) {
	yzm.lock.Lock()
	defer yzm.lock.Unlock()
	xyzm := CYzm{yzmid: xcm.GetRandomString(6, xcm.NSTR), timeAccessed: time.Now()}
	yzm.yzms[phone] = xyzm
	yzmid = xyzm.yzmid
	return
}

//Get 获取验证码
func (yzm *YzmManager) Get(phone string) (yzmid string, err error) {
	yzm.lock.Lock()
	defer yzm.lock.Unlock()
	yzms, ok := yzm.yzms[phone]
	if !ok {
		err = errors.New("no yzm")
		return
	}
	yzmid = yzms.yzmid
	return
}
