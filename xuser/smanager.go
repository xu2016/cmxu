package xuser

import (
	"cmxu/xcm"
	"errors"
	"net/http"
	"sync"
	"time"
)

//GSession 全局Session管理器
var GSession *SManager

//CName cookie名称
var CName string

//SManager Session管理器
type SManager struct {
	cookieName  string              //private cookiename
	lock        sync.Mutex          // protects session
	sid         map[string]ZSession //session id 唯一标示
	maxlifetime int64
}

//NewSessionManager 参加一个Session管理器
func NewSessionManager(cookieName string, maxlifetime int64) (*SManager, error) {
	sid := make(map[string]ZSession, 0)
	return &SManager{cookieName: cookieName, sid: sid, maxlifetime: maxlifetime}, nil
}

//createsid 生成Session唯一ID
func (sm *SManager) createsid(name string) (sid string) {
	sid = xcm.GetMD5(name) + xcm.GetRandomString(8, xcm.NSDSTR)
	return
}

//GC 定时对过期的Session进行删除
func (sm *SManager) GC() {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	for k, v := range sm.sid {
		if (v.timeAccessed.Unix() + sm.maxlifetime) < time.Now().Unix() {
			delete(sm.sid, k)
		}
	}
	time.AfterFunc(time.Duration(sm.maxlifetime*2), func() { sm.GC() })
}

//Set 添加Session
func (sm *SManager) Set(phone string, gid string) (sid string, err error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	sid = sm.createsid(phone)
	if sid == "" {
		err = errors.New("sid create error")
		return
	}
	zs := ZSession{}
	zs.timeAccessed = time.Now()
	zs.gid = gid
	zs.uid = phone
	sm.sid[sid] = zs
	return
}

//Get 获取Session
func (sm *SManager) Get(sid string) (zs ZSession, err error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	zs, ok := sm.sid[sid]
	if !ok {
		err = errors.New("no session")
	}
	return
}

//Del 删除Session
func (sm *SManager) Del(sid string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	delete(sm.sid, sid)
	return
}

//Update 更新Session最后访问时间
func (sm *SManager) Update(sid string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	zs, _ := sm.sid[sid]
	zs.timeAccessed = time.Now()
	sm.sid[sid] = zs
	return
}

//TimeOut 判断Session是否过期
func (sm *SManager) TimeOut(sid string) bool {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	if _, ok := sm.sid[sid]; !ok {
		return true
	}
	if (sm.sid[sid].timeAccessed.Unix() + sm.maxlifetime) < time.Now().Unix() {
		return true
	}
	return false
}

//GetCookie 获取客户端返回的Cookie
func (sm *SManager) GetCookie(r *http.Request, cookieName string) (sid string, err error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return
	}
	sid = cookie.Value
	if sid == "" {
		err = errors.New("sid is empty")
		return
	}
	return
}

//SetCookie 设置客户端返回的Cookie
func (sm *SManager) SetCookie(w http.ResponseWriter, cookieName, sid string) {
	cookie := http.Cookie{Name: cookieName, Value: sid, Path: "/", HttpOnly: true}
	http.SetCookie(w, &cookie)
	return
}

//DeleteCookie 设置客户端返回的Cookie
func (sm *SManager) DeleteCookie(w http.ResponseWriter, cookieName, sid string) {
	cookie := http.Cookie{Name: cookieName, Value: sid, Path: "/", HttpOnly: true, MaxAge: -1}
	http.SetCookie(w, &cookie)
	return
}
