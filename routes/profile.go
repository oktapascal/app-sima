package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/oktapascal/app-barayya/controllers"
	"time"
)

func NewProfileRoutes(router fiber.Router, profileControllers controllers.ProfileControllers) {
	router.Get("/test", timeout.New(func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Form Api Profile!")
	}, 10*time.Second))

	router.Get("/user", timeout.New(profileControllers.GetUser, 10*time.Second))
	router.Put("/user", timeout.New(profileControllers.UpdateUser, 10*time.Second))
}
