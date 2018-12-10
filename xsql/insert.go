package xsql

import (
	"database/sql"
	"log"
)

//Insertline 插入一条记录
func (xdb *CXSql) Insertline(qstr string) (err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(qstr)
	if err != nil {
		log.Println(qstr)
		return
	}
	if xdb.cmt {
		_, err = db.Exec(`commit`)
	}
	return
}
