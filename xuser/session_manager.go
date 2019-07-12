package xuser

import (
	"cmxu/xcm"
	"errors"
	"net/http"
	"sync"
	"time"
)

//Gsm 全局Session管理器
var Gsm *SManager

//SManager Session管理器
type SManager struct {
	cookieName  string             //private cookiename
	lock        sync.RWMutex       // protects session
	sid         map[string]session //session id 唯一标示
	maxlifetime int64
}

//session session存储结构
type session struct {
	uid          string    //用户账号
	timeAccessed time.Time //最后访问时间
	rid          []string  //角色组
	phone        string    //用户手机号
	city         string    //地市
}

//NewSessionManager 参加一个Session管理器
func NewSessionManager(cookieName string, maxlifetime int64) (*SManager, error) {
	sid := make(map[string]session, 0)
	return &SManager{cookieName: cookieName, sid: sid, maxlifetime: maxlifetime}, nil
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

//UserIsLogin 判断用户是否登陆
func (sm *SManager) UserIsLogin(r *http.Request) (sid string, err error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	cookie, err := r.Cookie(sm.cookieName)
	if err != nil {
		return
	}
	sid = cookie.Value
	if sid == "" {
		err = errors.New("No session")
		return
	}
	//判断Session是否存在
	zs, ok := sm.sid[sid]
	if !ok {
		err = errors.New("No session")
		return
	}
	//判断Session是否过期
	if (zs.timeAccessed.Unix() + sm.maxlifetime) < time.Now().Unix() {
		err = errors.New("Session time is out")
		return
	}
	//更新Session最近访问时间
	zs.timeAccessed = time.Now()
	sm.sid[sid] = zs
	return
}

//AddUserLogin 用户登陆添加Session
func (sm *SManager) AddUserLogin(w http.ResponseWriter, userid, phone, city string, rid []string) (err error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	sid := xcm.GetMD5(phone + userid + city + xcm.GetRandomString(8, xcm.NSDSTR))
	if sid == "" {
		err = errors.New("Add  session error")
		return
	}
	zs := session{uid: userid, phone: phone, rid: rid, city: city, timeAccessed: time.Now()}
	sm.sid[sid] = zs
	cookie := http.Cookie{Name: sm.cookieName, Value: sid, Path: "/", HttpOnly: true}
	http.SetCookie(w, &cookie)
	return
}

//DelUserLogin 用户注销删除Session
func (sm *SManager) DelUserLogin(w http.ResponseWriter, r *http.Request) (err error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	cookie, err := r.Cookie(sm.cookieName)
	if err != nil {
		return
	}
	sid := cookie.Value
	delete(sm.sid, sid)
	cookie = &http.Cookie{Name: sm.cookieName, Value: sid, Path: "/", HttpOnly: true, MaxAge: -1}
	http.SetCookie(w, cookie)
	return
}

//GetUserID 获取用户ID
func (sm *SManager) GetUserID(sid string) (uid string, err error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	zs, ok := sm.sid[sid]
	if !ok {
		err = errors.New("No session")
		return
	}
	uid = zs.uid
	return
}

//GetUserPhone 获取用户phone
func (sm *SManager) GetUserPhone(sid string) (phone string, err error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	zs, ok := sm.sid[sid]
	if !ok {
		err = errors.New("No session")
		return
	}
	phone = zs.phone
	return
}

//GetUserCity 获取用户所属地市
func (sm *SManager) GetUserCity(sid string) (city string) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	zs, ok := sm.sid[sid]
	if !ok {
		return
	}
	city = zs.city
	return
}

//GetUserRoles 获取用户角色组
func (sm *SManager) GetUserRoles(sid string) (rid []string) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	zs, ok := sm.sid[sid]
	if !ok {
		return
	}
	rid = zs.rid
	return
}

//GetUserAuth 判断用户是否有访问权限
func (sm *SManager) GetUserAuth(sid string) (b bool) {

	return
}
