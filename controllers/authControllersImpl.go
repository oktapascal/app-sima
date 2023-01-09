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

// AuthControllersImpl is a struct that represents a controller for handling authentication-related actions.
// It has fields for a validator.Validate, services.UserServices, bootstraps.Jwt, and bootstraps.Cookie.
type AuthControllersImpl struct {
	// Validate is a field of type *validator.Validate that is used to validate input data.
	Validate *validator.Validate
	// UserServices is a field of type services.UserServices that is used to perform actions related to users.
	UserServices services.UserServices
	// JwtConfig is a field of type bootstraps.Jwt that holds configuration information for JWT.
	JwtConfig bootstraps.Jwt
	// CookieConfig is a field of type bootstraps.Cookie that holds configuration information for cookies.
	CookieConfig bootstraps.Cookie
}

// NewAuthControllerImpl is a function that returns a new *AuthControllersImpl struct.
// It takes in four input parameters:
// - a validate of type *validator.Validate
// - a userServices of type services.UserServices
// - a jwtConfig of type bootstraps.Jwt
// - a cookieConfig of type bootstraps.Cookie
func NewAuthControllerImpl(validate *validator.Validate, userServices services.UserServices, jwtConfig bootstraps.Jwt, cookieConfig bootstraps.Cookie) *AuthControllersImpl {
	// Return a new *AuthControllersImpl struct with its fields initialized with the input parameters.
	return &AuthControllersImpl{
		Validate:     validate,
		UserServices: userServices,
		JwtConfig:    jwtConfig,
		CookieConfig: cookieConfig,
	}
}

// Register handles a request to register a new user.
func (controllers *AuthControllersImpl) Register(ctx *fiber.Ctx) error {
	// Initialize a RegisterRequest struct and parse the request body into it.
	request := new(web.RegisterRequest)
	err := ctx.BodyParser(request)
	utils.PanicIfError(err)

	// Validate the request.
	err = controllers.Validate.Struct(request)
	utils.PanicIfError(err)

	// Save the request to storage.
	cntx := ctx.UserContext()
	controllers.UserServices.Save(cntx, *request)

	// Construct and return a JSON response with status code 200 OK.
	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(responses)
}

// Login handles a request to log in an existing user.
func (controllers *AuthControllersImpl) Login(ctx *fiber.Ctx) error {
	// Initialize a LoginRequest struct and parse the request body into it.
	request := new(web.LoginRequest)
	err := ctx.BodyParser(request)
	utils.PanicIfError(err)

	// Validate the request.
	err = controllers.Validate.Struct(request)
	utils.PanicIfError(err)

	// Check if the provided username is valid.
	cntx := ctx.UserContext()
	user := controllers.UserServices.CheckUsername(cntx, request.Username)

	// Validate the provided password.
	isValidPassword := utils.ValidatePassword(request.Password, user.Password)
	if !isValidPassword {
		// If the password is invalid, raise an error.
		panic(exceptions.NewErrorUnprocessableEntity("Password tidak valid", "password"))
	}

	// Generate an access token and expiration date.
	token, expiration, err := controllers.JwtConfig.GenerateAccessToken(user)
	if err != nil {
		// If there is an error generating the token, raise an error.
		panic(exceptions.NewErrorUnauthorized("token is incorrect"))
	}

	// Save the session information.
	var session web.SessionRequest
	session.AuthToken = token
	session.IdUser = user.Id
	controllers.UserServices.SaveSessionUser(cntx, session)
	userSession := controllers.UserServices.GetSessionUser(cntx, token)

	// Create a token cookie.
	controllers.CookieConfig.CreateTokenAndCookie(token, expiration, ctx)

	// Construct and return a JSON response with status code 200 OK.
	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          userSession,
	}
	return ctx.Status(fiber.StatusOK).JSON(responses)
}

// Logout handles a request to log out the current user.
func (controllers *AuthControllersImpl) Logout(ctx *fiber.Ctx) error {
	// Get the user context and the current user's auth token.
	cntx := ctx.UserContext()
	authToken := ctx.Cookies(controllers.CookieConfig.GetCookieToken())

	// Destroy the current user's session.
	controllers.UserServices.DestroySessionUser(cntx, authToken)

	// Delete the auth token cookie.
	controllers.CookieConfig.DeleteCookie(ctx)

	// Construct and return a JSON response with status code 200 OK.
	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(responses)
}

// GetUserAccess handles a request to get the current user's access information.
func (controllers *AuthControllersImpl) GetUserAccess(ctx *fiber.Ctx) error {
	// Get the user context and the current user's auth token.
	cntx := ctx.UserContext()
	authToken := ctx.Cookies(controllers.CookieConfig.GetCookieToken())

	// Get the current user's session information.
	user := controllers.UserServices.GetSessionUser(cntx, authToken)

	// Construct and return a JSON response with status code 200 OK.
	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          user,
	}
	return ctx.Status(fiber.StatusOK).JSON(responses)
}
