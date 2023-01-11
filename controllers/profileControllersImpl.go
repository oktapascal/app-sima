package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/oktapascal/app-barayya/models/web"
	"github.com/oktapascal/app-barayya/services"
	"github.com/oktapascal/app-barayya/utils"
)

// ProfileControllersImpl struct is used to handle HTTP requests related to profile information
type ProfileControllersImpl struct {
	// Validate is used to validate data before performing any actions on them
	Validate *validator.Validate
	// UserServices is an instance of a struct that contains methods for interacting with User data
	UserServices services.UserServices
	// KaryawanServices is an instance of a struct that contains methods for interacting with Karyawan data
	KaryawanServices services.KaryawanServices
}

// NewProfileControllersImpl creates new instance of the ProfileControllersImpl struct
// it takes three arguments as input:
// - validate : an instance of validator.Validate that will be used to validate the data
// - userServices : an instance of UserServices that contains methods for interacting with User data
// - karyawanServices : an instance of KaryawanServices that contains methods for interacting with Karyawan data
func NewProfileControllersImpl(validate *validator.Validate, userServices services.UserServices, karyawanServices services.KaryawanServices) *ProfileControllersImpl {
	// create new instance of the struct and set its fields with the values of the input parameters
	return &ProfileControllersImpl{Validate: validate, UserServices: userServices, KaryawanServices: karyawanServices}
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

// UpdateUser is a method of ProfileControllersImpl struct, it is used to handle an HTTP request to update the profile information of a user
func (controllers *ProfileControllersImpl) UpdateUser(ctx *fiber.Ctx) error {
	// create a new instance of KaryawanUpdateRequest
	request := new(web.KaryawanUpdateRequest)
	// parse the request body into request object
	err := ctx.BodyParser(request)
	// check for any error that might occur during parsing
	utils.PanicIfError(err)

	err = controllers.Validate.Struct(request)
	// check for any validation error
	utils.PanicIfError(err)

	// get the user context
	cntx := ctx.UserContext()

	// get the user id from the request's context
	locals := ctx.Locals("user")
	exception, _ := locals.(*web.JwtClaims)

	// set the request's IdUser to the user id that was got from the context
	request.IdUser = exception.IdUser

	// update the user's profile information by calling the Update method on the KaryawanServices
	controllers.KaryawanServices.Update(cntx, *request)
	// create a new instance of JsonResponses to send the response
	responses := web.JsonResponses{
		// set status code as OK
		StatusCode: fiber.StatusOK,
		// set status message as OK
		StatusMessage: "OK",
		// set Data field as nil
		Data: nil,
	}
	// return the response to the client with status code as OK and json responses
	return ctx.Status(fiber.StatusOK).JSON(responses)
}
