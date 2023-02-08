package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/oktapascal/app-sima/models/web"
	"github.com/oktapascal/app-sima/services"
)

type MenuControllersImpl struct {
	Validate     *validator.Validate
	MenuServices services.MenuServices
}

func NewMenuControllersImpl(validate *validator.Validate, menuServices services.MenuServices) *MenuControllersImpl {
	return &MenuControllersImpl{Validate: validate, MenuServices: menuServices}
}

func (controllers *MenuControllersImpl) GetAll(ctx *fiber.Ctx) error {
	cntx := ctx.UserContext()

	response := controllers.MenuServices.GetAll(cntx)

	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          response,
	}

	return ctx.Status(fiber.StatusOK).JSON(responses)
}
