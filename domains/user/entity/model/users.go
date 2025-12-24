package model

import (
	"self_go_gin/internal/model"
)

type User struct {
	model.GormModel
	Account  string `gorm:"type:varchar(255);not null;uniqueIndex;" json:"account" binding:"required"`
	Password string `gorm:"type:varchar(255);not null;" json:"password" binding:"required"`
}
