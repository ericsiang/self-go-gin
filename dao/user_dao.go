package dao

import (
	"fmt"
	"self_go_gin/initialize"
	"self_go_gin/model"
)

type UserDaoInterface interface {
	GetGenericDao() GenericDaoInterface[model.User]
	GetUsersByAccount(account string) (*model.User, error)
}

type userDaoImpl struct {
	GenericDao GenericDaoInterface[model.User]
}

func NewUserDao() UserDaoInterface {
	return &userDaoImpl{
		GenericDao: NewGenericDAO[model.User](initialize.GetMysqlDB()),
	}
}

func (d *userDaoImpl) GetGenericDao() GenericDaoInterface[model.User] {
	return d.GenericDao
}
func (d *userDaoImpl) GetUsersByAccount(account string) (*model.User, error) {
	log_data := map[string]interface{}{
		"account": account,
	}
	var user model.User
	err := d.GenericDao.GetDB().Where("account =?", account).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("UserDaoImpl GetUsersByAccount() data: %s \n %w", log_data, err)
	}
	return &user, err
}
