package xcm

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type wxcodeInfo struct {
	Openid     string `json:"openid"`
	Sessionkey string `json:"session_key"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

/*GetOpenID 用于获取用户的微信openid
功能：用于获取用户的微信openid
参数：
	url：https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
		参数		必填	说明
		appid		是	小程序唯一标识
		secret		是	小程序的 app secret
		js_code		是	登录时获取的 code
		grant_type	是	填写为 authorization_code
返回：
	openid：成功返回用户微信openid，失败返回空字符串
	err:成功返回nil，失败返回错误
*/
func GetOpenID(wxAppID, wxAppSecret, wxcode string) (openid string, err error) {
	var wx wxcodeInfo
	url := `https://api.weixin.qq.com/sns/jscode2session?appid=` + wxAppID + `&secret=` + wxAppSecret + `&js_code=` + wxcode + `&grant_type=authorization_code`
	resp, err := http.Get(url)
	openid = "0"
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(body), &wx)
	if wx.Openid == "" {
		log.Println(wx.Errmsg)
		openid = "0"
		err = errors.New(wx.Errmsg)
	} else {
		openid = wx.Openid
		err = nil
	}
	return
}
