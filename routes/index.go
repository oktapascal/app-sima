package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/oktapascal/app-barayya/controllers"
	"github.com/oktapascal/app-barayya/middleware"
	"github.com/oktapascal/app-barayya/models/web"
	"strings"
	"time"
)

func NewRouter(app *fiber.App,
	middlewareAuth middleware.Authentication,
	authControllers controllers.AuthControllers) {
	authGroup := app.Group("/api/auth")
	NewAuthRoutes(authGroup, authControllers)

	authGroupWithMiddleware := authGroup.Group("/", middlewareAuth.MiddlewareCookie, middlewareAuth.MiddlewareBearer)
	authGroupWithMiddleware.Post("/logout", timeout.New(authControllers.Logout, 10*time.Second))
	authGroupWithMiddleware.Post("/profile", timeout.New(authControllers.UpdateUserProfile, 10*time.Second))
	authGroupWithMiddleware.Get("/profile", timeout.New(authControllers.GetUserProfile, 10*time.Second))

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
