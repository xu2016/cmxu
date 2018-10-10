package xsql

import (
	"database/sql"
	"log"
)

//QueryLine 查询一条记录
func (xdb *CXSql) QueryLine() (row *sql.Row, err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	row = db.QueryRow(xdb.sql)
	return
}

//Query 查询多条记录
func (xdb *CXSql) Query() (rows *sql.Rows, err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		log.Println("Open err:", err)
		return
	}
	defer db.Close()
	rows, err = db.Query(xdb.sql)
	if err != nil {
		log.Println("Query err:", err)
		return
	}
	return
}
