package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name        string `gorm:"type:varchar(40)"`
	UserId      int64
	User        User `gorm:"foreignKey:UserId"`
	Description string
	Cost        int64
	SalesVolume int64
}
