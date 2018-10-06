package xsql

import (
	"database/sql"

	_ "gopkg.in/goracle.v2"
)

//XOracle 对Oracle的操作
type XOracle struct {
	DB  string
	SQL string
}

//QueryLine 查询一条记录
func (xdb *XOracle) QueryLine() (row *sql.Row, err error) {
	db, err := sql.Open("goracle", xdb.DB)
	if err != nil {
		return
	}
	defer db.Close()
	row = db.QueryRow(xdb.SQL)
	return
}

//Query 查询多条记录
func (xdb *XOracle) Query() (rows *sql.Rows, err error) {
	db, err := sql.Open("goracle", xdb.DB)
	if err != nil {
		return
	}

	defer db.Close()
	rows, err = db.Query(xdb.SQL)
	return
}

//Insertline 插入一条记录
func (xdb *XOracle) Insertline() (err error) {
	db, err := sql.Open("goracle", xdb.DB)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(xdb.SQL)
	if err != nil {
		return
	}
	_, err = db.Exec(`commit`)
	return
}

//Updateline 更新一条记录
func (xdb *XOracle) Updateline() (err error) {
	db, err := sql.Open("goracle", xdb.DB)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(xdb.SQL)
	if err != nil {
		return
	}
	_, err = db.Exec(`commit`)
	return
}
