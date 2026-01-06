package gorm_mysql

import (
	"fmt"
	"os"
	"self_go_gin/infra/env"
	"self_go_gin/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

var db *gorm.DB
// InitMysql 初始化 MySQL 資料庫連接
func InitMysql(GetServerEnv func() *env.ServerConfig) {
	var config *gorm.Config
	gormZaplogger := zapgorm2.New(zap.L())
	logger.Default.LogMode(logger.Error)
	// zap.S().Info("logger level: ", logger.Info)
	// zap.S().Info("ori_loggger : ", ori_loggger)
	// zap.S().Info("gormZaplogger : ", gormZaplogger)
	if gin.Mode() == gin.ReleaseMode {
		config = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			// SkipDefaultTransaction:                   true,
		}
	} else {
		config = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			// SkipDefaultTransaction:                   true,
			Logger: gormZaplogger,
		}
	}

	//注意：User和Password为MySQL資料庫的管理員密碼，Host和Port為資料庫連接ip端口，DBname為要連接的資料庫
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", User,Password,Ip,Port,DBName)
	dsn := GetServerEnv().MysqlDB.Username + ":" + GetServerEnv().MysqlDB.Password + "@tcp(" + GetServerEnv().MysqlDB.Host + ":" + strconv.Itoa((GetServerEnv().MysqlDB.Port)) + ")/" + GetServerEnv().MysqlDB.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// zap.L().Info("dsn :"+ dsn)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mysql connect failed, err:", err)
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	model.SetDB(db)
	fmt.Println("mysql connect success")
}

// GetMysqlDB 返回 MySQL 資料庫連接
func GetMysqlDB() *gorm.DB {
	return db
}
