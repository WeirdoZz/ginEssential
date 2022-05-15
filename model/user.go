package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"varchar(20);not null"`
	PhoneNumber string `gorm:"type:varchar(11);not null;unique"`
	Password    string `gorm:"size:255;not null"`
}
