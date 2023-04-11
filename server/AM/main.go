package main

import (
	"NihiStore/server/shared/consts"
	"NihiStore/server/shared/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	user := ""
	password := ""
	host := ""
	port := 3306
	name := ""
	db := initDb(user, password, host, port, name)
	_ = db.AutoMigrate(
		&model.Goods{},
	)
}

func initDb(user, password, host string, port int, name string) *gorm.DB {
	dsn := fmt.Sprintf(consts.MySqlDSN, user, password, host, port, name) //user password host port name
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
