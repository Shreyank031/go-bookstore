package config

//Connect with database i.e mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //make sure you add underscore and space at start
)

var (
	db *gorm.DB
)

//The whole point of this file is to return return the variable db which will help other file to interact with db

// To open the connection with our database i.e mysql
func Connect() { //make sure the function name is Uppercase

	d, err := gorm.Open("mysql", "simple:password/simple?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}

//Done with config
