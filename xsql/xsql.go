package xsql

import (
	"database/sql"
	"errors"
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

//QJSON 查询时返回的数据JSON
type QJSON struct {
	NmInt    map[string]sql.NullInt64   `json:"nmint"`
	NmFloat  map[string]sql.NullFloat64 `json:"nmfloat"`
	NmString map[string]sql.NullString  `json:"nmstring"`
	NmBool   map[string]sql.NullBool    `json:"nmbool"`
}

//QRowsJSON 查询时返回的JSON
type QRowsJSON struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data []QJSON `json:"data"`
}

//RTypeSlice 返回确定类型的空接口切片
func RTypeSlice(tstr []string) (ci []interface{}, err error) {
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
	rp.Data = make([]QJSON, 0)
	cstr := strings.Split(r.FormValue("qcstr"), ",")
	tstr := strings.Split(r.FormValue("qtstr"), ",")
	ci, err := RTypeSlice(tstr)
	if err != nil {
		rp.Msg = err.Error()
		return
	}
	var db XDB
	switch dbtype {
	case "mysql":
		db = &XMariaDB{}
	case "oracle":
		db = &XOracle{}
	}
	sqlstr := "select " + r.FormValue("qcstr") + r.FormValue("qstr")
	qjsn := QJSON{}

	rows, err := db.Query(dbstr, sqlstr)
	err = rows.Scan(ci...)
	if err != nil {
		rp.Msg = err.Error()
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(ci...)
		if err != nil {
			rp.Msg = err.Error()
			return
		}
		for i, v := range ci {
			switch v.(type) {
			case sql.NullInt64:
				qjsn.NmInt[cstr[i]] = *v.(*sql.NullInt64)
			case sql.NullFloat64:
				qjsn.NmFloat[cstr[i]] = *v.(*sql.NullFloat64)
			case sql.NullString:
				qjsn.NmString[cstr[i]] = *v.(*sql.NullString)
			case sql.NullBool:
				qjsn.NmBool[cstr[i]] = *v.(*sql.NullBool)
			}
		}
		rp.Data = append(rp.Data, qjsn)
	}
	rp.Code = 0
	rp.Msg = "查询成功"
	qi = rp
	return
}
