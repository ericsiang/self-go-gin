package dao

import (
    "fmt"
    "self_go_gin/initialize"
    "self_go_gin/model"
)

type AdminDaoInterface interface {
    GetGenericDao() GenericDaoInterface[model.Admins]
    GetAdminByAccount(account string) (*model.Admins, error)
}

type adminDaoImpl struct {
    GenericDao GenericDaoInterface[model.Admins]
}

func NewAdminDao() AdminDaoInterface {
    return &adminDaoImpl{
        GenericDao: NewGenericDAO[model.Admins](initialize.GetMysqlDB()),
    }
}

func (d *adminDaoImpl) GetGenericDao() GenericDaoInterface[model.Admins] {
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
