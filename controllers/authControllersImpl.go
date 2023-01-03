package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/oktapascal/app-barayya/bootstraps"
	"github.com/oktapascal/app-barayya/exceptions"
	"github.com/oktapascal/app-barayya/models/web"
	"github.com/oktapascal/app-barayya/services"
	"github.com/oktapascal/app-barayya/utils"
)

type AuthControllerImpl struct {
	Validate     *validator.Validate
	UserServices services.UserServices
	JwtConfig    bootstraps.Jwt
	CookieConfig bootstraps.Cookie
}

func NewAuthControllerImpl(validate *validator.Validate, userServices services.UserServices, jwtConfig bootstraps.Jwt, cookieConfig bootstraps.Cookie) *AuthControllerImpl {
	return &AuthControllerImpl{
		Validate:     validate,
		UserServices: userServices,
		JwtConfig:    jwtConfig,
		CookieConfig: cookieConfig,
	}
}

func (controllers *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	request := new(web.RegisterRequest)

	err := ctx.BodyParser(request)
	utils.PanicIfError(err)

	err = controllers.Validate.Struct(request)
	utils.PanicIfError(err)

	cntx := ctx.UserContext()

	controllers.UserServices.Save(cntx, *request)

	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(responses)
}

func (controllers *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	request := new(web.LoginRequest)

	err := ctx.BodyParser(request)
	utils.PanicIfError(err)

	err = controllers.Validate.Struct(request)
	utils.PanicIfError(err)

	cntx := ctx.UserContext()

	user := controllers.UserServices.CheckUsername(cntx, request.Username)

	isValidPassword := utils.ValidatePassword(request.Password, user.Password)
	if !isValidPassword {
		panic(exceptions.NewErrorUnprocessableEntity("Password tidak valid", "password"))
	}

	token, expiration, err := controllers.JwtConfig.GenerateAccessToken(user)
	if err != nil {
		panic(exceptions.NewErrorUnauthorized("token is incorrect"))
	}

	var session web.SessionRequest
	session.AuthToken = token
	session.IdUser = user.Id

	controllers.UserServices.SaveSessionUser(cntx, session)

	controllers.CookieConfig.CreateTokenAndCookie(token, expiration, ctx)

	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(responses)
}

func (controllers *AuthControllerImpl) Logout(ctx *fiber.Ctx) error {
	cntx := ctx.UserContext()

	authToken := ctx.Cookies(controllers.CookieConfig.GetCookieToken())

	controllers.UserServices.DestroySessionUser(cntx, authToken)

	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(responses)
}
