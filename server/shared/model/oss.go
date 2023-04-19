package model

import "gorm.io/gorm"

type GoodsOSSRecord struct {
	gorm.Model
	GoodsId int64
	Path    string `gorm:"column:path;type:varchar(100);not null"`
}

type HeadOSSRecord struct {
	gorm.Model
	UserId int64
	Path   string `gorm:"column:path;type:varchar(100);not null"`
}
