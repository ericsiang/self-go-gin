package model

import "gorm.io/gorm"

var db *gorm.DB

func SetDB(gorm *gorm.DB) {
	db = gorm
}
