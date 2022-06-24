package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/go-fiber/configs"
	"github.com/heriant0/go-fiber/models"
)

func main() {
	db := configs.Load()
	app := fiber.New()

	db.AutoMigrate(&models.Book{})

	db.Create(&models.Book{ID: "A001", Title: "Golang", Author: "google team"})

	var books models.Book
	findById := db.First(&books, "id = ?", "A001")
	println(findById)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
