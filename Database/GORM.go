package main

import (
	"fmt"
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
	profile1 := Profile{
		FirstName: "Abhishek",
		LastName:  "Kumar",
		Address:   "C-2224, Sector 43",
		Contact:   "9661572653",
		IsActive:  true,
		Company:   "Synergy Labs",
	}

	profile2 := Profile{
		FirstName: "Sid",
		LastName:  "Kumar",
		Address:   "Sikanderpur",
		Contact:   "8210xxxxx",
		IsActive:  true,
		Company:   "KRITIKAL",
	}

	db.Create(&profile1)
	db.Create(&profile2)

	// R -> Read database records
	var profiles []Profile

	db.Find(&profiles)
	fmt.Printf("%+v", profiles)

	var profileWithId Profile

	db.Find(&profileWithId, Profile{Model: gorm.Model{ID: 2}})
	fmt.Printf("%+v", profileWithId)

	// U -> Update database record
	var newProfile Profile
	db.First(&newProfile, Profile{Model: gorm.Model{ID: 2}})

	newProfile.FirstName = "Siddheswar"
	newProfile.LastName = "Ojha"

	db.Save(&newProfile)

	// D -> Delete database record
	db.Delete(&Profile{}, Profile{Model: gorm.Model{ID: 3}})
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
