package dao

import (
	"fmt"
	"self_go_gin/infra/orm/gorm_mysql"
	"self_go_gin/domains/user/entity/model"
	"self_go_gin/internal/dao"
)

type UserDaoInterface interface {
	GetGenericDao() dao.GenericDaoInterface[model.User]
	GetUsersByAccount(account string) (*model.User, error)
}

type userDaoImpl struct {
	GenericDao dao.GenericDaoInterface[model.User]
}

func NewUserDao() UserDaoInterface {
	return &userDaoImpl{
		GenericDao: dao.NewGenericDAO[model.User](gorm_mysql.GetMysqlDB()),
	}
}

func (d *userDaoImpl) GetGenericDao() dao.GenericDaoInterface[model.User] {
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
