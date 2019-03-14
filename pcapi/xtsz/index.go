package xtsz

import "net/http"

//Rt 路由
var Rt map[string]func(http.ResponseWriter, *http.Request)

//Menus 菜单变量
var Menus = map[string]Menu{
	`ddgl`: {
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
	`zygl`: {
		Key:      `zygl`,
		Name:     `资源管理`,
		IconType: `deployment-unit`,
		SubMenu: []SubMenus{
			{Key: `zycx`, Name: `资源查询`},
			{Key: `zyxg`, Name: `资源修改`},
			{Key: `zyxz`, Name: `资源下载`},
		},
	},
	`sjfx`: {
		Key:      `sjfx`,
		Name:     `数据分析`,
		IconType: `pie-chart`,
		SubMenu:  []SubMenus{},
	},
	`qxsz`: {
		Key:      `qxsz`,
		Name:     `权限设置`,
		IconType: `lock`,
		SubMenu: []SubMenus{
			{Key: `cpgl`, Name: `产品管理`},
			{Key: `yhgl`, Name: `用户管理`},
		},
	},
	`xtsz`: {
		Key:      `xtsz`,
		Name:     `系统设置`,
		IconType: `setting`,
		SubMenu: []SubMenus{
			{Key: `mmxg`, Name: `密码修改`},
			{Key: `zlxg`, Name: `资料修改`},
		},
	},
}

func init() {
	Rt = make(map[string]func(http.ResponseWriter, *http.Request))
	//系统设置
	Rt["/login"] = login   //登陆
	Rt["/logout"] = logout //用户注销
	Rt["/reg"] = reg       //注册
	Rt["/mmxg"] = mmxg     //密码修改
	Rt["/zlxg"] = zlxg     //资料修改
}
