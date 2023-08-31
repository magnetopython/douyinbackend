package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	var err error
	dsn := C.Mysql.UserName + ":" + C.Mysql.Password + "@tcp(" + C.Mysql.IpAddress + ":" + C.Mysql.Port + ")/" + C.Mysql.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
