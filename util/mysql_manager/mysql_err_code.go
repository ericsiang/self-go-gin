package mysql_manager

import (
	mysqlErr "github.com/go-sql-driver/mysql"
)

// MysqlErrCode 根据mysql错误信息返回错误代码
/*
* 1062: Duplicate entry
*/
func MysqlErrCode(err error) int {
	mysqlErr, ok := err.(*mysqlErr.MySQLError)
	if !ok {
		return 0
	}
	return int(mysqlErr.Number)
}
