package initialize

import (
	"api/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

var db *gorm.DB

func initMysql() {
	var config *gorm.Config
	gormZaplogger := zapgorm2.New(zap.L())
	logger.Default.LogMode(logger.Error)
	// zap.S().Info("logger level: ", logger.Info)
	// zap.S().Info("ori_loggger : ", ori_loggger)
	// zap.S().Info("gormZaplogger : ", gormZaplogger)
	if gin.Mode() == gin.ReleaseMode {
		config = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   true,
		}
	} else {
		config = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   true,
			Logger : gormZaplogger,
		}
	}

	//注意：User和Password为MySQL資料庫的管理員密碼，Host和Port為資料庫連接ip端口，DBname為要連接的資料庫
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", User,Password,Ip,Port,DBName)
	dsn := ServerEnv.MysqlDB.Username + ":" + ServerEnv.MysqlDB.Password + "@tcp(" + ServerEnv.MysqlDB.Host + ":" + strconv.Itoa((ServerEnv.MysqlDB.Port)) + ")/" + ServerEnv.MysqlDB.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// zap.L().Info("dsn :"+ dsn)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		zap.S().Error("mysql connect error :"+ err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	model.SetDB(db)
}

func GetMysqlDB() *gorm.DB {
	return db
}
