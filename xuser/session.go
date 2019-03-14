package xuser

import (
	"errors"
	"net/http"
	"time"
)

//ZSession session存储结构
type ZSession struct {
	uid          string         //用户账号
	timeAccessed time.Time      //最后访问时间
	gid          map[string]int //群组ID
	city         string         //地市
}

//GetUID 获取用户账号
func (zs *ZSession) GetUID() (uid string) {
	uid = zs.uid
	return
}

//GetGIDS 获取群组IDS
func (zs *ZSession) GetGIDS() (gid map[string]int) {
	gid = zs.gid
	return
}

//GetGID 获取特定群组ID
func (zs *ZSession) GetGID(grptp string) (gid int) {
	gid = zs.gid[grptp]
	return
}

//UserContrl 用户权限控制
func (zs *ZSession) UserContrl(grptp string, subMenuNum int) (b bool) {
	if ((zs.gid[grptp] >> uint(subMenuNum)) & 1) == 1 {
		b = true
	}
	return
}

//GetCity 获取用户所属地市
func (zs *ZSession) GetCity() (city string) {
	city = zs.city
	return
}

//GetSession 获取并验证账号session是否存在、正确和未过期，
//如果存在且正确并未过期，更新session,gid为用户角色（权限)。
func GetSession(r *http.Request) (zs ZSession, err error) {
	sid, err := GSession.GetCookie(r, CName)
	if err != nil {
		return
	}
	zs, err = GSession.Get(sid)
	if err != nil {
		return
	}
	if GSession.TimeOut(sid) {
		err = errors.New("TimeOut")
		return
	}
	GSession.Update(sid)
	return
}

//DelSession 删除session
func DelSession(w http.ResponseWriter, r *http.Request) {
	sid, err := GSession.GetCookie(r, CName)
	if err != nil {
		return
	}
	GSession.Del(sid)
	GSession.DeleteCookie(w, CName, sid)
	return
}

//AddSession 添加session
func AddSession(w http.ResponseWriter, sid string) {
	GSession.SetCookie(w, CName, sid)
	return
}
