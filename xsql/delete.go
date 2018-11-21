package xsql

import (
	"database/sql"
	"log"
)

//Deleteline 删除一条记录
func (xdb *CXSql) Deleteline(qstr string) (err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(qstr)
	if err != nil {
		log.Println(err)
		log.Println(qstr)
		return
	}
	if xdb.cmt {
		_, err = db.Exec(`commit`)
	}
	return
}
