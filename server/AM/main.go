package main

import "NihiStore/server/shared/model"

func main() {
	user := ""
	password := ""
	host := ""
	port := 3306
	name := ""
	db := initDb(user, password, host, port, name)
	db.AutoMigrate(
		&model.User{},
		&model.CustomClaims{},
	)
}
