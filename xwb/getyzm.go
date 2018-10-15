package xwb

import (
	"bytes"
	"cmxu/xcm"
	"cmxu/xsql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//YzmInfo 返回验证码信息
type YzmInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Yzm  string `json:"yzm"`
}

//YzmJkInfo 返回验证码信息
type YzmJkInfo struct {
	Result int    `json:"result"`
	Errmsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Fee    int    `json:"fee"`
	Sid    string `json:"sid"`
}

//Getyzm 生成并发送验证码
func Getyzm(w http.ResponseWriter, r *http.Request, sdkappid, appkey, database, random string) {
	rp := &YzmInfo{}
	wx := &YzmJkInfo{}
	rp.Code = 1
	rp.Msg = "获取验证码失败"
	rp.Yzm = ""
	tm := strconv.FormatInt(time.Now().Unix(), 10)
	str := `appkey=` + appkey + `&random=` + random + `&time=` + tm + `&mobile=` + r.FormValue("phone")
	sig := xcm.GetSHA256(str)
	jsonStr := []byte(`{
		"ext": "",
		"extend": "",
		"params": [
			"` + random + `",
			"1"
		],
		"sig": "` + sig + `",
		"sign": "中山资源核查工具",
		"tel": {
			"mobile": "` + r.FormValue("phone") + `",
			"nationcode": "86"
		},
		"time": ` + tm + `,
		"tpl_id": 162364
	}`)
	url := `https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=` + sdkappid + `&random=` + random
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		rp.Msg = err.Error()
		JSONPage(rp, w)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rp.Msg = err.Error()
		JSONPage(rp, w)
		return
	}
	json.Unmarshal([]byte(body), &wx)
	if wx.Result != 0 {
		rp.Msg = wx.Errmsg
		JSONPage(rp, w)
		return
	}
	rp.Code = 0
	rp.Msg = "发送验证码成功"
	rp.Yzm = random
	JSONPage(rp, w)
	return
}

//SendYzm 发送验证码，msg必须以中文开头，且只能发送联通手机号码
func SendYzm(phone, msg string) (err error) {
	qstr := `insert into smagent.t_smsend@sm16(smseq, sender, receiver,  message, getdate, EXPIREDDATE,  commiter, pri)
	select SMAGENT.smseq.NEXTVAL@sm16, 760133, '` + phone + `','` + msg + `', SYSDATE-0.01, sysdate+2/24, '760', 1 from dual a`
	rr := xsql.RJSON{}
	rr.Init()
	_, err = XPost("http://211.95.193.112:8080/sqlms/insertline", qstr, "", "", &rr)
	return
}

//SetYZM 设置验证码
func SetYZM(phone, msg, dbstr, sql string) (err error) {
	db := xsql.NewSQL("mysql", dbstr, false)
	err = db.Insertline("insert " + sql)
	if err != nil {
		err = db.Updateline("update " + sql)
		if err != nil {
			return
		}
	}
	err = SendYzm(phone, msg)
	return
}

//YzYzm 验证验证码
func YzYzm(yzm, dbstr, sql string) (bl bool) {
	db := xsql.NewSQL("mysql", dbstr, false)
	row, err := db.QueryLine(sql)
	if err != nil {
		return
	}
	var dbyzm string
	err = row.Scan(&dbyzm)
	if err != nil {
		return
	}
	if yzm == dbyzm {
		bl = true
	}
	return
}
