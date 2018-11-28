package xuser

//Role 角色管理
type Role struct {
	rid   string
	menu  []XMenu
	inlmt map[string]int
}

//XMenu 菜单
type XMenu struct {
	Xmid    string
	Xrw     int
	Xmname  string
	Xmtype  string
	Xmurl   string
	Xumid   string
	Xmstate string
}

/*NewRole 角色初始化
mdb:菜单数据库,mtb:菜单数据表, mtp:菜单数据库类型
rdb:角色数据库, rtb:角色数据表, rtp:角色数据库类型
rmdb:角色菜单关系数据库, rmtb:角色菜单关系数据表, rmtp:角色菜单关系数据库类型
*/
func NewRole(mdb, mtb, mtp, rdb, rtb, rtp, rmdb, rmtb, rmtp string) (r []Role) {

	return
}

//GetMenu 返回该角色的菜单
func (r *Role) GetMenu() (xm []XMenu) {
	xm = r.menu
	return
}

//GetInLmt 获取角色是不是有接口访问权限
func (r *Role) GetInLmt(iid string) (b bool) {
	if r.inlmt[iid] == 1 {
		b = true
	}
	return
}
