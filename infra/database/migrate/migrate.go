package migrate

import (
	admin_model "self_go_gin/domains/admin/entity/model"
	user_model "self_go_gin/domains/user/entity/model"
	"self_go_gin/infra/orm/gorm_mysql"
)

var err error

func Migrate() {
	err = gorm_mysql.GetMysqlDB().AutoMigrate(&user_model.User{})
	panicErr(err)
	err = gorm_mysql.GetMysqlDB().AutoMigrate(&admin_model.Admins{})
	panicErr(err)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
