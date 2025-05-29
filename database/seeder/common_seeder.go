package seeder

import (
	"self_go_gin/initialize"
	"self_go_gin/model"
	"self_go_gin/util/bcryptEncap"
	"strconv"
)

func CreateUser() {
	db := initialize.GetMysqlDB()
	seeder := NewSeeder(db)
	if err := seeder.Clear("users"); err != nil {
		panic(err)
	}
	var users []*model.User
	//密碼加密
	bcryptPassword, err := bcryptEncap.GenerateFromPassword("123456")
	if err != nil {
		panic("Seeder CreateUser() bcrypt fail")
	}

	for i := 1; i < 4; i++ {
		users = append(users, &model.User{
			Account:  "user" + strconv.Itoa(i),
			Password: string(bcryptPassword),
		})
	}

	err = db.Create(&users).Error
	if err != nil {
		panic("Seeder CreateUser() Create fail")
	}
}

func CreateAdmin() {
	db := initialize.GetMysqlDB()
	seeder := NewSeeder(db)
	if err := seeder.Clear("admins"); err != nil {
		panic(err)
	}
	var admins []*model.Admins
	//密碼加密
	bcryptPassword, err := bcryptEncap.GenerateFromPassword("123456")
	if err != nil {
		panic("Seeder CreateAdmin() bcrypt fail")
	}

	for i := 1; i < 4; i++ {
		admins = append(admins, &model.Admins{
			Account:  "admin" + strconv.Itoa(i),
			Password: string(bcryptPassword),
		})
	}

	err = db.Create(&admins).Error
	if err != nil {
		panic("Seeder CreateUser() Create fail")
	}
}
