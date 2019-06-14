package xsqlx

//XSQLX 数据库访问类型
type XSQLX struct {
	dbtype string
	db     string
	cmt    bool
}

//NewXSQLX 创建一个数据库对象
func NewXSQLX(sdbtype, sdb string, scmt bool) (xdb XSQLX) {
	xdb = XSQLX{dbtype: sdbtype, db: sdb, cmt: scmt}
	return
}
