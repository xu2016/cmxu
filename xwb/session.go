package xwb

import (
	"cmxu/xcm"
	"errors"
	"net/http"
	"sync"
	"time"
)

//GSession 全局Session管理器
var GSession *SessionManager

//CookieName ...
var CookieName = "sidx08063600"

//然后在 init 函数中初始化
func init() {
	GSession, _ = NewSessionManager(CookieName, 3600)
	go GSession.GC()
}

//Session session存储结构
type Session struct {
	TimeAccessed time.Time //最后访问时间
	GroupID      int       //群组ID
}

//SessionManager Session管理器
type SessionManager struct {
	cookieName  string             //private cookiename
	lock        sync.Mutex         // protects session
	sid         map[string]Session //session id 唯一标示
	maxlifetime int64
}

//NewSessionManager 参加一个Session管理器
func NewSessionManager(cookieName string, maxlifetime int64) (*SessionManager, error) {
	sid := make(map[string]Session)
	return &SessionManager{cookieName: cookieName, sid: sid, maxlifetime: maxlifetime}, nil
}

//GetSID 生成Session唯一ID
func (sessionManager *SessionManager) GetSID(name string) (sid string) {
	sid = xcm.GetMD5(name) + xcm.GetRandomString(8, xcm.NSDSTR)
	return
}

//GC 定时对过期的Session进行删除
func (sessionManager *SessionManager) GC() {
	sessionManager.lock.Lock()
	defer sessionManager.lock.Unlock()
	for k, v := range sessionManager.sid {
		if (v.TimeAccessed.Unix() + sessionManager.maxlifetime) < time.Now().Unix() {
			delete(sessionManager.sid, k)
		}
	}
	time.AfterFunc(time.Duration(sessionManager.maxlifetime*2), func() { sessionManager.GC() })
}

//Add 添加Session
func (sessionManager *SessionManager) Add(sid string, gid int) error {
	sessionManager.lock.Lock()
	defer sessionManager.lock.Unlock()
	zs := Session{}
	zs.TimeAccessed = time.Now()
	zs.GroupID = gid
	sessionManager.sid[sid] = zs
	return nil
}

//Get 获取Session
func (sessionManager *SessionManager) Get(sid string) (zs Session, err error) {
	err = nil
	zs, ok := sessionManager.sid[sid]
	if !ok {
		err = errors.New("no session")
	}
	return
}

//Del 删除Session
func (sessionManager *SessionManager) Del(sid string) error {
	sessionManager.lock.Lock()
	defer sessionManager.lock.Unlock()
	delete(sessionManager.sid, sid)
	return nil
}

//Update 更新Session最后访问时间
func (sessionManager *SessionManager) Update(sid string) error {
	sessionManager.lock.Lock()
	defer sessionManager.lock.Unlock()
	zs, _ := sessionManager.Get(sid)
	zs.TimeAccessed = time.Now()
	sessionManager.sid[sid] = zs
	return nil
}

//TimeOut 判断Session是否过期
func (sessionManager *SessionManager) TimeOut(sid string) bool {
	if _, ok := sessionManager.sid[sid]; !ok {
		return true
	}
	if (sessionManager.sid[sid].TimeAccessed.Unix() + sessionManager.maxlifetime) < time.Now().Unix() {
		return true
	}
	return false
}

//GetCookie 获取客户端返回的Cookie
func (sessionManager *SessionManager) GetCookie(r *http.Request, cookieName string) (sid string, err error) {
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
func (sessionManager *SessionManager) SetCookie(w http.ResponseWriter, cookieName, sid string) {
	cookie := http.Cookie{Name: cookieName, Value: sid, Path: "/", HttpOnly: true}
	http.SetCookie(w, &cookie)
	return
}

//DeleteCookie 设置客户端返回的Cookie
func (sessionManager *SessionManager) DeleteCookie(w http.ResponseWriter, cookieName, sid string) {
	cookie := http.Cookie{Name: cookieName, Value: sid, Path: "/", HttpOnly: true, MaxAge: -1}
	http.SetCookie(w, &cookie)
	return
}

//GetGroupID 获取群组ID
func (sessionManager *SessionManager) GetGroupID(sid string) (gid int, err error) {
	err = nil
	zs, ok := sessionManager.sid[sid]
	if !ok {
		err = errors.New("no session")
		return
	}
	gid = zs.GroupID
	return
}
