package model

import "gorm.io/gorm"

type Favorites struct {
	gorm.Model
	Name     string `gorm:"type:varchar(40)"`
	Describe string
	UserId   int64
	User     User `gorm:"foreignKey:UserId"`
}

type Collection struct {
	gorm.Model
	GoodsId     int64
	Goods       Goods `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:GoodsId"`
	FavoritesId int64
	Favorites   Favorites `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:FavoritesId"`
	UserId      int64
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserId"`
}
