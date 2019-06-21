package xuser

import "sync"

//Guflag 全局唯一性标示管理器
var Guflag *UniFlag

//UniFlag 唯一性标示管理器
type UniFlag struct {
	lock        sync.Mutex      // protects session
	yzms        map[string]CYzm //号码验证码
	maxlifetime int64
}
