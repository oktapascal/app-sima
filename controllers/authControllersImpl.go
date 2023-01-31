package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/oktapascal/app-sima/bootstraps"
	"github.com/oktapascal/app-sima/exceptions"
	"github.com/oktapascal/app-sima/models/web"
	"github.com/oktapascal/app-sima/services"
	"github.com/oktapascal/app-sima/utils"
)

type AuthControllersImpl struct {
	Validate     *validator.Validate
	AuthServices services.AuthServices
	JwtConfig    bootstraps.Jwt
	CookieConfig bootstraps.Cookie
}

func NewAuthControllerImpl(validate *validator.Validate, authServices services.AuthServices, jwtConfig bootstraps.Jwt, cookieConfig bootstraps.Cookie) *AuthControllersImpl {
	return &AuthControllersImpl{
		Validate:     validate,
		AuthServices: authServices,
		JwtConfig:    jwtConfig,
		CookieConfig: cookieConfig,
	}
}

func (controllers *AuthControllersImpl) Register(ctx *fiber.Ctx) error {
	request := new(web.RegisterRequest)
	err := ctx.BodyParser(request)
	utils.PanicIfError(err)

	err = controllers.Validate.Struct(request)
	utils.PanicIfError(err)

	cntx := ctx.UserContext()

	controllers.AuthServices.Register(cntx, *request)

	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(responses)
}

func (controllers *AuthControllersImpl) Login(ctx *fiber.Ctx) error {
	request := new(web.LoginRequest)
	err := ctx.BodyParser(request)
	utils.PanicIfError(err)

	err = controllers.Validate.Struct(request)
	utils.PanicIfError(err)

	cntx := ctx.UserContext()
	user := controllers.AuthServices.Login(cntx, *request)

	token, expiration, errGenerate := controllers.JwtConfig.GenerateAccessToken(user)
	if errGenerate != nil {
		panic(exceptions.NewErrorUnauthorized("token is incorrect"))
	}

	session := web.SessionRequest{
		IdUser:    user.IdUser,
		Nik:       user.Nik,
		AuthToken: token,
	}

	controllers.AuthServices.StoreSessionUser(cntx, session)

	controllers.CookieConfig.CreateTokenAndCookie(token, expiration, ctx)

	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          user,
	}

	return ctx.Status(fiber.StatusOK).JSON(responses)
}

func (controllers *AuthControllersImpl) Logout(ctx *fiber.Ctx) error {
	cntx := ctx.UserContext()
	authToken := ctx.Cookies(controllers.CookieConfig.GetCookieToken())

	locals := ctx.Locals("user")
	exception, _ := locals.(*web.JwtClaims)

	session := web.SessionRequest{
		IdUser:    exception.IdUser,
		AuthToken: authToken,
	}

	controllers.AuthServices.Logout(cntx, session)

	controllers.CookieConfig.DeleteCookie(ctx)

	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
	}
	return ctx.Status(fiber.StatusOK).JSON(responses)
}

func (controllers *AuthControllersImpl) GetUserProfile(ctx *fiber.Ctx) error {
	cntx := ctx.UserContext()

	locals := ctx.Locals("user")
	exception, _ := locals.(*web.JwtClaims)

	user := controllers.AuthServices.GetUserProfile(cntx, exception.Nik)

	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          user,
	}

	return ctx.Status(fiber.StatusOK).JSON(responses)
}

func (controllers *AuthControllersImpl) UpdateUserProfile(ctx *fiber.Ctx) error {
	cntx := ctx.UserContext()

	request := new(web.UpdateUserProfileRequest)
	err := ctx.BodyParser(request)
	utils.PanicIfError(err)

	err = controllers.Validate.Struct(request)
	utils.PanicIfError(err)

	locals := ctx.Locals("user")
	exception, _ := locals.(*web.JwtClaims)

	request.Nik = exception.Nik

	controllers.AuthServices.UpdateUserProfile(cntx, *request)

	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
	}

	return ctx.Status(fiber.StatusOK).JSON(responses)
}
