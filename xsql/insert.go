package xsql

import (
	"database/sql"
)

//Insertline 插入一条记录
func (xdb *CXSql) Insertline() (err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(xdb.sql)
	if err != nil {
		return
	}
	if xdb.cmt {
		_, err = db.Exec(`commit`)
	}
	return
}
