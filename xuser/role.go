package xuser

import (
	"cmxu/xsql"
	"cmxu/xwb"
	"log"
)

//Role 角色管理
type Role struct {
	rname string
	menu  *LMenu
	inlmt map[string]int
}

//XMenu 菜单
type XMenu struct {
	Xmid   string //菜单ID
	Xrw    int    //排序序号
	Xmname string //菜单名称
	Xmtype string //菜单类型
	Xmurl  string //URL地址
	Xumid  string //上级菜单
}

//LMenu 菜单链
type LMenu struct {
	Menu  *XMenu //我的菜单
	CMenu *LMenu //首个子菜单链接
	NMenu *LMenu //下一个兄弟菜单链接
}

//RoleMenu 角色菜单关系表
type RoleMenu struct {
	role string
	menu string
}

/*NewRole 角色初始化
mdb:菜单数据库,mtb:菜单数据表, mtp:菜单数据库类型
rdb:角色数据库, rtb:角色数据表, rtp:角色数据库类型
rmdb:角色菜单关系数据库, rmtb:角色菜单关系数据表, rmtp:角色菜单关系数据库类型
*/
func NewRole(mdb, mtb, mtp, rdb, rtb, rtp, rmdb, rmtb, rmtp string) (r []Role) {

	return
}

//GetRoleName 返回该角色的名称
func (r *Role) GetRoleName() (rn string) {
	rn = r.rname
	return
}

//GetMenu 返回该角色的菜单
func (r *Role) GetMenu() (xm *LMenu) {
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

//readMenu 读取所有有效菜单
func readMenu(mdb, mtb, mtp string) (menu map[string]XMenu) {
	qstr := `select OBJID,RW,OBJNAME,OBJTYPE,CURL,UP_OBJID from ZSKDZS_OBJINFO where STATE='1'`
	qcstr := `OBJID,RW,OBJNAME,OBJTYPE,CURL,UP_OBJID`
	qtstr := `string,int,string,string,string,string`
	rr := xsql.QRowsJSON{}
	rr.Init()
	_, err := xwb.XPost("http://211.95.193.112:8080/sqlms/qrow", qstr, qcstr, qtstr, &rr)
	if err != nil {
		log.Println(err)
		return
	}
	if xsql.QIsEmpty(rr) {
		log.Println("读取所有有效菜单失败")
		return
	}
	for _, v := range rr.Data {
		xm := XMenu{
			Xmid:   xsql.GetString(v.NmString["OBJID"], "null"),
			Xrw:    int(xsql.GetInt64(v.NmInt["RW"], 0)),
			Xmname: xsql.GetString(v.NmString["OBJNAME"], "null"),
			Xmtype: xsql.GetString(v.NmString["OBJTYPE"], "null"),
			Xmurl:  xsql.GetString(v.NmString["CURL"], "null"),
			Xumid:  xsql.GetString(v.NmString["UP_OBJID"], "null"),
		}
		menu[xm.Xmid] = xm
	}
	return
}

//readRole 读取所有角色
func readRole(rdb, rtb, rtp string) (role map[string]string) {
	qstr := `select GRPID,GRPNAME from ZSKDZS_GROUP`
	qcstr := `GRPID,GRPNAME`
	qtstr := `string,string`
	rr := xsql.QRowsJSON{}
	rr.Init()
	_, err := xwb.XPost("http://211.95.193.112:8080/sqlms/qrow", qstr, qcstr, qtstr, &rr)
	if err != nil {
		log.Println(err)
		return
	}
	if xsql.QIsEmpty(rr) {
		log.Println("读取所有角色失败")
		return
	}
	for _, v := range rr.Data {
		role[xsql.GetString(v.NmString["GRPID"], "null")] = xsql.GetString(v.NmString["GRPNAME"], "null")
	}
	return
}

//readRoleMenu 读取所有角色菜单关系
func readRoleMenu(rmdb, rmtb, rmtp string) (rolemenu []RoleMenu) {
	qstr := `select GRPID,OBJID from ZSKDZS_GROUPING`
	qcstr := `GRPID,OBJID`
	qtstr := `string,string`
	rr := xsql.QRowsJSON{}
	rr.Init()
	_, err := xwb.XPost("http://211.95.193.112:8080/sqlms/qrow", qstr, qcstr, qtstr, &rr)
	if err != nil {
		log.Println(err)
		return
	}
	if xsql.QIsEmpty(rr) {
		log.Println("读取所有角色失败")
		return
	}
	for _, v := range rr.Data {
		xrm := RoleMenu{
			role: xsql.GetString(v.NmString["GRPID"], "null"),
			menu: xsql.GetString(v.NmString["OBJID"], "null"),
		}
		rolemenu = append(rolemenu, xrm)
	}
	return
}

//menuLevel 菜单分层
func menuLevel(menu map[string]XMenu) (m *LMenu) {
	xmenu := make(map[string]XMenu)
	//按顺序添加主菜单
	for i, v := range menu {
		//判断是不是菜单
		if v.Xmtype == "menu" {
			//判断是不是主菜单
			if v.Xumid == "null" {
				xm := &LMenu{Menu: &v}
				//如果主菜单为nil或主菜单第一个节点比新节点大
				if m == nil || m.Menu.Xrw > v.Xrw {
					xm.NMenu = m
					m = xm
				} else {
					for p := m; p != nil; p = p.NMenu {
						if p.Menu.Xrw < v.Xrw && (p.NMenu.Menu.Xrw > v.Xrw || p.NMenu == nil) {
							xm.NMenu = p.NMenu
							p.NMenu = xm
							break
						}
					}
				}
			} else {
				xmenu[i] = v
			}
		}
	}
	//添加除主菜单外的所有子菜单
	for {

	}
	return
}
