package seeder

import (
	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func RunSeeder() {
	//common_seeder
	// CreateUser()
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) Clear(tableName string) error {
	return s.db.Exec("truncate table " + tableName).Error
}





