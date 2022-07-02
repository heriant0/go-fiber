package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/go-fiber/configs"
	"github.com/heriant0/go-fiber/handlers"
	"github.com/heriant0/go-fiber/repositories"
	"github.com/heriant0/go-fiber/services"
)

func RegisterUserRoutes(router fiber.Router) {
	service := services.User{
		Repository: repositories.NewUserRepository(configs.Db),
	}
	user := handlers.User{
		Service: service,
	}

	router.Post("/users", user.Create)
	router.Put("/users/:id", user.Update)
	router.Delete("/users/:id", user.Delete)
	router.Get("/users/:id", user.Get)
	router.Get("/users", user.GetPaginated)
}
