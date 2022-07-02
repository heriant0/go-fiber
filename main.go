package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/go-fiber/configs"
	"github.com/heriant0/go-fiber/models"
	"github.com/heriant0/go-fiber/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configs.Load()
	configs.Db.AutoMigrate(&models.User{})
}
func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	routes.RegisterUserRoutes(app)
	routes.RegisterBookRoutes(app)

	// app.Listen(":5000")
	app.Listen(fmt.Sprintf(":%d", configs.Env.AppPort))
}
