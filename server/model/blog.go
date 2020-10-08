package model

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	Title string `gorm:"not null"`
	Content string
	Author	string `gorm:"not null"`
}
