package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/oktapascal/app-barayya/models/web"
	"github.com/oktapascal/app-barayya/services"
)

// ProfileControllersImpl is a struct that contains fields related to profile management.
type ProfileControllersImpl struct {
	// Validate is a pointer to a validator.Validate value, which can be used to validate data.
	Validate *validator.Validate
	// UserServices is a value of type services.UserServices, which can be used to perform actions related to user management.
	UserServices services.UserServices
}

// NewProfileControllersImpl creates and returns a new ProfileControllersImpl value with the given validator.Validate and services.UserServices values as its Validate and UserServices fields, respectively.
func NewProfileControllersImpl(validate *validator.Validate, userServices services.UserServices) *ProfileControllersImpl {
	return &ProfileControllersImpl{Validate: validate, UserServices: userServices}
}

// GetUser retrieves a user's profile information using the UserServices field, and returns it in an HTTP response.
func (controllers *ProfileControllersImpl) GetUser(ctx *fiber.Ctx) error {
	// Get the user context from the fiber.Ctx value
	cntx := ctx.UserContext()

	// Get the "user" local value stored in the fiber.Ctx value
	locals := ctx.Locals("user")
	exception, _ := locals.(*web.JwtClaims)

	// Retrieve the user's profile information
	user := controllers.UserServices.GetUserProfile(cntx, exception.IdUser)

	// Construct a web.JsonResponses value with the retrieved profile information and return it in an HTTP response
	responses := web.JsonResponses{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "OK",
		Data:          user,
	}
	return ctx.Status(fiber.StatusOK).JSON(responses)
}
