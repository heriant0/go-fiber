package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/go-fiber/configs"
	"github.com/heriant0/go-fiber/models"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configs.Load()
	configs.Db.AutoMigrate(&models.User{})
}
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":5000")
}
