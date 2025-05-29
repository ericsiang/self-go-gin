package migrate

import (
	"self_go_gin/initialize"
	"self_go_gin/model"
)

var err error

func Migrate() {
	err = initialize.GetMysqlDB().AutoMigrate(&model.User{})
	panicErr(err)
	err = initialize.GetMysqlDB().AutoMigrate(&model.Admins{})
	panicErr(err)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
