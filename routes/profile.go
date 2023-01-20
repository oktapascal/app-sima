package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"time"
)

func NewProfileRoutes(router fiber.Router) {
	router.Get("/test", timeout.New(func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Form Api Profile!")
	}, 10*time.Second))
}
