package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/go-fiber/configs"
	"github.com/heriant0/go-fiber/handlers"
	"github.com/heriant0/go-fiber/repositories"
	"github.com/heriant0/go-fiber/services"
)

func RegisterBookRoutes(router fiber.Router) {
	service := services.Book{
		Repository: repositories.NewBookRepository(configs.Db),
	}

	book := handlers.Book{
		Service: service,
	}

	router.Post("/books", book.Create)
	router.Put("/books/:id", book.Update)
	router.Delete("/books/:id", book.Delete)
	router.Get("/books/:id", book.Get)

}
