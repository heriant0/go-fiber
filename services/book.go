package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/go-fiber/models"
	"github.com/heriant0/go-fiber/repositories"
	"github.com/jinzhu/copier"
)

type Book struct {
	Repository repositories.BookRepository
}

func (s Book) Get(id int) (models.Book, *models.Error) {
	result, err := s.Repository.Find(id)
	if err != nil {
		return result, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	return result, nil
}

func (s Book) Create(book models.CreateBook) (models.Book, *models.Error) {
	model := models.Book{}
	copier.Copy(&model, &book)
	err := s.Repository.Save(&model)
	if err != nil {
		return model, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}
	return model, nil
}

func (s Book) Update(book models.UpdateBook) (models.Book, *models.Error) {
	model, err := s.Repository.Find(book.ID)
	if err != nil {
		return model, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}
	copier.Copy(&model, &book)
	err = s.Repository.Save(&model)
	if err != nil {
		return model, &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}
	return model, nil
}

func (s Book) Delete(id int) *models.Error {
	model, err := s.Repository.Find(id)
	if err != nil {
		return &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	err = s.Repository.Delete(&model)
	if err != nil {
		return &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}
	return nil
}
