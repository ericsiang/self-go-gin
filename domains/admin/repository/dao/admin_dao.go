package dao

import (
	"fmt"
	"self_go_gin/domains/admin/entity/model"
	"self_go_gin/infra/orm/gorm_mysql"
	"self_go_gin/internal/dao"
)

// AdminDaoInterface 管理員帳號密碼表 DAO 介面
type AdminDaoInterface interface {
	GetGenericDao() dao.GenericDaoInterface[model.Admins]
	GetAdminByAccount(account string) (*model.Admins, error)
}

type adminDaoImpl struct {
	GenericDao dao.GenericDaoInterface[model.Admins]
}

// NewAdminDao 建立管理員帳號密碼表 DAO
func NewAdminDao() AdminDaoInterface {
	return &adminDaoImpl{
		GenericDao: dao.NewGenericDAO[model.Admins](gorm_mysql.GetMysqlDB()),
	}
}

func (d *adminDaoImpl) GetGenericDao() dao.GenericDaoInterface[model.Admins] {
	return d.GenericDao
}

// GetAdminByAccount 根據帳號查詢管理員
func (d *adminDaoImpl) GetAdminByAccount(account string) (*model.Admins, error) {
	logData := map[string]interface{}{
		"account": account,
	}
	var admin model.Admins
	err := d.GenericDao.GetDB().Where("account =?", account).First(&admin).Error
	if err != nil {
		return nil, fmt.Errorf("AdminDaoImpl GetAdminByAccount() data: %s \n %w", logData, err)
	}
	return &admin, err
}
