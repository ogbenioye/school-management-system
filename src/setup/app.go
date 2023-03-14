package setup

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connection() {
	d, err := gorm.Open("mysql", "root:my_secret@tcp(mysql:3306)/sms-mysql?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
