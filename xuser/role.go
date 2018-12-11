package xuser

import (
	"cmxu/xsql"
	"cmxu/xwb"
	"log"
)

//GRole 全局角色
var GRole map[string]Role

//Role 角色管理
type Role struct {
	rname string
	menu  []LMenu
	inlmt map[string]int
}

//XMenu 菜单
type XMenu struct {
	Xmid   string //菜单ID
	Xrw    int    //排序序号
	Xmname string //菜单名称
	Xmurl  string //URL地址
	Xumid  string //上级菜单
}

//LMenu 菜单链
type LMenu struct {
	Menu XMenu //我的菜单
	MNum int   //菜单级别
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
func NewRole() (r map[string]Role) {
	r = make(map[string]Role)
	mm := readMenu()
	m := menuLevel(mm)
	rr := readRole()
	rm := readRoleMenu()
	for i, v := range rr {
		rl := Role{}
		rl.inlmt = make(map[string]int)
		rl.rname = v
		for _, vv := range rm {
			if vv.role == i {
				rl.inlmt[vv.menu] = 1
			}
		}
		rl.menu = m
		r[i] = rl
	}
	return
}

//GetRoleName 返回该角色的名称
func (r *Role) GetRoleName() (rn string) {
	rn = r.rname
	return
}

//GetMenu 返回该角色的菜单
func (r *Role) GetMenu() (xm []LMenu) {
	xm = r.menu
	return
}

//GetInLmt 获取角色是不是有接口访问权限
func (r *Role) GetInLmt(iid string) (b bool) {
	if r.inlmt[iid] == 1 || iid == "xu20181210" {
		b = true
	}
	return
}

//readMenu 读取所有有效菜单
func readMenu() (menu map[string]XMenu) {
	menu = make(map[string]XMenu)
	qstr := `select OBJID,RW,OBJNAME,CURL,UP_OBJID from ZSKDZS_OBJINFO where STATE='1' and OBJTYPE='menu'`
	qcstr := `OBJID,RW,OBJNAME,CURL,UP_OBJID`
	qtstr := `string,int,string,string,string`
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
			Xmurl:  xsql.GetString(v.NmString["CURL"], "null"),
			Xumid:  xsql.GetString(v.NmString["UP_OBJID"], "null"),
		}
		menu[xm.Xmid] = xm
	}
	return
}

//readRole 读取所有角色
func readRole() (role map[string]string) {
	role = make(map[string]string)
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
func readRoleMenu() (rolemenu []RoleMenu) {
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
func menuLevel(menu map[string]XMenu) (m []LMenu) {
	m = make([]LMenu, 0)
	temp := make([]LMenu, 0)
	for i, v := range menu {
		if v.Xumid == "null" {
			xm := LMenu{Menu: v, MNum: 0}
			temp = append(temp, xm)
			delete(menu, i)
		}
	}
	//log.Println(temp)
	temp = menuSort(temp)
	for _, vv := range temp {
		temp1 := make([]LMenu, 0)
		for i, v := range menu {
			if v.Xumid == vv.Menu.Xmid {
				xm := LMenu{Menu: v, MNum: 1}
				temp1 = append(temp1, xm)
				delete(menu, i)
			}
		}
		temp1 = menuSort(temp1)
		m = append(m, vv)
		for _, v := range temp1 {
			m = append(m, v)
		}
	}
	return
}

//menuSort 同级菜单排序
func menuSort(xm []LMenu) (m []LMenu) {
	for i := 0; i < len(xm); i++ {
		mm := xm[i]
		for j := i + 1; j < len(xm); j++ {
			if mm.Menu.Xrw > xm[j].Menu.Xrw {
				mm, xm[j] = xm[j], mm
			}
		}
		xm[i] = mm
	}
	m = xm
	return
}
