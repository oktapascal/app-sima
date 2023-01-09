package exceptions

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/oktapascal/app-barayya/models/web"
	"github.com/oktapascal/app-barayya/utils"
)

// ErrorHandler handles different types of errors that may occur during the
// processing of a request. It checks for specific error types and returns
// appropriate responses to the client. If no specific error type is found,
// it returns a generic internal server error response.
func ErrorHandler(ctx *fiber.Ctx, err error) error {

	if notFoundError(ctx, err) {
		return nil
	}

	if badRequestError(ctx, err) {
		return nil
	}

	if unauthorizedError(ctx, err) {
		return nil
	}

	if unprocessableEntityError(ctx, err) {
		return nil
	}

	if validationError(ctx, err) {
		return nil
	}

	if forbiddenError(ctx, err) {
		return nil
	}

	return internalServerError(ctx, err)
}

func forbiddenError(ctx *fiber.Ctx, err interface{}) bool {
	exception := err.(*fiber.Error)
	if exception.Error() == "Forbidden" {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusForbidden,
			StatusMessage: "Forbidden",
			Data:          nil,
		}

		err := ctx.Status(fiber.StatusForbidden).JSON(response)
		utils.PanicIfError(err)

		return true
	}
	return false
}

// validationError checks if the given error is a validation error and returns a
// response with validation error messages to the client if it is. It returns
// false if the error is not a validation error.
func validationError(ctx *fiber.Ctx, err interface{}) bool {
	exceptions, ok := err.(validator.ValidationErrors)

	if ok {
		var errorFields []utils.FieldError
		var message string

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s wajib diisi", err.Field())
			default:
				message = exceptions.Error()
			}

			errorField := utils.FieldError{
				Param:   err.Field(),
				Message: message,
			}

			errorFields = append(errorFields, errorField)
		}

		response := web.JsonResponses{
			StatusCode:    fiber.StatusUnprocessableEntity,
			StatusMessage: "Unprocessable Entity",
			Data:          errorFields,
		}

		err := ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
		utils.PanicIfError(err)

		return true
	}
	return false
}

// unprocessableEntityError checks if the given error is an unprocessable entity
// error and returns a response with the error message to the client if it is.
// It returns false if the error is not an unprocessable entity error.
func unprocessableEntityError(ctx *fiber.Ctx, err error) bool {
	exceptions, ok := err.(ErrorUnprocessableEntity)

	if ok {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusUnprocessableEntity,
			StatusMessage: "Unprocessable Entity",
			Data:          exceptions,
		}

		err := ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
		utils.PanicIfError(err)

		return true
	}

	return false
}

// unauthorizedError checks if the given error is an unauthorized error and
// returns a response with the error message to the client if it is. It returns
// false if the error is not an unauthorized error.
func unauthorizedError(ctx *fiber.Ctx, err interface{}) bool {
	exceptions, ok := err.(ErrorUnauthorized)

	if ok {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusUnauthorized,
			StatusMessage: "Unauthorized",
			Data:          exceptions.Error,
		}

		err := ctx.Status(fiber.StatusUnauthorized).JSON(response)
		utils.PanicIfError(err)

		return true
	}

	return false
}

// badRequestError checks if the given error is a bad request error and returns
// a response with the error message to the client if it is. It returns false
// if the error is not a bad request error.
func badRequestError(ctx *fiber.Ctx, err interface{}) bool {
	exceptions, ok := err.(ErrorBadRequest)

	if ok {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusBadRequest,
			StatusMessage: "Bad Request",
			Data:          exceptions.Error,
		}

		err := ctx.Status(fiber.StatusBadRequest).JSON(response)
		utils.PanicIfError(err)

		return true
	}

	return false
}

// notFoundError checks if the given error is a not found error and returns a
// response with the error message to the client if it is. It returns false
// if the error is not a not found error.
func notFoundError(ctx *fiber.Ctx, err interface{}) bool {
	exceptions, ok := err.(ErrorNotFound)

	if ok {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusNotFound,
			StatusMessage: "Not Found",
			Data:          exceptions.Error,
		}

		err := ctx.Status(fiber.StatusNotFound).JSON(response)
		utils.PanicIfError(err)

		return true
	}
	return false
}

// internalServerError returns a generic internal server error response to the
// client.
func internalServerError(ctx *fiber.Ctx, err interface{}) error {
	response := web.JsonResponses{
		StatusCode:    fiber.StatusInternalServerError,
		StatusMessage: "Internal Server Error",
		Data:          nil,
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(response)
}
