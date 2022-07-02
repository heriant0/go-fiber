package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/go-fiber/models"
	"github.com/heriant0/go-fiber/repositories"
	"github.com/heriant0/go-fiber/utils"
	"github.com/jinzhu/copier"
)

type User struct {
	Repository repositories.UserRepository
}

func (s User) Get(id int) (models.User, *models.Error) {
	result, err := s.Repository.Find(id)
	fmt.Println(result, "result")
	if err != nil {
		return result, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	return result, nil
}

func (s User) GetUserName(username string) (models.User, *models.Error) {
	result, err := s.Repository.FindByUsername(username)
	if err != nil {
		return result, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	return result, nil
}

func (s User) GetPaginated(paginator utils.Paginator) (utils.Paginator, *models.Error) {
	result, err := s.Repository.FindPaginated(paginator)
	if err != nil {
		return result, &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}

	return result, nil
}

func (s User) Create(user models.CreateUser) (models.User, *models.Error) {
	model := models.User{}
	copier.Copy(&model, &user)
	err := s.Repository.Save(&model)
	if err != nil {
		return model, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}
	return model, nil
}

func (s User) Update(user models.UpdateUser) (models.User, *models.Error) {
	model, err := s.Repository.Find(user.ID)
	if err != nil {
		return model, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	copier.Copy(&model, &user)
	err = s.Repository.Save(&model)
	if err != nil {
		return model, &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}

	return model, nil
}

func (s User) Delete(id int) *models.Error {
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
