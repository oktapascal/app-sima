package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/oktapascal/app-barayya/bootstraps"
	"github.com/oktapascal/app-barayya/models/web"
)

// Authentication is an interface that defines methods for handling
// authentication in a request.
type Authentication interface {
	MiddlewareCookie(ctx *fiber.Ctx) error
	MiddlewareBearer(ctx *fiber.Ctx) error
}

// AuthenticationImpl is a struct that implements the Authentication interface
// and has fields for handling cookies and JWTs.
type AuthenticationImpl struct {
	Cookies bootstraps.Cookie
	Jwt     bootstraps.Jwt
}

// NewAuthenticationImpl creates a new AuthenticationImpl struct.
func NewAuthenticationImpl(cookies bootstraps.Cookie, jwt bootstraps.Jwt) *AuthenticationImpl {
	return &AuthenticationImpl{Cookies: cookies, Jwt: jwt}
}

// MiddlewareCookie is a middleware that checks if the request includes a
// cookie with the specified name. If the cookie is not present, it returns a
// response indicating that the request is unauthorized. If the cookie is
// present, it sets the Authorization header with the value of the cookie and
// calls the Next function to continue processing the request.
func (middleware *AuthenticationImpl) MiddlewareCookie(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies(middleware.Cookies.GetCookieToken())
	if cookie == "" {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusUnauthorized,
			StatusMessage: "Unauthorized",
			Data:          "Cookie Access is required",
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	ctx.Request().Header.Set("Authorization", cookie)

	return ctx.Next()
}

// MiddlewareBearer is a middleware function that authenticates requests using
// JSON Web Tokens (JWTs). If the request is not authorized, it returns an
// unauthorized response to the client. If the request is authorized, it stores
// the JWT claims in the context object and passes control to the next handler.
func (middleware *AuthenticationImpl) MiddlewareBearer(ctx *fiber.Ctx) error {
	// Get the Authorization header from the request
	authorization := ctx.Get("Authorization")

	// If the header is not present, return an unauthorized response
	if authorization == "" {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusUnauthorized,
			StatusMessage: "Unauthorized",
			Data:          "Authorization is required",
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	verify := middleware.Jwt.GetJwtKey()
	claims := &web.JwtClaims{}

	// Verify the JWT using the jwt package
	withClaims, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(verify), nil
	})

	// If the verification fails, return an unauthorized response with an appropriate error message
	if err != nil {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusUnauthorized,
			StatusMessage: "Unauthorized",
			Data:          err.Error(),
		}

		if err == jwt.ErrSignatureInvalid {
			response = web.JsonResponses{
				StatusCode:    fiber.StatusUnauthorized,
				StatusMessage: "Unauthorized",
				Data:          "Signature is invalid",
			}
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// If the token is not valid, return an unauthorized response
	if !withClaims.Valid {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusUnauthorized,
			StatusMessage: "Unauthorized",
			Data:          "Token is invalid",
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	// If the request is authorized, store the JWT claims in the context object and pass control to the next handler
	claims = withClaims.Claims.(*web.JwtClaims)
	ctx.Locals("user", claims)

	return ctx.Next()
}
