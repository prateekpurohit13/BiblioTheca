package config

import(
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "name:AddyourSQLHere@/simplerest?charset = utf8&parseTime = True&loc = Local") // Kindly add your db here in the specified format.
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
