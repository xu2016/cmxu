package xsql

import (
	_ "gopkg.in/goracle.v2"
)

//XOracle 对Oracle的操作
type XOracle struct {
	DB  string
	SQL string
}
