package kdjson

//DscxInfo ...
type DscxInfo struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

//ZYInfo 返回的资源信息结构体
type ZYInfo struct {
	City       string  `json:"city"`       //归属地市
	Sbid       string  `json:"sbid"`       //设备ID
	Yf         string  `json:"yf"`         //归属营服
	Dy         string  `json:"dy"`         //归属单元
	Sbjd       float64 `json:"sbjd"`       //设备经度（可返回Bd和Gcj）
	Sbwd       float64 `json:"sbwd"`       //设备纬度（可返回Bd和Gcj）
	Wydks      int     `json:"wydks"`      //可用端口数
	Sbaddress  string  `json:"sbaddress"`  //设备地址
	Fgfw       string  `json:"fgfw"`       //覆盖范围
	Statename  string  `json:"statename"`  //设备状态
	Sfxcflag   int     `json:"sfxcflag"`   //清查类型
	Sfxc       string  `json:"sfxc"`       //清查类型标志
	Sfxnflag   string  `json:"sfxnflag"`   //设备类型
	Sfxn       string  `json:"sfxn"`       //设备类型标志
	Updateflag int     `json:"updateflag"` //更新标记
	Zwlx       string  `json:"zwlx"`       //装维类型
	Tjlx       string  `json:"tjlx"`       //投建类型
	Juli       float64 `json:"juli"`       //距离
}
type Zyx []ZYInfo

func (c Zyx) Len() int {
	return len(c)
}
func (c Zyx) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Zyx) Less(i, j int) bool {
	return c[i].Juli < c[j].Juli
}

//ZyWXInfo 返回附近资源信息，主要用于小程序
type ZyWXInfo struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Openid string   `json:"openid"`
	Cxid   string   `json:"cxid"`
	Data   []ZYInfo `json:"data"`
}

//ZyPCInfo 返回附近资源信息，主要用于PC
type ZyPCInfo struct {
	Code   int        `json:"code"`
	Msg    string     `json:"msg"`
	Openid string     `json:"openid"`
	Lng    float64    `json:"lng"`
	Lat    float64    `json:"lat"`
	Sdz    string     `json:"sdz"`
	Cxid   string     `json:"cxid"`
	Data   []ZYInfo   `json:"data"`
	Fjzy   []FJZYInfo `json:"fjzy"`
}

//FJZYInfo 附近PC端最近5个点的距离和步行距离
type FJZYInfo struct {
	Sbid   string  `json:"sbid"`
	Juli   float64 `json:"juli"`
	Bxjuli float64 `json:"bxjuli"`
}

//WgcxInfo 查询返回的JSON
type WgcxInfo struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	OpenID string   `json:"openid"`
	Data   []ZYInfo `json:"data"`
}
