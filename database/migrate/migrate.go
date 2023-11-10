package migrate

import (
	"api/initialize"
	"api/model"

)

func Migrate(){
	initialize.GetMysqlDB().AutoMigrate(&model.Users{})
}