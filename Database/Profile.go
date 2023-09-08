package main

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	FirstName string
	LastName  string
	Address   string
	Contact   string
	IsActive  bool
	Company   string
}
