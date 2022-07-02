package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/go-fiber/models"
	"github.com/heriant0/go-fiber/services"
	"github.com/heriant0/go-fiber/utils"
)

type User struct {
	Service services.User
}

func (h User) Get(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	user, err := h.Service.Get(id)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Message,
		})

		return ctx.SendStatus(err.Code)
	}
	return ctx.JSON(user)
}

func (h User) GetPaginated(ctx *fiber.Ctx) error {
	users, err := h.Service.GetPaginated(*utils.NewPaginator(ctx))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "unable to get all users",
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(users)
}

func (h User) Create(ctx *fiber.Ctx) error {
	form := models.CreateUser{}
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

func (h User) Update(ctx *fiber.Ctx) error {
	form := models.UpdateUser{}
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

func (h User) Delete(ctx *fiber.Ctx) error {
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
