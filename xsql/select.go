package xsql

import (
	"database/sql"
	"log"
)

//QueryLine 查询一条记录
func (xdb *CXSql) QueryLine(qstr string) (row *sql.Row, err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	row = db.QueryRow(qstr)
	return
}

//Query 查询多条记录
func (xdb *CXSql) Query(qstr string) (rows *sql.Rows, err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		log.Println("Open err:", err)
		return
	}
	defer db.Close()
	rows, err = db.Query(qstr)
	if err != nil {
		log.Println("Query err:", err)
		log.Println(qstr)
		return
	}
	return
}
