package xwb

import (
	"database/sql"
	"log"
	"strconv"
)

//ZBZYInfo ...
type ZBZYInfo struct {
	Sbid       string  `json:"sbid"`
	Yf         string  `json:"yf"`
	Dy         string  `json:"dy"`
	Sbjd       float64 `json:"sbjd"`
	Sbwd       float64 `json:"sbwd"`
	Slsb       string  `json:"slsb"`
	Wydks      int     `json:"wydks"`
	Sbaddress  string  `json:"sbaddress"`
	Fgfw       string  `json:"fgfw"`
	Statename  string  `json:"statename"`
	Sfxcflag   int     `json:"sfxcflag"`
	Sfxc       string  `json:"sfxc"`
	Sfxnflag   string  `json:"sfxnflag"`
	Sfxn       string  `json:"sfxn"`
	Updateflag int     `json:"updateflag"`
	Zwlx       string  `json:"zwlx"`
	Tjlx       string  `json:"tjlx"`
	Juli       float64 `json:"juli"`
}

//Zbzyss 资源搜索
func Zbzyss(lng, lat, juli float64, cnt int, database string, ctbl string) (zbzy []ZBZYInfo, err error) {
	//log.Println(lng, lat, juli, cnt, database)
	zbzy = make([]ZBZYInfo, 0)
	db, err := sql.Open("mysql", database)
	if err != nil {
		log.Println("Zbzyss:", err)
		return
	}
	defer db.Close()
	jstr := `ROUND(6378.138*2*ASIN(SQRT(POW(SIN((` + strconv.FormatFloat(lat, 'f', -1, 64) + `*PI()/180-wsbbdwd*PI()/180)/2),2)+COS(` + strconv.FormatFloat(lat, 'f', -1, 64) + `*PI()/180)*COS(wsbbdwd*PI()/180)*POW(SIN((` + strconv.FormatFloat(lng, 'f', -1, 64) + `*PI()/180-wsbbdjd*PI()/180)/2),2)))*1000)`
	cntstr, julistr := "", ""
	if juli > 0 {
		julistr = ` and ` + jstr + `<` + strconv.FormatFloat(juli, 'f', -1, 64)
	}
	if cnt != 0 {
		cntstr = ` LIMIT ` + strconv.Itoa(cnt)
	}
	qstr := `SELECT sbid,yf,accessdy,wsbbdjd,wsbbdwd,slsb,wydks,sbaddress,fgfw,statename,sfxcflag,sfxc,sfxnflag,sfxn,updateflag,zwlx,tjlx,` +
		jstr + `AS juli FROM ` + ctbl + ` WHERE statename='正常'` + julistr + ` ORDER BY juli ASC` + cntstr
	//log.Println(qstr)
	rows, err := db.Query(qstr)
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		zy := ZBZYInfo{}
		err = rows.Scan(&zy.Sbid, &zy.Yf, &zy.Dy, &zy.Sbjd, &zy.Sbwd, &zy.Slsb, &zy.Wydks, &zy.Sbaddress, &zy.Fgfw, &zy.Statename, &zy.Sfxcflag, &zy.Sfxc, &zy.Sfxnflag, &zy.Sfxn, &zy.Updateflag, &zy.Zwlx, &zy.Tjlx, &zy.Juli)
		if err == nil {
			if zy.Sfxnflag == "2" {
				zy.Sbaddress = "*"
				zy.Sbid = "ZS-ODB-YHJG-*"
			}
			zbzy = append(zbzy, zy)
		}
	}
	err = nil
	return
}
