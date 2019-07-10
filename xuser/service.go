package xuser

//CService 服务管理类
type CService struct {
	DBType    string
	DBString  string
	DBTable   string
	DBCommint bool
}

//QueryService 服务查询
func (cs *CService) QueryService() (err error) {

	return
}

//AddService 服务添加
func (cs *CService) AddService(val []interface{}) (err error) {
	//istr:=`insert into ` +cs.DBTable+`(field1,field2,...fieldN)VALUES(value1,value2,...valueN)`
	//xdb := xsqlx.NewSQLX(cs.DBType, cs.DBString, cs.DBCommint)

	return
}

//DelService 服务删除
func (cs *CService) DelService() (err error) {

	return
}

//ChangeService 服务变更
func (cs *CService) ChangeService() (err error) {

	return
}
