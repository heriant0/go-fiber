package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     string `gorm:"column:id"`
	Title  string `gorm:"column:title"`
	Author string `gorm:"column:author"`
}
