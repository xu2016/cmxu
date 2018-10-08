package xsql

import (
	"database/sql"
)

//Updateline 更新一条记录
func (xdb *CXSql) Updateline() (err error) {
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
