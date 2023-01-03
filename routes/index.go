package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oktapascal/app-barayya/controllers"
)

func NewRouter(app *fiber.App, authControllers controllers.AuthControllers) {
	authGroup := app.Group("/api/auth")

	NewAuthRoutes(authGroup, authControllers)
}
