package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name        string `gorm:"type:varchar(40)"`
	Describe    string
	Cost        int64
	SalesVolume int64
}
