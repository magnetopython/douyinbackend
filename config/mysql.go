package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() error {
	var err error
	dsn := C.Mysql.UserName + ":" + C.Mysql.Password + "@tcp(" + C.Mysql.IpAddress + ":" + C.Mysql.Port + ")/" + C.Mysql.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
