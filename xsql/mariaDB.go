package xsql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//XMariaDB 对MySQL和MariaDB的操作
type XMariaDB struct {
}

//QueryLine 查询一条记录
func (xdb *XMariaDB) QueryLine(dbstr, sqlstr string) (row *sql.Row, err error) {
	db, err := sql.Open("mysql", dbstr)
	if err != nil {
		return
	}
	defer db.Close()
	row = db.QueryRow(sqlstr)
	return
}

//Query 查询多条记录
func (xdb *XMariaDB) Query(dbstr, sqlstr string) (rows *sql.Rows, err error) {
	db, err := sql.Open("mysql", dbstr)
	if err != nil {
		return
	}

	defer db.Close()
	rows, err = db.Query(sqlstr)
	return
}

//Insertline 插入一条记录
func (xdb *XMariaDB) Insertline(dbstr, sqlstr string) (err error) {
	db, err := sql.Open("mysql", dbstr)
	if err != nil {
		return
	}

	defer db.Close()
	_, err = db.Exec(sqlstr)
	return
}

//Updateline 更新一条记录
func (xdb *XMariaDB) Updateline(dbstr, sqlstr string) (err error) {
	db, err := sql.Open("mysql", dbstr)
	if err != nil {
		return
	}

	defer db.Close()
	_, err = db.Exec(sqlstr)
	return
}
