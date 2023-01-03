package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/oktapascal/app-barayya/bootstraps"
	"github.com/oktapascal/app-barayya/models/web"
)

type Authentication interface {
	MiddlewareCookie(ctx *fiber.Ctx) error
	MiddlewareBearer(ctx *fiber.Ctx) error
}

type jwtClaims struct {
	IdUser     uint   `json:"id_user"`
	KodeLokasi string `json:"kode_lokasi"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}

type AuthenticationImpl struct {
	Cookies bootstraps.Cookie
	Jwt     bootstraps.Jwt
}

func NewAuthenticationImpl(cookies bootstraps.Cookie, jwt bootstraps.Jwt) *AuthenticationImpl {
	return &AuthenticationImpl{Cookies: cookies, Jwt: jwt}
}

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

func (middleware *AuthenticationImpl) MiddlewareBearer(ctx *fiber.Ctx) error {
	authorization := ctx.Get("Authorization")

	if authorization == "" {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusUnauthorized,
			StatusMessage: "Unauthorized",
			Data:          "Authorization is required",
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	verify := middleware.Jwt.GetJwtKey()
	claims := &jwtClaims{}

	withClaims, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(verify), nil
	})

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

	if !withClaims.Valid {
		response := web.JsonResponses{
			StatusCode:    fiber.StatusUnauthorized,
			StatusMessage: "Unauthorized",
			Data:          "Token is invalid",
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	claims = withClaims.Claims.(*jwtClaims)
	ctx.Locals("user", claims)

	return ctx.Next()
}
