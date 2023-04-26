package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"type:varchar(40)"`
	Password    string
	SellerAliId string
	HeadId      int64
}
