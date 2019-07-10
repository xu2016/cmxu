package xsqlx

import (
	"database/sql"
	"log"
)

/*QueryLine 查询一条记录
  parameter：
	qstr：查询语句，里面包含一些需要替换的值在val中,替换标识号取决于不同的数据库
		  MySQL：?,?,....,?
		  Oracle：:1,:2,.....,:n
		注意：在qstr中要处理好可能出现NULL的情况，比如在Oracle中，nvl(字段，为NULL是的值)
	val：要替换的值的slice
	colnum：需要取的字段长度。
  return：返回结果是一个字符串slice，请获取到数据后自行进行数据转换,操作错误err。
*/
func (xdb *CXSqlx) QueryLine(qstr string, val []interface{}, colnum int) (gdata []string, err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	sci := make([]interface{}, colnum)
	for i := range sci {
		sci[i] = new(string)
	}
	row := db.QueryRow(qstr, val...)
	if err != nil {
		log.Println("Query Error:", err)
		log.Println("Query Sql:", qstr)
		return
	}
	err = row.Scan(sci...)
	if err != nil {
		return
	}
	ds := make([]string, colnum)
	for i, v := range sci {
		ds[i] = *v.(*string)
	}
	gdata = ds
	return
}

/*Query 查询多条记录
  parameter：
	qstr：查询语句，里面包含一些需要替换的值在val中,替换标识号取决于不同的数据库
		  MySQL：?,?,....,?
		  Oracle：:1,:2,.....,:n
		注意：在qstr中要处理好可能出现NULL的情况，比如在Oracle中，nvl(字段，为NULL是的值)
	val：要替换的值的slice
	colnum：需要取的字段长度。
  return：返回结果是一个二维的字符串slice，请获取到数据后自行进行数据转换。
*/
func (xdb *CXSqlx) Query(qstr string, val []interface{}, colnum int) (gdata [][]string, err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	sci := make([]interface{}, colnum)
	for i := range sci {
		sci[i] = new(string)
	}
	rows, err := db.Query(qstr, val...)
	if err != nil {
		log.Println("Query Error:", err)
		log.Println("Query Sql:", qstr)
		return
	}
	for rows.Next() {
		err = rows.Scan(sci...)
		if err != nil {
			continue
		}
		ds := make([]string, colnum)
		for i, v := range sci {
			ds[i] = *v.(*string)
		}
		gdata = append(gdata, ds)
	}
	return
}
