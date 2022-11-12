package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() *gorm.DB {
	dsn := "root@chirag@7(127.0.0.1:3306)/go_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
func DB() *gorm.DB {
	return Db
}
