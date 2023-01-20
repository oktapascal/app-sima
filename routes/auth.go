package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/oktapascal/app-sima/controllers"
	"time"
)

func NewAuthRoutes(router fiber.Router, authControllers controllers.AuthControllers) {
	router.Get("/test", timeout.New(func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Form Api Auth!")
	}, 10*time.Second))

	router.Post("/register", timeout.New(authControllers.Register, 10*time.Second))
	router.Post("/login", timeout.New(authControllers.Login, 10*time.Second))
}
