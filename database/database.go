package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDatabase() (err error) {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "edesa-go"

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	GlobalDB, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return

	}
	return
}
