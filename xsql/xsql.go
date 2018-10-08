package xsql

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "gopkg.in/goracle.v2"
)

//CXSql 数据库访问类型
type CXSql struct {
	dbtype string
	db     string
	sql    string
	cmt    bool
}

func NewSQL(sdbtype, sdb, ssql string, scmt bool) (xdb CXSql) {
	xdb = CXSql{dbtype: sdbtype, db: sdb, sql: ssql, cmt: scmt}
	return
}

//RJSON insert和update返回的JSON
type RJSON struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//Init ...
func (rp RJSON) Init() {
	rp.Code = 1
	rp.Msg = "失败"
}

//QJSON 查询时返回的JSON
type QJSON struct {
	NmInt    map[string]sql.NullInt64
	NmFloat  map[string]sql.NullFloat64
	NmString map[string]sql.NullString
	NmBool   map[string]sql.NullBool
}

//Init ...
func (rr QJSON) Init() {
	rr.NmInt = make(map[string]sql.NullInt64, 0)
	rr.NmFloat = make(map[string]sql.NullFloat64, 0)
	rr.NmString = make(map[string]sql.NullString, 0)
	rr.NmBool = make(map[string]sql.NullBool, 0)
}

//QRowsJSON 查询时返回的JSON
type QRowsJSON struct {
	Code int
	Msg  string
	Data []QJSON
}

//Init ...
func (rp QRowsJSON) Init() {
	rp.Code = 1
	rp.Msg = "查询失败"
	rp.Data = make([]QJSON, 0)
}

//GetFloat64 sql.NullFloat64转化为float64类型
func GetFloat64(null sql.NullFloat64, dft float64) (res float64) {
	res = dft
	if null.Valid {
		res = null.Float64
	}
	return
}

//GetString sql.NullString转化为string类型
func GetString(null sql.NullString, dft string) (res string) {
	res = dft
	if null.Valid {
		res = null.String
	}
	return
}

//GetBool sql.NullBool转化为bool类型
func GetBool(null sql.NullBool, dft bool) (res bool) {
	res = dft
	if null.Valid {
		res = null.Bool
	}
	return
}

//GetInt64 sql.NullInt64转化为int64类型
func GetInt64(null sql.NullInt64, dft int64) (res int64) {
	res = dft
	if null.Valid {
		res = null.Int64
	}
	return
}

//QIsEmpty 判断查询结果是否为空
func QIsEmpty(rp QRowsJSON) (bl bool) {
	bl = false
	if len(rp.Data) != 0 {
		bl = true
	}
	return
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
func QRows(dbtype, dbstr string, cmt bool, r *http.Request) (rp *QRowsJSON) {
	rp = &QRowsJSON{}
	rp.Init()
	rr := QJSON{}
	rr.Init()
	cstr := strings.Split(r.FormValue("qcstr"), ",")
	tstr := strings.Split(r.FormValue("qtstr"), ",")
	ci, err := RTypeSlice(tstr)
	if err != nil {
		rp.Msg = err.Error()
		log.Println("RTypeSlice:", err)
		return
	}
	db := NewSQL(dbtype, dbstr, r.FormValue("qstr"), cmt)
	rows, err := db.Query()
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
				rr.NmInt[cstr[i]] = *v.(*sql.NullInt64)
			case *sql.NullFloat64:
				rr.NmFloat[cstr[i]] = *v.(*sql.NullFloat64)
			case *sql.NullString:
				rr.NmString[cstr[i]] = *v.(*sql.NullString)
			case *sql.NullBool:
				rr.NmBool[cstr[i]] = *v.(*sql.NullBool)
			}
		}
		rp.Data = append(rp.Data, rr)
	}
	rp.Code = 0
	rp.Msg = "查询成功"
	return
}
