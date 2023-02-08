package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/oktapascal/app-sima/controllers"
	"time"
)

func NewSettingRoutes(router fiber.Router, menuControllers controllers.MenuControllers) {
	router.Get("/test", timeout.New(func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Form Api Setting!")
	}, 10*time.Second))

	router.Get("/menus", timeout.New(menuControllers.GetAll, 10*time.Second))
}
