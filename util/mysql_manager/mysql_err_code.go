package mysql_manager

import (
	"strconv"

	mysqlErr "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

const DuplicateEntryCode = 1062

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

func CheckRecordNotFound(result *gorm.DB) error {
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func StringToUint64(s string) (uint64, error) {
	ui64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return ui64, nil
}

func StringToUint16(s string) (uint16, error) {
	ui64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	ui16 := uint16(ui64)
	return ui16, nil
}
