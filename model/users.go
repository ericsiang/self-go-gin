package model

type User struct {
	GormModel
	Account  string `gorm:"type:varchar(255);not null;uniqueIndex;" json:"account" binding:"required"`
	Password string `gorm:"type:varchar(255);not null;" json:"password" binding:"required"`
}

func (model User) CreateUser() (User, error) {
	err := db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (model User) GetUsersById(UserId int64) (User, error) {
	err := db.Where("id = ?", UserId).First(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (model User) GetUsersByAccount(account string) (User, error) {
	err := db.Where("account = ?", account).First(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
