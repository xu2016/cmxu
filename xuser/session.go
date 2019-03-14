package xuser

import (
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

//GetGID 获取群组ID
func (zs *ZSession) GetGID() (gid map[string]int) {
	gid = zs.gid
	return
}

//GetCity 获取用户所属地市
func (zs *ZSession) GetCity() (city string) {
	city = zs.city
	return
}
