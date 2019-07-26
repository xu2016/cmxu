package kdjson

/* 宽带订单管理相关JSON */
//OpDataJSON 宽带订单领取和展示返回的订单JSON定义
type OpDataJSON struct {
	Logid   string      `json:"logid"`
	Xdtime  string      `json:"xdtime"`
	Ddtype  string      `json:"ddtype"`
	City    string      `json:"city"`
	Lcnum   string      `json:"lcnum"`
	Ddstate string      `json:"ddstate"`
	Opname  string      `json:"opname"`
	Opphone string      `json:"Opphone"`
	DdInfo  interface{} `json:"ddInfo"`
	User    interface{} `json:"user"`
	Product interface{} `json:"product"`
	Jwd     interface{} `json:"jwd"`
	Log     interface{} `json:"log"`
}

//NewDdData 新订单(未指定操作人，需要指派人员)
type NewDdData struct {
	Key      string `json:"key"`
	LOGID    string `json:"logid"`    //订单号
	CONTACT  string `json:"contact"`  //客户姓名
	INSDATE  string `json:"insdate"`  //下单时间
	PLANNAME string `json:"planname"` //预约套餐
	WFID     string `json:"wfid"`     //订单类型编号
	WF       string `json:"wf"`       //订单类型
	WFJDID   string `json:"wfjdid"`   //所处流程编号
	DDSTATE  string `json:"ddstate"`  //所处流程
}

//NewDdJSON 获取所属地市所有未领取订单返回JSON
type NewDdJSON struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data []NewDdData `json:"data"`
}

//MyDdJSON 获取我的订单
type MyDdJSON struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data []RMyDD `json:"data"`
}
type RMyDD struct {
	Wfid   string `json:"wfid"`
	Wfjdid int    `json:"wfjdid"`
	Logid  string `json:"logid"`
}

//PdDdInfoJSON
type PdDdInfoJSON struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	Kdczinfo CKdczinfo `json:"kdczinfo"`
	Ddinfo   CDdinfo   `json:"ddinfo"`
	Fzrinfo  CFzrinfo  `json:"fzrinfo"`
	Kdinfo   CKdinfo   `json:"kdinfo"`
	Zjinfo   CZjinfo   `json:"zjinfo"`
	Jwdinfo  CJwdinfo  `json:"jwdinfo"`
}

//CKdczinfo ...
type CKdczinfo struct {
	NoZpsm      string `json:"noZpsm"`
	Zpsm        string `json:"Zpsm"`        //
	KdddDisplay string `json:"kdddDisplay"` //
	Czlx        string `json:"czlx"`        //
	Czbz        string `json:"czbz"`        //
	Wfid        string `json:"wfid"`        //
	Wfjdid      int    `json:"wfjdid"`      //
}

//CDdinfo 订单基本信息
type CDdinfo struct {
	Logid     string `json:"logid"`     //订单编号
	Username  string `json:"username"`  //客户姓名
	Userphone string `json:"userphone"` //联系电话
	Xdtime    string `json:"xdtime"`    //下单时间
	Yytime    string `json:"yytime"`    //预约时间
	Zjdz      string `json:"zjdz"`      //装机地址
	Bcdz      string `json:"bcdz:`      //补充地址
}

type CFzrinfo struct {
	Qdmc     string `json:"qdmc"`     //渠道名称
	Bmid     string `json:"bmid"`     //渠道编码
	Zbid     string `json:"zbid"`     //总部渠道编码
	Zbgh     string `json:"zbgh"`     //总部工号
	Zbname   string `json:"zbname"`   //总部姓名
	Bdgh     string `json:"bdgh"`     //本地工号
	Ghname   string `json:"ghname"`   //工号姓名
	Fzrphone string `json:"fzrphone"` //联系号码
}

type CKdinfo struct {
	Kfphone string `json:"kfphone"` //开户号码
	Yytc    string `json:"yytc"`    //预约套餐
	Kdsl    string `json:"kdsl"`    //宽带速率
	Tcfy    string `json:"tcfy"`    //套餐费用
	Azfy    string `json:"azfy"`    //安装费用
	Xzhsf   string `json:"xzhsf"`   //先装后收费
	Gmxq    string `json:"gmxq"`    //光猫需求
	Gmfy    string `json:"gmfy"`    //光猫费用
	Iptvxq  string `json:"iptvxq"`  //IPTV需求
	Iptvfy  string `json:"iptvfy"`  //IPTV费用
	Jdhxq   string `json:"jdhxq"`   //机顶盒需求
	Jdhfy   string `json:"jdhfy"`   //机顶盒费用
	Zjbz    string `json:"zjbz"`    //装机备注
}
type CZjinfo struct {
	Zjlx  string `json:"zjlx"`  //证件类型
	Zjhm  string `json:"zjhm"`  //证件号码
	Zjyxq string `json:"zjyxq"` //证件有效期
	Zjdz  string `json:"zjdz"`  //证件地址
}

type CJwdinfo struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

//KdProXXX 宽带套餐明细
type KdProXXX struct {
	Proid   string  `json:"proid"`
	Proname string  `json:"proname"`
	KDSL    float64 `json:"KDSL"`
	PROFEE  float64 `json:"PROFEE"`
	AZFEE   float64 `json:"AZFEE"`
	IPTVFEE float64 `json:"IPTVFEE"`
	GMFEE   float64 `json:"GMFEE"`
	JDHFEE  float64 `json:"JDHFEE"`
	REMARK  string  `json:"remark"`
}

//XdRJSON 下单返回信息
type XdRJSON struct {
	Code   int             `json:"code"`
	Msg    string          `json:"msg"`
	Openid string          `json:"openid"`
	Ddid   string          `json:"ddid"`
	Proxx  KdProXXX        `json:"proxx"`
	Data   []NearPointInfo `json:"data"`
	Remark string          `json:"remark"`
}

//NearPointInfo 就近资源点信息
type NearPointInfo struct {
	Rn   int     `json:"rn"`
	Jl   float64 `json:"jl"`
	Sbid string  `json:"sbid"`
	Sbwd float64 `json:"sbwd"`
	Sbjd float64 `json:"sbjd"`
}

//XdqyJSON 下单确应返回
type XdqyJSON struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	OpenID string `json:"openId"`
	DDID   string `json:"ddId"`
}

//NewDdCntJSON 返回待派订单总数
type NewDdCntJSON struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Cnt  int    `json:"cnt"`
}
