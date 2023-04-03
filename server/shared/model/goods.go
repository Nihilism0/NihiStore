package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name        string `gorm:"type:varchar(40)"`
	User_Id     int64
	User        User `gorm:"foreignKey:User_Id"`
	Describe    string
	Cost        int64
	SalesVolume int64
}
