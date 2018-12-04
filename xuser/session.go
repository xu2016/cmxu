package xuser

import (
	"cmxu/xsql"
	"log"
	"time"
)

//ZSession session存储结构
type ZSession struct {
	uid          string    //用户账号
	timeAccessed time.Time //最后访问时间
	gid          string    //群组ID
}

//GetUID 获取用户账号
func (zs *ZSession) GetUID() (uid string) {
	uid = zs.uid
	return
}

//GetGID 获取用户账号
func (zs *ZSession) GetGID() (gid string) {
	gid = zs.gid
	return
}

//SetUserGid ...
func SetUserGid() {
	m := make(map[string]int)
	o := make(map[string]string)
	var phone, uid string
	var gid int
	db := xsql.NewSQL("mysql", "root:VcXzSa_1328a@tcp(157.122.176.134)/zslt?charset=utf8", false)
	rows, err := db.Query("select phone,gid from zslt_tbl")
	for rows.Next() {
		err = rows.Scan(&phone, &gid)
		if err == nil {
			m[phone] = gid
		}
	}

	dba := xsql.NewSQL("goracle", "kdzy/kdzy_20180705@pdbzsdmc", true)
	rows, err = dba.Query("select USERID,PHONE from ZSKDZS_USER")
	for rows.Next() {
		err = rows.Scan(&uid, &phone)
		if err == nil {
			o[phone] = uid
		}
	}
	for i, v := range o {
		gidx := "101"
		if m[i] == 1 {
			gidx = "100"
		}
		qstr := `insert into ZSKDZS_USERINFO (USERID,GRPID) values ('` + v + `','` + gidx + `')`
		err := dba.Insertline(qstr)
		log.Println(err)
	}
}
