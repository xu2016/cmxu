package xsql

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"
)

//XDB 数据库访问接口
type XDB interface {
	QueryLine(dbstr, sqlstr string) (row *sql.Row, err error)
	Query(dbstr, sqlstr string) (rows *sql.Rows, err error)
	Insertline(dbstr, sqlstr string) (err error)
	Updateline(dbstr, sqlstr string) (err error)
}

//QRowsJSON 查询时返回的JSON
type QRowsJSON struct {
	Code     int
	Msg      string
	NmInt    map[string]sql.NullInt64
	NmFloat  map[string]sql.NullFloat64
	NmString map[string]sql.NullString
	NmBool   map[string]sql.NullBool
}

//RTypeSlice 返回确定类型的空接口切片
func RTypeSlice(tstr []string) (ci []interface{}, err error) {
	ci = make([]interface{}, len(tstr))
	for i, v := range tstr {
		switch v {
		case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint", "NullInt64":
			ci[i] = new(sql.NullInt64)
		case "float32", "float64", "NullFloat64":
			ci[i] = new(sql.NullFloat64)
		case "string", "NullString":
			ci[i] = new(sql.NullString)
		case "bool", "NullBool":
			ci[i] = new(sql.NullBool)
		default:
			err = errors.New("Type not implements")
		}
	}
	return
}

//QRows 查询数据库
func QRows(dbtype, dbstr string, r *http.Request) (qi interface{}, err error) {
	rp := &QRowsJSON{}
	rp.Code = 1
	rp.Msg = "查询失败"
	//rp.Data = make([]QJSON, 0)
	rp.NmInt = make(map[string]sql.NullInt64, 0)
	rp.NmFloat = make(map[string]sql.NullFloat64, 0)
	rp.NmString = make(map[string]sql.NullString, 0)
	rp.NmBool = make(map[string]sql.NullBool, 0)
	cstr := strings.Split(r.FormValue("qcstr"), ",")
	tstr := strings.Split(r.FormValue("qtstr"), ",")
	ci, err := RTypeSlice(tstr)
	if err != nil {
		rp.Msg = err.Error()
		log.Println("RTypeSlice:", err)
		return
	}
	var db XDB
	switch dbtype {
	case "mysql":
		db = &XMariaDB{}
	case "oracle":
		db = &XOracle{}
	}
	rows, err := db.Query(dbstr, r.FormValue("qstr"))
	if err != nil {
		rp.Msg = err.Error()
		log.Println("Query:", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(ci...)
		if err != nil {
			rp.Msg = err.Error()
			log.Println("rows.Scan:", err)
			return
		}
		for i, v := range ci {
			switch v.(type) {
			case *sql.NullInt64:
				rp.NmInt[cstr[i]] = *v.(*sql.NullInt64)
			case *sql.NullFloat64:
				rp.NmFloat[cstr[i]] = *v.(*sql.NullFloat64)
			case *sql.NullString:
				rp.NmString[cstr[i]] = *v.(*sql.NullString)
			case *sql.NullBool:
				rp.NmBool[cstr[i]] = *v.(*sql.NullBool)
			}
		}
		//rp.Data = append(rp.Data, qjsn)
	}
	rp.Code = 0
	rp.Msg = "查询成功"
	qi = rp
	return
}
