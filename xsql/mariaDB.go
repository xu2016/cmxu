package xsql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//XMariaDB 对MySQL和MariaDB的操作
type XMariaDB struct {
	DB  string
	SQL string
}

//QueryLine 查询一条记录
func (xdb *XMariaDB) QueryLine() (row *sql.Row, err error) {
	db, err := sql.Open("mysql", xdb.DB)
	if err != nil {
		return
	}
	defer db.Close()
	row = db.QueryRow(xdb.SQL)
	return
}

//Query 查询多条记录
func (xdb *XMariaDB) Query() (rows *sql.Rows, err error) {
	db, err := sql.Open("mysql", xdb.DB)
	if err != nil {
		return
	}

	defer db.Close()
	rows, err = db.Query(xdb.SQL)
	return
}
