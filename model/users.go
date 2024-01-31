package model

type Users struct {
	GormModel
	Account  string `gorm:"type:varchar(255);not null;uniqueIndex;" json:"account" binding:"required"`
	Password string `gorm:"type:varchar(255);not null;" json:"password" binding:"required"`
}

func (model Users) CreateUser() (Users, error) {
	err := db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (model Users) GetUsersById(UserId int64) (Users, error) {
	err := db.Where("id = ?", UserId).First(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (model Users) GetUsersByAccount(account string) (Users, error) {
	err := db.Where("account = ?", account).First(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
