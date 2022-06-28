package models

type Book struct {
	ID     string `gorm:"column:id"`
	Title  string `gorm:"column:title"`
	Author string `gorm:"column:author"`
}
