package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oktapascal/app-barayya/controllers"
	"github.com/oktapascal/app-barayya/middleware"
)

func NewRouter(app *fiber.App, middlewareAuth middleware.Authentication, authControllers controllers.AuthControllers) {
	authGroup := app.Group("/api/auth")

	NewAuthRoutes(authGroup, authControllers)

	profileGroup := app.Group("/api/profile", middlewareAuth.MiddlewareCookie, middlewareAuth.MiddlewareBearer)

	NewProfileRoutes(profileGroup)
}
