package xtsz

//SubMenus 子菜单
type SubMenus struct {
	Key  string `json:"Key"`
	Name string `json:"Name"`
}

//Menu 菜单
type Menu struct {
	Key      string     `json:"Key"`
	Name     string     `json:"Name"`
	IconType string     `json:"IconType"`
	SubMenu  []SubMenus `json:"subMenu"`
}

//RMenu 菜单返回JSON
type RMenu struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	MenuMain []Menu `json:"menuMain"`
}

//Menus 菜单变量
var Menus = []Menu{
	{
		Key:      `ddgl`,
		Name:     `订单管理`,
		IconType: `file-text`,
		SubMenu: []SubMenus{
			{Key: `ddcx`, Name: `订单查询`},
			{Key: `htxd`, Name: `后台下单`},
			{Key: `ddxg`, Name: `订单修改`},
			{Key: `ztpd`, Name: `中台派单`},
			{Key: `ddjk`, Name: `订单监控`},
			{Key: `ddxz`, Name: `订单下载`},
		},
	},
	{
		Key:      `zygl`,
		Name:     `资源管理`,
		IconType: `deployment-unit`,
		SubMenu: []SubMenus{
			{Key: `zycx`, Name: `资源查询`},
			{Key: `zyxg`, Name: `资源修改`},
			{Key: `zyxz`, Name: `资源下载`},
		},
	},
	{
		Key:      `sjfx`,
		Name:     `数据分析`,
		IconType: `pie-chart`,
		SubMenu:  []SubMenus{},
	},
	{
		Key:      `qxsz`,
		Name:     `权限设置`,
		IconType: `lock`,
		SubMenu: []SubMenus{
			{Key: `cpgl`, Name: `产品管理`},
			{Key: `yhgl`, Name: `用户管理`},
		},
	},
}
var xtsz = Menu{
	Key:      `xtsz`,
	Name:     `系统设置`,
	IconType: `setting`,
	SubMenu: []SubMenus{
		{Key: `mmxg`, Name: `密码修改`},
		{Key: `zlxg`, Name: `资料修改`},
	},
}

//MenuSort 一级菜单的排序
var MenuSort = map[string]int{`ddgl`: 1, `zygl`: 2, `sjfx`: 3, `qxsz`: 4}

func Getmenu(grpids map[string]int) (rmenu []Menu) {
	rmenu = make([]Menu, 0)
	xm := make([]int, len(MenuSort))
	for k, v := range grpids {
		xm[MenuSort[k]-1] = v
	}
	for k, v := range xm {
		if v != 0 {
			xmm := Menu{Key: Menus[k].Key, Name: Menus[k].Name, IconType: Menus[k].IconType, SubMenu: make([]SubMenus, 0)}
			xn := len(Menus[k].SubMenu)
			for i := 0; i < xn; i++ {
				if ((v >> uint(i)) & 1) == 1 {
					xmm.SubMenu = append(xmm.SubMenu, Menus[k].SubMenu[i])
				}
			}
			rmenu = append(rmenu, xmm)
		}
	}
	rmenu = append(rmenu, xtsz)
	return
}
