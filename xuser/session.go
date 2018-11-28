package xuser

import (
	"time"
)

//ZSession session存储结构
type ZSession struct {
	uid          string    //用户账号
	timeAccessed time.Time //最后访问时间
	gid          int       //群组ID
}

//GetUID 获取用户账号
func (zs *ZSession) GetUID() (uid string) {
	uid = zs.uid
	return
}

//GetGID 获取用户账号
func (zs *ZSession) GetGID() (gid int) {
	gid = zs.gid
	return
}
