package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string     	`gorm:"type:varchar(20);not null;"`
	Age           string     	`gorm:"type:varchar(10);not null;"`
	Sex           string     	`gorm:"type:varchar(20);not null;"`
	Mobile		  string 		`gorm:"type:varchar(20);not null;unique_index"`
}

