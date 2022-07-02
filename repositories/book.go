package repositories

import (
	"errors"

	"github.com/heriant0/go-fiber/models"
	"gorm.io/gorm"
)

type book struct {
	storage *gorm.DB
}

func NewBookRepository(storage *gorm.DB) BookRepository {
	storage.AutoMigrate(&models.Book{})

	return book{storage: storage}
}

func (r book) Find(id int) (models.Book, error) {
	book := models.Book{}
	err := r.storage.First(&book, "id = ?", id).Error
	if book.ID == 0 {
		return book, errors.New("book not found")
	}

	return book, err
}

func (r book) Save(book *models.Book) error {
	return r.storage.Save(book).Error
}

func (r book) Delete(book *models.Book) error {
	return r.storage.Delete(book).Error
}
