package kdjson

/* 用户管理相关JSON */
//RUserInfo ...
type RUserInfo struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data []UserInfo `json:"data"`
}

//UserInfo 用户信息
type UserInfo struct {
	UserID   string         `json:"userid"`
	UserName string         `json:"username"`
	Phone    string         `json:"phone"`
	City     string         `json:"city"`
	Yf       string         `json:"yf"`
	Qd       string         `json:"qd"`
	DeptID   string         `json:"deptid"` //部门ID
	BSSGH    string         `json:"bssgh"`  //BSS七位工号编码
	OAID     string         `json:"oaid"`
	GRPID    map[string]int `json:"grpid"`
}
type GRUserInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data []GUserInfo `json:"data"`
}

type GUserInfo struct {
	UserName string `json:"username"`
	Phone    string `json:"phone"`
	City     string `json:"city"`
}

type GORUserInfo struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data GOUserInfo `json:"data"`
}

type GOUserInfo struct {
	UserID   string `json:"userid"`
	UserName string `json:"username"`
	Phone    string `json:"phone"`
	City     string `json:"city"`
	Yf       string `json:"yf"`
	Qd       string `json:"qd"`
	DeptID   string `json:"deptid"` //部门ID
	BSSGH    string `json:"bssgh"`  //BSS七位工号编码
	OAID     string `json:"oaid"`
	Ddgl     []int  `json:"ddgl"`
	Zygl     []int  `json:"zygl"`
	Sjfx     []int  `json:"sjfx"`
	Qxsz     []int  `json:"qxsz"`
}

//LoginJSON 登陆返回信息
type LoginJSON struct {
	Code   int            `json:"code"`
	Msg    string         `json:"msg"`
	Userid string         `json:"userid"`
	Phone  string         `json:"phone"`
	City   string         `json:"city"`
	Gid    map[string]int `json:"gid"`
}
