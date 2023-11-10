package model

type Users struct {
	GormModel
	Account  string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password string `gorm:"type:varchar(255);not null"`
}

func (model Users) GetUsersById(UserId int) (Users, error) {
	err := db.Where("id = ?", UserId).First(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
