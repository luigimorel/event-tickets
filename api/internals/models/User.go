package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName   string `gorm:"size:255;not null;" json:"firstName"`
	LastName    string `gorm:"size:255;not null;" json:"lastName"`
	PhoneNumber string `gorm:"type:text;not null;unique" json:"phoneNumber"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	Password    string `gorm:"size:100;not null;" json:"password"`
}
