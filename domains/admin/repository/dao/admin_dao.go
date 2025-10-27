package dao

import (
    "fmt"
    "self_go_gin/domains/admin/entity/model"
	"self_go_gin/internal/dao"
    "self_go_gin/infra/orm/gorm_mysql"

)

type AdminDaoInterface interface {
    GetGenericDao() dao.GenericDaoInterface[model.Admins]
    GetAdminByAccount(account string) (*model.Admins, error)
}

type adminDaoImpl struct {
    GenericDao dao.GenericDaoInterface[model.Admins]
}

func NewAdminDao() AdminDaoInterface {
    return &adminDaoImpl{
        GenericDao: dao.NewGenericDAO[model.Admins](gorm_mysql.GetMysqlDB()),
    }
}

func (d *adminDaoImpl) GetGenericDao() dao. GenericDaoInterface[model.Admins] {
    return d.GenericDao
}

func (d *adminDaoImpl) GetAdminByAccount(account string) (*model.Admins, error) {
    log_data := map[string]interface{}{
        "account": account,
    }
    var admin model.Admins
    err := d.GenericDao.GetDB().Where("account =?", account).First(&admin).Error
    if err != nil {
        return nil, fmt.Errorf("AdminDaoImpl GetAdminByAccount() data: %s \n %w", log_data, err)
    }
    return &admin, err
}
