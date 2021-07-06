package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDatabase() (err error) {
	USER := viper.Get("DATABASE_USERNAME")
	PASS := viper.Get("DATABASE_PASSWORD")
	HOST := viper.Get("DATABASE_HOST")
	PORT := viper.Get("DATABASE_PORT")
	DBNAME := viper.Get("DATABASE_NAME")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	GlobalDB, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return

	}
	return
}
