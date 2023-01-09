package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oktapascal/app-barayya/controllers"
	"github.com/oktapascal/app-barayya/middleware"
	"github.com/oktapascal/app-barayya/models/web"
	"strings"
)

func NewRouter(app *fiber.App,
	middlewareAuth middleware.Authentication,
	authControllers controllers.AuthControllers,
	profileControllers controllers.ProfileControllers) {
	authGroup := app.Group("/api/auth")

	NewAuthRoutes(authGroup, authControllers)

	sessionApi := authGroup.Group("/session", middlewareAuth.MiddlewareCookie, middlewareAuth.MiddlewareBearer)
	sessionApi.Get("/user-access", authControllers.GetUserAccess)

	profileGroup := app.Group("/api/profile", middlewareAuth.MiddlewareCookie, middlewareAuth.MiddlewareBearer)

	NewProfileRoutes(profileGroup, profileControllers)

	app.Use(func(ctx *fiber.Ctx) error {
		if strings.HasPrefix(ctx.OriginalURL(), "/api") {
			response := web.JsonResponses{
				StatusCode:    fiber.StatusNotFound,
				StatusMessage: "API tidak ditemukan",
				Data:          nil,
			}

			return ctx.Status(fiber.StatusNotFound).JSON(response)
		}

		return ctx.Next()
	})
}
