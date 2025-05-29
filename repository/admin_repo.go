package repository

import (
	"fmt"
	"self_go_gin/dao"
	"self_go_gin/model"

	"gorm.io/gorm"
)

type AdminRepositoryInterface interface {
	GetDB() *gorm.DB
	GetAdminByAccount(account string) (*model.Admins, error)
	CreateAdmin(newAdmin model.Admins) (*model.Admins, error)
}

type adminRepositoryImpl struct {
	dao dao.AdminDaoInterface
}

func NewAdminRepository() AdminRepositoryInterface {
	return &adminRepositoryImpl{
		dao: dao.NewAdminDao(),
	}
}

func (r *adminRepositoryImpl) GetDB() *gorm.DB {
	return r.dao.GetGenericDao().GetDB()
}

func (r *adminRepositoryImpl) GetAdminByAccount(account string) (*model.Admins, error) {
	log_data := map[string]interface{}{
		"account": account,
	}
	admin, err := r.dao.GetAdminByAccount(account)
	if err != nil {
		return nil, fmt.Errorf("AdminRepositoryImpl GetAdminByAccount() data: %s \n %w", log_data, err)
	}

	return admin, err
}

func (r *adminRepositoryImpl) CreateAdmin(newAdmin model.Admins) (*model.Admins, error) {
	log_data := map[string]interface{}{
		"newAdmin": newAdmin,
	}
	admin, err := r.dao.GetGenericDao().Create(&newAdmin)
	if err != nil {
		return nil, fmt.Errorf("AdminRepositoryImpl CreateAdmin() data: %s \n %w", log_data, err)
	}
	return admin, nil
}
