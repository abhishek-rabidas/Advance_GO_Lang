package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
Install these dependencies:
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/mysql
*/

func main() {

	var db *gorm.DB = InitializeDB()

	// C -> Create database table record

}

func InitializeDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/syndb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // connect to the datasource name

	if err != nil {
		panic("Error connecting to db")
	}

	err = db.AutoMigrate(&Profile{})

	if err != nil {
		panic(err)
	}

	return db
}
