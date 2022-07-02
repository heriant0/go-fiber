package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/go-fiber/models"
	"github.com/heriant0/go-fiber/services"
	"github.com/heriant0/go-fiber/utils"
)

type Book struct {
	Service services.Book
}

func (h Book) Get(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	book, err := h.Service.Get(id)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Message,
		})

		return ctx.SendStatus(err.Code)
	}
	return ctx.JSON(book)
}

func (h Book) Create(ctx *fiber.Ctx) error {
	form := models.CreateBook{}
	ctx.BodyParser(&form)
	message, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": message,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	model, status := h.Service.Create(form)
	if status != nil {
		ctx.JSON(map[string]interface{}{
			"message": status.Message,
		})
		return ctx.SendStatus(status.Code)
	}

	ctx.JSON(model)

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Book) Update(ctx *fiber.Ctx) error {
	form := models.UpdateBook{}
	ctx.BodyParser(&form)
	message, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": message,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	form.ID, _ = strconv.Atoi(ctx.Params("id"))
	model, status := h.Service.Update(form)
	if status != nil {
		ctx.JSON(map[string]interface{}{
			"message": status.Message,
		})

		return ctx.SendStatus(status.Code)
	}
	return ctx.JSON(model)
}

func (h Book) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	err := h.Service.Delete(id)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Message,
		})
		return ctx.SendStatus(err.Code)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
