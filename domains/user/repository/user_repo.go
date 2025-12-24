package repository

import (
	"fmt"
	"self_go_gin/domains/user/entity/model"
	"self_go_gin/domains/user/repository/dao"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetDB() *gorm.DB
	GetUsersByAccount(account string) (*model.User, error)
	CreateUser(newUser *model.User) (*model.User, error)
}

type userRepositoryImpl struct {
	dao dao.UserDaoInterface
}

func NewUserRepository() UserRepositoryInterface {
	return &userRepositoryImpl{
		dao: dao.NewUserDao(),
	}
}

func (r *userRepositoryImpl) GetDB() *gorm.DB {
	return r.dao.GetGenericDao().GetDB()
}

func (r *userRepositoryImpl) GetUsersByAccount(account string) (*model.User, error) {
	log_data := map[string]interface{}{
		"account": account,
	}
	user, err := r.dao.GetUsersByAccount(account)
	if err != nil {
		return nil, fmt.Errorf("UserRepositoryImpl GetUsersByAccount() data: %s \n %w", log_data, err)
	}

	return user, err
}

func (r *userRepositoryImpl) CreateUser(newUser *model.User) (*model.User, error) {
	log_data := map[string]interface{}{
		"newUser": newUser,
	}
	user, err := r.dao.GetGenericDao().Create(newUser)
	if err != nil {
		return nil, fmt.Errorf("UserRepositoryImpl CreateUser() data: %s \n %w", log_data, err)
	}
	return user, nil
}
