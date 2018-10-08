package xsql

import (
	"database/sql"

	_ "gopkg.in/goracle.v2"
)

//XOracle 对Oracle的操作
type XOracle struct {
}

//QueryLine 查询一条记录
func (xdb *XOracle) QueryLine(dbstr, sqlstr string) (row *sql.Row, err error) {
	db, err := sql.Open("goracle", dbstr)
	if err != nil {
		return
	}
	defer db.Close()
	row = db.QueryRow(sqlstr)
	return
}

//Query 查询多条记录
func (xdb *XOracle) Query(dbstr, sqlstr string) (rows *sql.Rows, err error) {
	db, err := sql.Open("goracle", dbstr)
	if err != nil {
		return
	}

	defer db.Close()
	rows, err = db.Query(sqlstr)
	return
}

//Insertline 插入一条记录
func (xdb *XOracle) Insertline(dbstr, sqlstr string) (err error) {
	db, err := sql.Open("goracle", dbstr)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(sqlstr)
	if err != nil {
		return
	}
	_, err = db.Exec(`commit`)
	return
}

//Updateline 更新一条记录
func (xdb *XOracle) Updateline(dbstr, sqlstr string) (err error) {
	db, err := sql.Open("goracle", dbstr)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(sqlstr)
	if err != nil {
		return
	}
	_, err = db.Exec(`commit`)
	return
}
