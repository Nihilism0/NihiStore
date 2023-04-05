package model

import "gorm.io/gorm"

type Favorties struct {
	gorm.Model
	Name     string `gorm:"type:varchar(40)"`
	Describe string
	User_Id  int64
	User     User `gorm:"foreignKey:User_Id"`
}

type Collection struct {
	gorm.Model
	Goods_Id     int64
	Goods        Goods `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Goods_Id"`
	Favorites_Id int64
	Favorites    Favorties `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Favorites_Id"`
	User_Id      int64
	User         User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:User_Id"`
}
