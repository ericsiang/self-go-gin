package migrate

import (
	"api/initialize"
	"api/model"
)

var err error

func Migrate() {
	err = initialize.GetMysqlDB().AutoMigrate(&model.Users{})
	panicErr(err)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
