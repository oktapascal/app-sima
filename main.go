package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/oktapascal/app-sima/bootstraps"
	"github.com/oktapascal/app-sima/controllers"
	"github.com/oktapascal/app-sima/database"
	"github.com/oktapascal/app-sima/exceptions"
	"github.com/oktapascal/app-sima/middleware"
	"github.com/oktapascal/app-sima/repository"
	"github.com/oktapascal/app-sima/routes"
	"github.com/oktapascal/app-sima/services"
	"github.com/oktapascal/app-sima/utils"
	"reflect"
	"strings"
	"time"
)

func main() {
	config := bootstraps.New(".env")
	jwtConfig := bootstraps.NewJwtImpl(config)
	cookiesConfig := bootstraps.NewCookieImpl(config)

	app := fiber.New(fiber.Config{
		ErrorHandler: exceptions.ErrorHandler,
	})

	//mySql := database.NewMysql(config)

	ctx := context.Background()
	fireStore := database.NewFirestoreClient(config, ctx)

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	middlewareAuth := middleware.NewAuthenticationImpl(cookiesConfig, jwtConfig)

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "cookie:csrf_",
		CookieName:     "csrf_",
		CookieHTTPOnly: true,
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.GenerateUUID,
	}))

	userRepository := repository.NewUserRepositoryImpl()
	userServices := services.NewUserServicesImpl(userRepository, fireStore)

	authControllers := controllers.NewAuthControllerImpl(validate, userServices, jwtConfig, cookiesConfig)

	routes.NewRouter(app, middlewareAuth, authControllers)

	// load static files
	app.Static("/", "./ui/dist")
	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./ui/dist/index.html")
	})

	server := app.Listen(":8080")
	utils.PanicIfError(server)
}
