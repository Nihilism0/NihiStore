package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Amount  int64
	GoodsId int64
	Goods   Goods `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:GoodsId"`
	UserId  int64
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserId"`
}
