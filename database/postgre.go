package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DbConn *gorm.DB
)

func InitDB(){
	dsn := "host=localhost user=postgres password=Password! dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database tidak dapat terkoneksi")
	}
}