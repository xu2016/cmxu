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

func getmenu(userid map[string]int) (rmenu []Menu, err error) {

	return
}
